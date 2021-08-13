package users

import (
	"go-ddd-example/api/configs"
	appUsers "go-ddd-example/application/users"
	domainUsers "go-ddd-example/domain/users"
	"go-ddd-example/infrastructure/common/persistence"
)

func NewUserRepositoryResolve(config configs.Config) domainUsers.UserRepository {
	return NewUserRepository(persistence.NewMongoDb(config.User.MongoDb, config.User.Database))
}

func NewUserServiceResolve(config configs.Config) appUsers.UserService {
	return appUsers.NewUserService(NewUserRepositoryResolve(config))
}