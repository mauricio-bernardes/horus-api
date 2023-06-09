package main

import (
	"fmt"
	"horus-api/configs"
	"horus-api/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/subscribe", handlers.HandleSubscribe)
	r.Post("/add", handlers.HandleInserService)
	r.Get("/list", handlers.HandleListServices)
	r.Delete("/remove", handlers.HandleRemoveService)
	r.Delete("/unsubscribe", handlers.HandleUnsubscribe)
	r.Get("/health", handlers.HandleHealthCheck)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetAPIConfig().Port), r)
}
