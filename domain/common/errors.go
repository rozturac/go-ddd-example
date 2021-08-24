package common

import (
	"errors"
	"fmt"
)

const PREFIX = "Domain - "
const (
	IsNullOrEmptyText    string = PREFIX + "%s is null or empty !"
	AlreadyExistRoleText string = PREFIX + "The User already has '%s' role !"
)

func IsNullOrEmptyError(name string) error {
	return errors.New(fmt.Sprintf(IsNullOrEmptyText, name))
}

func AlreadyExistRoleError(name string) error {
	return errors.New(fmt.Sprintf(AlreadyExistRoleText, name))
}
