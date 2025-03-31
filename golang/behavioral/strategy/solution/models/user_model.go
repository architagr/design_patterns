package models

import "fmt"

type UserModel struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *UserModel) GetHeader() string {
	return "id,name,email\n"
}

func (u *UserModel) GetRow() string {
	return fmt.Sprintf("%d,%s,%s\n", u.Id, u.Name, u.Email)
}
