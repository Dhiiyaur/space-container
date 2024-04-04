package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	http.ListenAndServe(port, r)
}
