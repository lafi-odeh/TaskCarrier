package repository

import (
	"TaskCarrier/pkg/model"
	"errors"
	"log"
)

var userDatabase = make(map[string][]model.Task)

func StoreUserTasks(email string, tasks []model.Task) {
	userDatabase[email] = tasks
	log.Println(userDatabase)
}

func GetUserTasks(email string) ([]model.Task, error) {
	tasks, exists := userDatabase[email]
	if !exists {
		return nil, errors.New("tasks not found for this email")
	}
	return tasks, nil
}
