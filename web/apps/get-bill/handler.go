package get_bill

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ryanbyrne30/what-the-bill/server/templates/layouts"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BillAction struct {
	Text string    `json:"text"`
	Date time.Time `json:"date"`
}

type Bill struct {
	ID         primitive.ObjectID `bson:"_id"`
	Issued     time.Time          `bson:"issued"`
	Updated    time.Time          `bson:"updated"`
	Url        string             `bson:"url"`
	Title      string             `bson:"title"`
	Text       string             `bson:"text"`
	ShortTitle string             `bson:"short_title"`
	Actions    []BillAction       `bson:"actions"`
}

type Database interface {
	GetBill(context.Context, string) (*Bill, error)
}

func Handler(ctx context.Context, db Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		bill, err := db.GetBill(r.Context(), id)
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
			w.Write([]byte("An error has occurred"))
			return
		}

		renderPage := Page(bill)
		component := layouts.Layout(renderPage)
		component.Render(ctx, w)
	}
}
