package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/db"
	"github.com/julianNot/helpet-api-rest/models"
)

func GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	db.DB.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	params := mux.Vars(r)
	db.DB.First(&post, params["id"])
	if post.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Post Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&post)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	createdPost := db.DB.Create(&post)
	err := createdPost.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&post)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	params := mux.Vars(r)
	db.DB.First(&post, params["id"])
	if post.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Post Not Found"))
		return
	}
	db.DB.Unscoped().Delete(&post)
	w.WriteHeader(http.StatusNoContent)
}