package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserRepository interface {
	FindOneById(ctx context.Context, id primitive.ObjectID) (*User, error)
	FindOneByUsername(ctx context.Context, username string) (*User, error)
	Add(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}
