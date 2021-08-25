package users

import (
	"go-ddd-example/api/configs"
	appUsers "go-ddd-example/application/users"
	domainUsers "go-ddd-example/domain/users"
	common_di "go-ddd-example/infrastructure/common"
	"go-ddd-example/infrastructure/common/persistence"
)

func NewUserRepositoryResolve(config configs.Config) domainUsers.IUserRepository {
	rbt := common_di.NewRabbitMQResolve(config)
	eventHandler := common_di.NewEventHandlerResolve(rbt)
	return newUserRepository(persistence.NewMongoDb(config.User.MongoDb, config.User.Database), eventHandler)
}

func NewUserServiceResolve(config configs.Config) appUsers.UserService {
	return appUsers.NewUserService(NewUserRepositoryResolve(config))
}
