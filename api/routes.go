package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages/stats", GetMessageStats).Methods("GET")
	return router
}
