package model

type User struct {
	Email string `json:"email"`
	Tasks []Task `json:"tasks"`
}
