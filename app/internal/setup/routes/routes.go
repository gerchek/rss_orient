package routes

import (
	"log"
	"net/http"
	fetchPostsConstructor "rss/internal/services/fetchPosts/constructor"
	rssLinksConstructor "rss/internal/services/rssLinks/constructor"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type MyRoute struct {
}

func (rr *MyRoute) Routes() {
	r := mux.NewRouter().StrictSlash(true)
	// POSTS
	r.HandleFunc("/posts", fetchPostsConstructor.FetchPostsController.GetAllPosts).Methods("GET")
	// LINKS
	r.HandleFunc("/links/all", rssLinksConstructor.RssLinksController.LinkAll).Methods("GET")
	r.HandleFunc("/link/create", rssLinksConstructor.RssLinksController.LinkCreate).Methods("POST")
	r.HandleFunc("/link/delete/{id}", rssLinksConstructor.RssLinksController.LinkDelete).Methods("DELETE")
	r.HandleFunc("/link/update/{id}", rssLinksConstructor.RssLinksController.LinkUpdate).Methods("PUT")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "Delete"})

	log.Fatal(http.ListenAndServe("95.85.124.41:8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
