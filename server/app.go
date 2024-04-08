package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ryanbyrne30/what-the-bill/server/templates/layouts"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		component := layouts.Layout()
		component.Render(context.Background(), w)
	})
	http.ListenAndServe(":3000", r)
}
