package users

import (
	"go-ddd-example/api/configs"
	appUsers "go-ddd-example/application/users"
	domainUsers "go-ddd-example/domain/users"
	common_di "go-ddd-example/infrastructure/common"
	"go-ddd-example/infrastructure/common/persistence"
)

func NewUserRepositoryResolve(config configs.Config) domainUsers.IUserRepository {
	return newUserRepository(persistence.NewMongoDb(config.User.MongoDb, config.User.Database), common_di.NewEventHandlerResolve())
}

func NewUserServiceResolve(config configs.Config) appUsers.UserService {
	return appUsers.NewUserService(NewUserRepositoryResolve(config))
}
