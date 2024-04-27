package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ryanbyrne30/what-the-bill/monitor/queue/proto_consumer_us_bills"
	"github.com/ryanbyrne30/what-the-bill/monitor/queue/queue/mongo"
)

func DispatchUSBillCreatedEvent(ctx context.Context, client *ProtoClient, event mongo.Event) error {
	log.Printf("dispatching US bill created event: %s", event.ID.Hex())
	var parsed struct {
		PackageID   string `json:"PackageID"`
		PackageLink string `json:"PackageLink"`
	}
	if err := json.Unmarshal([]byte(event.Message), &parsed); err != nil {
		return err
	}

	return client.PostUSBillCreatedEvent(ctx, &proto_consumer_us_bills.USBillCreatedEvent{
		PackageId:   parsed.PackageID,
		PackageLink: parsed.PackageLink,
	})
}

func DispatchUSBillUpdatedEvent(ctx context.Context, client *ProtoClient, event mongo.Event) error {
	log.Printf("dispatching US bill updated event: %s", event.ID.Hex())
	var e proto_consumer_us_bills.USBillUpdatedEvent
	if err := json.Unmarshal([]byte(event.Message), &e); err != nil {
		return err
	}

	return client.PostUSBillUpdatedEvent(ctx, &e)
}
