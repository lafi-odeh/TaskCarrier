package handler

import (
	"TaskCarrier/pkg/model"
	"TaskCarrier/pkg/service"
	"encoding/json"
	"log"
	"net/http"
)

func HandleUserTasksPost(writer http.ResponseWriter, request *http.Request) {
	var reqBody model.RequestBody
	err := json.NewDecoder(request.Body).Decode(&reqBody)
	if err != nil {
		http.Error(writer, "Invalid data", http.StatusBadRequest)
		return
	}
	if reqBody.Email == "" {
		http.Error(writer, "Email is required", http.StatusBadRequest)
		return
	}
	log.Println(reqBody)
	service.StoreUserTasks(reqBody.Email, reqBody.Tasks)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(map[string]string{"message": "Tasks saved successfully"})
}

func HandleUserTasksGet(writer http.ResponseWriter, request *http.Request) {
	email := request.URL.Query().Get("email")
	if email == "" {
		http.Error(writer, "Email is required", http.StatusBadRequest)
		return
	}
	tasks, err := service.GetUserTasks(email)
	if err != nil {
		http.Error(writer, "Tasks not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(writer).Encode(tasks)
}
