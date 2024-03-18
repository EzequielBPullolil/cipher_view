package main

import (
	"log"
	"net/http"

	"github.com/EzequielBPullolil/cipher_view/src/endpoints"
	"github.com/go-chi/chi/v5"
	"github.com/kataras/blocks"
)

func main() {
	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
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
	r.Post("/verify_password", endpoints.VerifyPassword)

	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		if err := views.ExecuteTemplate(w, "home", "base", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	})
	http.ListenAndServe(":8000", r)
}
