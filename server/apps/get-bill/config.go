package get_bill

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewGetBillHandler(ctx context.Context, col *mongo.Collection) http.HandlerFunc {
	db := NewMongoDB(col)
	return Handler(ctx, db)
}
