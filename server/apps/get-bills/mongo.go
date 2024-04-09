package get_bills

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	col *mongo.Collection
}

func NewMongoDB(col *mongo.Collection) *MongoDB {
	return &MongoDB{col: col}
}

func (m *MongoDB) GetBills(ctx context.Context, opts *GetBillsOpts) ([]Bill, error) {
	o := options.Find().SetSort(bson.D{{Key: "updated", Value: -1}})
	if opts.Limit > 0 {
		o.SetLimit(int64(opts.Limit))
	}
	if opts.Offset > 0 {
		o.SetSkip(int64(opts.Offset))
	}

	cursor, err := m.col.Find(ctx, bson.D{{}}, o)
	if err != nil {
		return nil, err
	}

	var results []Bill
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		err = cursor.Decode(&result)
		if err != nil {
			log.Printf("Could not decode bill from MongoDB. %s\n", err.Error())
		}
	}

	return results, nil
}
