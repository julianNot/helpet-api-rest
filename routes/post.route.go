package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/services"
)

func setRouterPost(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/posts", services.GetAllPostsHandler).Methods("GET")
	apiRouter.HandleFunc("/posts/{id}", services.GetPostHandler).Methods("GET")
	apiRouter.HandleFunc("/posts", services.CreatePostHandler).Methods("POST")
}
