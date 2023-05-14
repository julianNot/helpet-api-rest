package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/services"
)

func setRouterPet(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/pets", services.GetPetsHandler).Methods("GET")
	apiRouter.HandleFunc("/pets/{id}", services.GetPetHandler).Methods("GET")
	apiRouter.HandleFunc("/pets", services.CreatePetHandler).Methods("POST")
}
