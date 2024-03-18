package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kataras/blocks"
)

func main() {
	r := chi.NewRouter()
	views := blocks.New("./src/views")
	if err := views.Load(); err != nil {
		log.Fatal(err)
	}
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if err := views.ExecuteTemplate(w, "index", "base", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
