package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Passowrd string `json:"password"`
}
