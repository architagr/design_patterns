package models

type UserModel struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
