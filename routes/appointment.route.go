package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/services"
)

func setRouterAppointment(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/appointments", services.GetAppointmentsHandler).Methods("GET")
	apiRouter.HandleFunc("/appointments/{id}", services.GetAppointmentHandler).Methods("GET")
	apiRouter.HandleFunc("/appointments/attendants/{id}", services.GetAttendantsByAttendant).Methods("GET")
	apiRouter.HandleFunc("/appointments", services.CreateAppointmentHandler).Methods("POST")
}
