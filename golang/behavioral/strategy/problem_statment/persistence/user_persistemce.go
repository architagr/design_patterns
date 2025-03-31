package persistence

import "strategy/models"

type UserPersistence struct {
}

func (usrPers *UserPersistence) ListUsers(paginationObj models.Pagination) []models.UserModel {
	return []models.UserModel{
		{
			Id:    1,
			Name:  "User1",
			Email: "user_1@foo.bar",
		},
		{
			Id:    2,
			Name:  "User2",
			Email: "user_2@foo.bar",
		},
	}
}
