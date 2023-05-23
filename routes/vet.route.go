package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/services"
)

func setRouterVet(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/vets", services.GetVetsHandler).Methods("GET")
	apiRouter.HandleFunc("/vets/{id}", services.GetVetHandler).Methods("GET")
	apiRouter.HandleFunc("/vets/Professionals/{id}", services.GetProfessinalsVet).Methods("GET")
	apiRouter.HandleFunc("/vets/posts/{id}", services.GetPostsVet).Methods("GET")
	apiRouter.HandleFunc("/vets", services.CreateVetHandler).Methods("POST")
}
