package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/db"
	"github.com/julianNot/helpet-api-rest/routes"
)

func main() {
	db.SetupDatabase()

	router := mux.NewRouter()
	routes.SetRouter(router)

	http.ListenAndServe(":3000", router)
}
