package users

import (
	"go-ddd-example/domain/common"
)

type User struct {
	Id                string             `json:"id"`
	FirstName         string             `json:"first_name"`
	LastName          string             `json:"last_name"`
	UserName          string             `json:"user_name"`
	EncryptedPassword *EncryptedPassword `json:"encrypted_password"`
	Roles             []*UserRole        `json:"roles"`
}

func NewUser(firstName, lastName, username, password string) *User {

	var user *User

	if len(username) == 0 {
		panic(common.IsNullOrEmptyError("username"))
	}

	user = &User{
		FirstName:         firstName,
		LastName:          lastName,
		UserName:          username,
		EncryptedPassword: NewEncryptedPassword(password),
	}

	return user
}

func NewGuestUser() *User {

	user := NewUser("", "", "Guest", "12345")
	user.AddUserRole(UserRole_Guest)

	return user
}

func NewAdminUser(firstName, lastName, username, password string) *User {

	user := NewUser(firstName, lastName, username, password)
	user.AddUserRole(UserRole_Admin)

	return user
}

func (u *User) AddUserRole(role *UserRole) {

	if role == nil {
		panic(common.IsNullOrEmptyError("role"))
	}

	for _, roleItem := range u.Roles {
		if roleItem.Name == role.Name {
			panic(common.AlreadyExistRoleError(role.Name))
		}
	}

	u.Roles = append(u.Roles, role)
}

func (u *User) ChangePassword(oldPassword, newPassword string) {

	if !u.EncryptedPassword.VerifyPassword(oldPassword) {
		panic("")
	}

	u.EncryptedPassword = NewEncryptedPassword(newPassword)
}
