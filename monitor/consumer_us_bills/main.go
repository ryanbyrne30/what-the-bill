package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/ryanbyrne30/what-the-bill/monitor/queue/bills_us"
	"github.com/ryanbyrne30/what-the-bill/monitor/queue/proto_consumer_us_bills"
	"google.golang.org/grpc"
)

var (
	port     = flag.Int("port", 6052, "Port for server to listen on")
	apiKey   = flag.String("apiKey", "", "API Key for govinfo.gov")
	mongoUrl = flag.String("mongoUrl", "mongo+srv://localhost:27017", "MongoDB connection url")
	mongoDb  = flag.String("mongoDb", "bills", "MongoDB database name")
	mongoCol = flag.String("mongoCol", "us_bills", "MongoDB collection name")
)

type server struct {
	mongo *bills_us.Mongo
	fetch *Fetch
	proto_consumer_us_bills.UnimplementedConsumerUSBillsServer
}

func (s *server) SendUSBillCreated(ctx context.Context, in *proto_consumer_us_bills.USBillCreatedEvent) (*proto_consumer_us_bills.EventPosted, error) {
	log.Printf("received bill created event: %s", in.GetPackageId())
	bill := &bills_us.Bill{BillID: in.GetPackageId()}

	details, err := s.fetch.FetchBillDetails(in.GetPackageLink())
	if err != nil {
		return nil, err
	}

	text, err := s.fetch.FetchBillText(details.Download.TextLink)
	if err != nil {
		return nil, err
	}
	bill.Text = text

	actionsResponse, err := s.fetch.FetchBillActions(details.Related.BillStatusLink)
	if err != nil {
		return nil, err
	}
	for _, action := range actionsResponse.Bill.Actions.Items {
		d, err := time.Parse("2006-01-02", action.ActionDate)
		if err != nil {
			log.Printf("Could not parse action date: %v", err)
			continue
		}
		bill.Actions = append(bill.Actions, bills_us.BillActions{
			Date:   d,
			Action: action.Text,
		})
	}

	bill.Title = details.Title
	bill.Url = details.Url
	bill.Congress = details.Congress
	bill.Session = details.Session
	bill.Version = details.Version
	bill.Type = details.Type

	pages, _ := strconv.Atoi(details.Pages)
	bill.Pages = pages

	if len(details.ShortTitle) > 0 {
		bill.ShortTitle = details.ShortTitle[0].Title
	}

	publishedAt, err := time.Parse("2006-01-02", details.PublishedAt)
	if err != nil {
		log.Printf("ERROR parsing publish date '%s': %v", details.PublishedAt, err)
		return nil, err
	}
	bill.PublishedAt = publishedAt

	updatedAt, err := time.Parse("2006-01-02T15:04:05Z", details.UpdatedAt)
	if err != nil {
		log.Printf("ERROR parsing update date '%s': %v", details.UpdatedAt, err)
		return nil, err
	}
	bill.UpdatedAt = updatedAt

	for _, member := range details.Members {
		bill.Members = append(bill.Members, bills_us.BillMembers{
			Role:    member.Role,
			Chamber: member.Chamber,
			BioID:   member.BioID,
			Name:    member.Name,
			State:   member.State,
			Party:   member.Party,
		})
	}

	for _, committee := range details.Committees {
		bill.Committees = append(bill.Committees, bills_us.BillCommittees{
			AuthorityID: committee.AuthorityID,
			Chamber:     committee.Chamber,
			Name:        committee.Name,
			Type:        committee.Type,
		})
	}

	err = s.mongo.InsertBill(ctx, bill)
	if err != nil {
		return nil, err
	}

	return &proto_consumer_us_bills.EventPosted{Status: "Ok"}, nil
}

func (s *server) SendUSBillUpdated(ctx context.Context, in *proto_consumer_us_bills.USBillUpdatedEvent) (*proto_consumer_us_bills.EventPosted, error) {
	log.Printf("received update bill event: %s", in.PackageId)
	return &proto_consumer_us_bills.EventPosted{Status: "Ok"}, nil
}

func main() {
	flag.Parse()

	ctx := context.Background()

	db := bills_us.NewMongo(*mongoUrl, *mongoDb, *mongoCol)
	defer db.Disconnect(ctx)

	fetch := NewFetch(*apiKey)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto_consumer_us_bills.RegisterConsumerUSBillsServer(s, &server{mongo: db, fetch: fetch})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
