package get_bills

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ryanbyrne30/what-the-bill/server/templates/layouts"
)

type BillAction struct {
	Text string    `json:"text"`
	Date time.Time `json:"date"`
}

type Bill struct {
	ID         string       `bson:"_id"`
	Issued     time.Time    `bson:"issued"`
	Updated    time.Time    `bson:"updated"`
	Url        string       `bson:"url"`
	Title      string       `bson:"title"`
	ShortTitle string       `bson:"short_title"`
	Actions    []BillAction `bson:"actions"`
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
			w.Write([]byte("An error has occured"))
			return
		}

		log.Printf("Number of bills found: %d\n", len(bills))
		for _, b := range bills {
			log.Println(b.ShortTitle)
		}

		renderPage := Page(bills)
		component := layouts.Layout(renderPage)
		component.Render(ctx, w)
	}

}
