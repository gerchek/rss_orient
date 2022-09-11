package routes

import (
	"log"
	"net/http"
	"rss/internal/services/rss/constructor"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

type MyRoute struct {
}

func (rr *MyRoute) Routes() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/posts", constructor.RssController.GetAll)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	
	log.Fatal(http.ListenAndServe("95.85.124.41:8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
