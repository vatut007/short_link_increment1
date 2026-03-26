package main

import (
	"net/http"
	"shortener/handlers"
)

func main() {
	handlers.RegisterHandlers()
	router := handlers.Mux
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
