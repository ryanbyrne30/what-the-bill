package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ryanbyrne30/what-the-bill/monitor/queue/proto_queue"
	"github.com/ryanbyrne30/what-the-bill/monitor/queue/queue/mongo"
	"google.golang.org/grpc"
)

var (
	port     = flag.Int("port", 6051, "Port for server to listen on")
	mongoUrl = flag.String("mongoUrl", "mongo+srv://localhost:27017", "MongoDB connection url")
	mongoDb  = flag.String("mongoDb", "queue", "MongoDB database name")
	mongoCol = flag.String("mongoCol", "bill_events", "MongoDB collection name")
)

type server struct {
	mongo *mongo.Mongo
	proto_queue.UnimplementedQueueServer
}

func (s *server) PostUSBillUpdatedEvent(ctx context.Context, in *proto_queue.USBillUpdatedEvent) (*proto_queue.EventPosted, error) {
	log.Printf("Received event: %s", mongo.USBillUpdatedEventName)
	updatedAt, err := time.Parse("2006-01-02T15:04:05Z", in.UpdatedAt)
	if err != nil {
		return &proto_queue.EventPosted{Status: "Error parsing UpdatedAt"}, err
	}

	message, err := json.Marshal(&mongo.USBillUpdatedEvent{
		PackageID:   in.PackageId,
		UpdatedAt:   updatedAt,
		PackageLink: in.PackageLink,
	})
	if err != nil {
		return &proto_queue.EventPosted{Status: "Error"}, err
	}

	s.mongo.SaveEvent(ctx, &mongo.Event{
		EventName: string(mongo.USBillUpdatedEventName),
		Message:   string(message),
	})

	return &proto_queue.EventPosted{Status: "Ok"}, nil
}

func (s *server) PostUSBillCreatedEvent(ctx context.Context, in *proto_queue.USBillCreatedEvent) (*proto_queue.EventPosted, error) {
	log.Printf("Received event: %s", mongo.USBillCreatedEventName)
	updatedAt, err := time.Parse("2006-01-02T15:04:05Z", in.UpdatedAt)
	if err != nil {
		return &proto_queue.EventPosted{Status: "Error parsing UpdatedAt"}, err
	}

	message, err := json.Marshal(&mongo.USBillCreatedEvent{
		PackageID:   in.PackageId,
		UpdatedAt:   updatedAt,
		PackageLink: in.PackageLink,
	})
	if err != nil {
		return &proto_queue.EventPosted{Status: "Error"}, err
	}

	s.mongo.SaveEvent(ctx, &mongo.Event{
		EventName: string(mongo.USBillCreatedEventName),
		Message:   string(message),
	})

	return &proto_queue.EventPosted{Status: "Ok"}, nil
}

func main() {
	flag.Parse()

	ctx := context.Background()

	db := mongo.NewMongo(*mongoUrl, *mongoDb, *mongoCol)
	defer db.Disconnect(ctx)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto_queue.RegisterQueueServer(s, &server{mongo: db})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
