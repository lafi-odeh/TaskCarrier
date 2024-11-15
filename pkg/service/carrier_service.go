package service

import (
	"TaskCarrier/pkg/model"
	"TaskCarrier/pkg/repository"
)

func StoreUserTasks(email string, tasks []model.Task) {
	repository.StoreUserTasks(email, tasks)
}

func GetUserTasks(email string) ([]model.Task, error) {
	return repository.GetUserTasks(email)
}
