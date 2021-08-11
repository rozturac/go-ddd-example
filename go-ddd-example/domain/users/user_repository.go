package users

type UserRepository interface {
	FindOne(id string) (*User, error)
	Add(user *User) error
	Update(user *User) error
}