package users

import "go-ddd-example/domain/users"

type userRepository struct {
}

func NewUserRepository() users.UserRepository {
	return &userRepository{}
}

func (u userRepository) FindOne(id string) (*users.User, error) {
	return &users.User{
		UserName: "Default",
	}, nil
}

func (u userRepository) Add(user *users.User) error {
	return nil
}

func (u userRepository) Update(user *users.User) error {
	return nil
}
