package users

import (
	"go-ddd-example/application/users/models"
	"go-ddd-example/domain/users"
)

type (
	UserService interface {
		AddNewUser(newUserModel *models.NewUserModel) (*users.User, error)
		AddNewAdminUser(newUserModel *models.NewUserModel) (*users.User, error)
		AddNewGuestUser() (*users.User, error)
	}
	userService struct {
		Repository users.UserRepository
	}
)

func NewUserService(repository users.UserRepository) UserService {
	return &userService{Repository: repository}
}

func (u userService) AddNewUser(newUserModel *models.NewUserModel) (*users.User, error) {

	user := users.NewUser(newUserModel.FirstName, newUserModel.LastName, newUserModel.UserName, newUserModel.Password)

	if err := u.Repository.Add(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u userService) AddNewAdminUser(newUserModel *models.NewUserModel) (*users.User, error) {

	user := users.NewAdminUser(newUserModel.FirstName, newUserModel.LastName, newUserModel.UserName, newUserModel.Password)

	if err := u.Repository.Add(user); err != nil {
		return nil, err
	}

	return user, nil

}

func (u userService) AddNewGuestUser() (*users.User, error) {
	user := users.NewGuestUser()

	if err := u.Repository.Add(user); err != nil {
		return nil, err
	}

	return user, nil
}
