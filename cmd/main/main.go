package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Metarock/go-bookstore/pkg/routes"
)

// specify where our routers are
func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
