package api

import (
	"TaskCarrier/pkg/handler"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/sender", handler.HandleUserTasksPost).Methods("POST")
	r.HandleFunc("/requester", handler.HandleUserTasksGet).Methods("GET")
	return r
}
