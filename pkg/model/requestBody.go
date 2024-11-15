package model

type RequestBody struct {
	Email string `json:"email"`
	Tasks []Task `json:"tasks"`
}
