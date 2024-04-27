package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Bill struct {
	ID       primitive.ObjectID `bson:"_id"`
	BillID   string             `bson:"bill_id"`
	UpdateAt time.Time          `bson:"updated_at"`
	Link     string             `bson:"link"`
}

type Mongo struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewMongo(url, db, col string) *Mongo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalf("could not connect to Mongo: %v", err)
	}
	log.Printf("successfully connected to Mongo")

	coll := client.Database(db).Collection(col)
	return &Mongo{client: client, col: coll}
}

func (m *Mongo) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *Mongo) GetBillById(ctx context.Context, id string) (*Bill, error) {
	var bill Bill
	err := m.col.FindOne(ctx, bson.D{{Key: "bill_id", Value: id}}).Decode(&bill)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &bill, nil
}

func (m *Mongo) SaveBill(ctx context.Context, bill *Bill) error {
	bill.ID = primitive.NewObjectID()
	_, err := m.col.InsertOne(ctx, bill)
	return err
}

func (m *Mongo) UpdateBill(ctx context.Context, bill *Bill) error {
	_, err := m.col.UpdateOne(ctx, bson.D{{Key: "bill_id", Value: bill.BillID}}, bson.D{{Key: "$set", Value: bill}})
	return err
}
