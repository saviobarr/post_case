package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saviobarr/post_case/controller"
)

//Start starts the entire app
func Start() {

	r := mux.NewRouter()
	r.HandleFunc("/post/{id}", controller.GetPost).Methods("GET")
	r.HandleFunc("/post", controller.CreatePost).Methods("POST")
	r.HandleFunc("/comments/{postId}", controller.CreatePost).Methods("GET")
	r.HandleFunc("/comment", controller.CreatePost).Methods("POST")

	http.Handle("/", r)

}
