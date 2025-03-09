package models

type GetUserResponse struct {
	UserId    int    `json:"userId"`
	UserName  string `json:"userName"`
	AvatarUrl string `json:"avatarUrl"`
}
