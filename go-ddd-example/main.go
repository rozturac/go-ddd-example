package main

import (
	"fmt"
	"go-ddd-example/application/users"
	"go-ddd-example/application/users/models"
	infUsers "go-ddd-example/infrastructure/users"
)

func main() {

	var userRepository = infUsers.NewUserRepository()
	var userService = users.NewUserService(userRepository)

	user, err := userService.AddNewUser(&models.NewUserModel{
		FirstName: "Rıdvan",
		LastName: "ÖZTURAÇ",
		UserName: "rozturac",
		Password: "123456",
	})

	if err != nil {
		panic(err)
	}

	if user.EncryptedPassword.VerifyPassword("123456") {
		fmt.Println("123456 true password.")
	}

	fmt.Println(user)
}
