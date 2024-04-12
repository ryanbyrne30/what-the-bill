package get_bill

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	col *mongo.Collection
}

func NewMongoDB(col *mongo.Collection) *MongoDB {
	return &MongoDB{col: col}
}

func (m *MongoDB) GetBill(ctx context.Context, id string) (*Bill, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var bill Bill
	err = m.col.FindOne(ctx, bson.D{{Key: "_id", Value: objID}}).Decode(&bill)
	if err != nil {
		return nil, err
	}

	return &bill, nil
}
