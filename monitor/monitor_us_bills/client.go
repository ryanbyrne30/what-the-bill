package main

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	proto_queue "github.com/ryanbyrne30/what-the-bill/monitor/queue/proto_queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProtoClient struct {
	conn   *grpc.ClientConn
	client proto_queue.QueueClient
}

func NewProtoClient(addr string) *ProtoClient {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := proto_queue.NewQueueClient(conn)

	return &ProtoClient{
		conn:   conn,
		client: c,
	}
}

func (p *ProtoClient) Close() {
	p.conn.Close()
}

func (p *ProtoClient) PostUSBillUpdatedEvent(ctx context.Context, event *proto_queue.USBillUpdatedEvent) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	result, err := p.client.PostUSBillUpdatedEvent(ctx, event)
	if err != nil {
		return err
	}
	if strings.ToLower(result.Status) != "ok" {
		return errors.New(result.Status)
	}

	return nil
}

func (p *ProtoClient) PostUSBillCreatedEvent(ctx context.Context, event *proto_queue.USBillCreatedEvent) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	result, err := p.client.PostUSBillCreatedEvent(ctx, event)
	if err != nil {
		return err
	}
	if strings.ToLower(result.Status) != "ok" {
		return errors.New(result.Status)
	}

	return nil
}
