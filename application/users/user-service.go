package users

import (
	"context"
	"go-ddd-example/application/users/mappers"
	"go-ddd-example/application/users/models"
	"go-ddd-example/domain/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	UserService interface {
		AddNewUser(ctx context.Context, newUserModel *models.NewUserModel) (*models.NewUserModel, error)
		AddNewAdminUser(ctx context.Context, newUserModel *models.NewUserModel) (*models.NewUserModel, error)
		AddNewGuestUser(ctx context.Context) (*models.NewUserModel, error)
		GetUserById(ctx context.Context, id string) (*models.NewUserModel, error)
		AuthUser(ctx context.Context, username, password string) (bool, error)
	}
	userService struct {
		Repository users.IUserRepository
	}
)

func NewUserService(repository users.IUserRepository) UserService {
	return &userService{Repository: repository}
}

func (service userService) GetUserById(ctx context.Context, id string) (*models.NewUserModel, error) {

	var (
		user     *users.User
		objectId primitive.ObjectID
		err      error
	)

	if objectId, err = primitive.ObjectIDFromHex(id); err != nil {
		return nil, err
	}

	if user, err = service.Repository.FindOneById(ctx, objectId); err != nil {
		return nil, err
	}

	return mappers.MapNewUserModel(user), nil
}

func (service userService) AddNewUser(ctx context.Context, newUserModel *models.NewUserModel) (*models.NewUserModel, error) {

	user := users.NewUser(newUserModel.FirstName, newUserModel.LastName, newUserModel.UserName, newUserModel.Password)

	if err := service.Repository.Add(ctx, user); err != nil {
		return nil, err
	}

	return mappers.MapNewUserModel(user), nil
}

func (service userService) AddNewAdminUser(ctx context.Context, newUserModel *models.NewUserModel) (*models.NewUserModel, error) {

	user := users.NewAdminUser(newUserModel.FirstName, newUserModel.LastName, newUserModel.UserName, newUserModel.Password)

	if err := service.Repository.Add(ctx, user); err != nil {
		return nil, err
	}

	return mappers.MapNewUserModel(user), nil

}

func (service userService) AddNewGuestUser(ctx context.Context) (*models.NewUserModel, error) {
	user := users.NewGuestUser()

	if err := service.Repository.Add(ctx, user); err != nil {
		return nil, err
	}

	return mappers.MapNewUserModel(user), nil
}

func (service userService) AuthUser(ctx context.Context, username, password string) (bool, error) {

	var (
		user *users.User
		err  error
	)

	if user, err = service.Repository.FindOneByUsername(ctx, username); err != nil {
		return false, err
	}

	return user.EncryptedPassword.VerifyPassword(password), nil
}
