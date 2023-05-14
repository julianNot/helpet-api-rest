package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/auth"
	"github.com/julianNot/helpet-api-rest/middlewares"
)

func SetRouter(router *mux.Router) {
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	
	// Login Partners
	apiRouter.HandleFunc("/auth/partners", auth.AuthenticatePartnerHandler).Methods("POST")

	//Middleware Auth
	apiRouter.Use(middlewares.AuthMiddlewarePartner)

	// Routes
	setRouterAppointment(apiRouter)
	setRouterAttendant(apiRouter)
	setRouterPartner(apiRouter)
	setRouterPet(apiRouter)
	setRouterPost(apiRouter)
	setRouterProfessional(apiRouter)
	setRouterVet(apiRouter)
}