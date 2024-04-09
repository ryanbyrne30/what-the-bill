package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	get_bills "github.com/ryanbyrne30/what-the-bill/server/apps/get-bills"
	"github.com/ryanbyrne30/what-the-bill/server/templates/layouts"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dbUrl := os.Getenv("MONGODB_URL")
	if dbUrl == "" {
		panic("MONGODB_URL not defined")
	}
	dbName := os.Getenv("MONGODB_DATABASE")
	if dbName == "" {
		panic("MONGODB_DATABASE not defined")
	}
	dbCol := os.Getenv("MONGODB_COLLECTION")
	if dbCol == "" {
		panic("MONGODB_COLLECTION not defined")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUrl))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(dbName).Collection(dbCol)

	fs := http.FileServer(http.Dir("static"))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Get("/bills", get_bills.NewGetBillsHandler(context.TODO(), coll))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		opts := &get_bills.GetBillsOpts{
			Limit: 10,
		}
		mongo := get_bills.NewMongoDB(coll)
		bills, err := mongo.GetBills(context.TODO(), opts)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		for _, bill := range bills {
			fmt.Println(bill.ShortTitle)
		}
		// var result []struct{}
		// opts := options.Find().SetLimit(10)
		// cur, err := coll.Find(context.TODO(), bson.D{{}}, opts)
		// if err != nil {
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }
		// err = cur.All(context.TODO(), &result)
		// if err != nil {
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }
		// log.Println("Number of docs", len(result))
		// for _, r := range result {
		// 	cur.Decode(&r)
		// 	output, err := json.MarshalIndent(r, "", "  ")
		// 	if err != nil {
		// 		continue
		// 	}
		// 	fmt.Printf("%s\n", output)
		// }

		component := layouts.Layout(nil)
		component.Render(context.Background(), w)
	})
	http.ListenAndServe(":3000", r)
}
