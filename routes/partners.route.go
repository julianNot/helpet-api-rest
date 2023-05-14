package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/services"
)

func setRouterPartner(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/partners", services.GetPartnersHandler).Methods("GET")
	apiRouter.HandleFunc("/partners/{id}", services.GetPartnerHandler).Methods("GET")
	apiRouter.HandleFunc("/partners", services.CreatePartnerHandler).Methods("POST")
	apiRouter.HandleFunc("/partners/{id}", services.UpdatePartnerHandler).Methods("PUT")
	apiRouter.HandleFunc("/partners/{id}", services.DeletePartnerHandler).Methods("DELETE")
}
