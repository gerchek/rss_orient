package routes

import (
	"log"
	"net/http"
	"rss/internal/services/rss/constructor"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type MyRoute struct {
}

func (rr *MyRoute) Routes() {
	r := mux.NewRouter().StrictSlash(true)
	// POSTS
	r.HandleFunc("/posts", constructor.RssController.GetAllPosts).Methods("GET")
	// LINKS
	r.HandleFunc("/links/all", constructor.RssController.LinkAll).Methods("GET")
	r.HandleFunc("/link/create", constructor.RssController.LinkCreate).Methods("POST")
	r.HandleFunc("/link/delete/{id}", constructor.RssController.LinkDelete).Methods("DELETE")
	r.HandleFunc("/link/update/{id}", constructor.RssController.LinkUpdate).Methods("PUT")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "Delete"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
