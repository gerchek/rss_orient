package routes

import (
	"log"
	"net/http"
	"rss/internal/services/rss/constructor"

	"github.com/gorilla/mux"
)

type MyRoute struct {
}

func (rr *MyRoute) Routes() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/posts", constructor.RssController.GetAll)
	log.Fatal(http.ListenAndServe("95.85.124.41:8080", r))
}
