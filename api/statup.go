package api

import (
	context2 "context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-ddd-example/api/configs"
	"go-ddd-example/api/controllers/v1"
	"go-ddd-example/application/users/consumers"
	common_di "go-ddd-example/infrastructure/common"
	infUsers "go-ddd-example/infrastructure/users"
	"net/http"
	"os"
)

func Init() {

	var (
		config configs.Config
		err    error
	)

	if config, err = configs.LoadConfig("./api", os.Getenv("ENV")); err != nil {
		panic(err)
	}

	var userService = infUsers.NewUserServiceResolve(config)

	BindConsumers(config)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		return userService.AuthUser(context2.Background(), username, password)
	}))

	v1 := e.Group("/api/v1")
	controllers_v1.CreateGuestUser(v1, userService)
	controllers_v1.GetUserByObjectId(v1, userService)

	e.Start(":8080")
}

func BindConsumers(config configs.Config) {
	rbt := common_di.NewRabbitMQResolve(config)
	rbt.BindConsumer(consumers.NewUserCreatedConsumer())
	rbt.Start()
}
