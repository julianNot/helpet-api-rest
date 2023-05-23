package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/services"
)

func setRouterProfessional(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/professionals", services.GetProfessionalsHandler).Methods("GET")
	apiRouter.HandleFunc("/professionals/{id}", services.GetProfessionalHandler).Methods("GET")
	apiRouter.HandleFunc("/professionals", services.CreateProfessionalHandler).Methods("POST")
	apiRouter.HandleFunc("/professionals/{id}", services.UpdateProfessionalHandler).Methods("PUT")
	apiRouter.HandleFunc("/professionals/{id}", services.DeleteProfessionalHandler).Methods("DELETE")
}
