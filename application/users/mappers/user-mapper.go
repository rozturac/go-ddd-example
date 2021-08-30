package mappers

import (
	"go-ddd-example/application/users/models"
	"go-ddd-example/domain/users"
)

func MapNewUserModel(user *users.User) *models.NewUserModel {
	return &models.NewUserModel{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
	}
}
