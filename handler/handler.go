package handler

import (
	"main/api"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlers() {
	r := mux.NewRouter()

	r.HandleFunc("/ticket_options", api.Ticket_options).Methods("POST")
	r.HandleFunc("/ticket/{id}", api.Ticket).Methods("GET")
	r.HandleFunc("/ticket_options/{id}/purchases", api.Ticket_purchases).Methods("POST")
	r.HandleFunc("/users", api.Users).Methods("POST")

	http.ListenAndServe(":8080", r)
}
