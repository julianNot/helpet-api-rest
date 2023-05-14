package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/services"
)

func setRouterAttendant(apiRouter *mux.Router) {
		apiRouter.HandleFunc("/attendants", services.GetAttendantsHandler).Methods("GET")
		apiRouter.HandleFunc("/attendants/{id}", services.GetAttendantHandler).Methods("GET")
		apiRouter.HandleFunc("/attendants", services.CreateAttendantHandler).Methods("POST")
}