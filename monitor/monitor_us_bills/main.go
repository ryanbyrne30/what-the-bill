package main

import (
	"context"
	"flag"
	"log"
	"time"

	proto_queue "github.com/ryanbyrne30/what-the-bill/monitor/queue/proto_queue"
)

var (
	queueAddr = flag.String("qaddr", "localhost:6051", "Address of queue")
	apiKey    = flag.String("apiKey", "", "API key for govinfo.gov")
	mongoUrl  = flag.String("mongoUrl", "mongo+srv://localhost:27017", "Mongo connection url")
	mongoDb   = flag.String("mongoDb", "monitor", "Database name for Mongo")
	mongoCol  = flag.String("mongoCol", "us_bills", "Collection name for Mongo")
)

func main() {
	flag.Parse()

	client := NewProtoClient(*queueAddr)
	defer client.Close()

	mongo := NewMongo(*mongoUrl, *mongoDb, *mongoCol)
	fetch := NewFetch(*apiKey, mongo)

	since := time.Now().Add(-24 * time.Hour)
	packages, err := fetch.BillsSince(since, 1000, 0)
	if err != nil {
		log.Fatalf("error fetching bills: %v", err)
	}

	ctx := context.Background()

	for _, pack := range packages {
		HandlePack(ctx, mongo, client, pack)
	}

	log.Printf("Ok")
}

func HandlePack(ctx context.Context, mongo *Mongo, client *ProtoClient, pack Package) {
	bill, err := mongo.GetBillById(ctx, pack.PackageID)
	if err != nil {
		log.Printf("Error getting bill from Mongo: %s", pack.PackageID)
		return
	}

	updatedAt, err := time.Parse("2006-01-02T15:04:05Z", pack.UpdatedAt)
	if err != nil {
		log.Printf("Unexpected date format: %s for bill %s", pack.UpdatedAt, pack.PackageID)
		return
	}
	isNew := bill == nil
	isUpdated := !isNew && updatedAt.Compare(bill.UpdateAt) > 0

	if isNew {
		log.Printf("New bill: %s", pack.PackageID)
		err = mongo.SaveBill(ctx, &Bill{
			BillID:   pack.PackageID,
			UpdateAt: updatedAt,
			Link:     pack.PackageLink,
		})
		if err == nil {
			err = client.PostUSBillCreatedEvent(ctx, &proto_queue.USBillCreatedEvent{
				PackageId:   pack.PackageID,
				UpdatedAt:   pack.UpdatedAt,
				PackageLink: pack.PackageLink,
			})
		}
	} else if isUpdated {
		log.Printf("Updating bill: %s", pack.PackageID)
		bill.UpdateAt = updatedAt
		err = mongo.UpdateBill(ctx, bill)
		if err == nil {
			err = client.PostUSBillUpdatedEvent(ctx, &proto_queue.USBillUpdatedEvent{
				PackageId:   pack.PackageID,
				UpdatedAt:   pack.UpdatedAt,
				PackageLink: pack.PackageLink,
			})
		}
	}

	if err != nil {
		log.Printf("error sending event: %v", err)
	}
}
