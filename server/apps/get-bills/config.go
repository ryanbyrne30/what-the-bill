package get_bills

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewGetBillsHandler(ctx context.Context, col *mongo.Collection) http.HandlerFunc {
	db := NewMongoDB(col)
	return Handler(ctx, db)
}
