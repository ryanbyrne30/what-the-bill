package main

import (
	"context"
	"flag"
	"log"

	"github.com/ryanbyrne30/what-the-bill/monitor/queue/queue/mongo"
)

var (
	mongoUrl            = flag.String("mongoUrl", "mongo+srv://localhost:27017", "Mongo URL")
	mongoDb             = flag.String("mongoDb", "queue", "MongoDB database name")
	mongoCol            = flag.String("mongoCol", "bill_events", "MongoDB collection name")
	consumerUSBillsAddr = flag.String("consumerUSBillsAddr", "localhost:6052", "Address of consumer for US bills")
	batchSize           = flag.Int("batchSize", 1000, "Size of batches for dispatching events")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	db := mongo.NewMongo(*mongoUrl, *mongoDb, *mongoCol)
	defer db.Disconnect(ctx)

	client := NewProtoClient(*consumerUSBillsAddr)

	for i := 0; i < 1_000_000_000; i++ {
		eventCount, err := DispatchEvents(ctx, db, client)
		if err != nil {
			log.Printf("ERROR dispatching events: %v", err)
		}
		if eventCount == 0 {
			break
		}
	}
}

func DispatchEvents(ctx context.Context, db *mongo.Mongo, client *ProtoClient) (int, error) {
	events, err := db.GetEvents(ctx, *batchSize)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, event := range events {
		err = DispatchEvent(ctx, client, event)
		if err != nil {
			log.Printf("ERROR dispatching event %s: %v", event.EventName, err)
			db.DeleteEvent(ctx, event.ID)
			db.SaveEvent(ctx, &event)
		} else if err = db.DeleteEvent(ctx, event.ID); err != nil {
			log.Printf("ERROR deleting event: %v", err)
		} else {
			count += 1
		}
	}

	return count, nil
}

func DispatchEvent(ctx context.Context, client *ProtoClient, event mongo.Event) error {
	switch event.EventName {
	case string(mongo.USBillCreatedEventName):
		return DispatchUSBillCreatedEvent(ctx, client, event)
	case string(mongo.USBillUpdatedEventName):
		return DispatchUSBillUpdatedEvent(ctx, client, event)
	default:
		return nil
	}
}
