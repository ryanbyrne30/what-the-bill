package get_bills

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ryanbyrne30/what-the-bill/server/templates/layouts"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bill struct {
	ID         primitive.ObjectID `bson:"_id"`
	Updated    time.Time          `bson:"updated"`
	ShortTitle string             `bson:"short_title"`
}

type GetBillsOpts struct {
	Limit  int
	Offset int
}

type Database interface {
	GetBills(context.Context, *GetBillsOpts) ([]Bill, error)
}

func Handler(ctx context.Context, db Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryPage := r.URL.Query().Get("page")
		page, _ := strconv.Atoi(queryPage)
		limit := 16

		bills, err := db.GetBills(r.Context(), &GetBillsOpts{
			Limit:  limit,
			Offset: page * limit,
		})
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
			w.Write([]byte("An error has occurred"))
			return
		}

		renderPage := Page(bills)
		component := layouts.Layout(renderPage)
		component.Render(ctx, w)
	}

}
