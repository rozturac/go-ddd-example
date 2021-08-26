package users

import (
	"context"
	"go-ddd-example/domain/common"
	"go-ddd-example/domain/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const _collectionName = "Users"

type userRepository struct {
	db              *mongo.Database
	eventDispatcher common.IEventDispatcher
}

func newUserRepository(db *mongo.Database, eventDispatcher common.IEventDispatcher) users.IUserRepository {
	return &userRepository{db: db, eventDispatcher: eventDispatcher}
}

func (repository userRepository) FindOneById(ctx context.Context, id primitive.ObjectID) (*users.User, error) {
	var user *users.User
	err := repository.db.Collection(_collectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (repository userRepository) FindOneByUsername(ctx context.Context, username string) (*users.User, error) {
	var user *users.User
	err := repository.db.Collection(_collectionName).FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return user, err
}

func (repository userRepository) Add(ctx context.Context, user *users.User) error {
	_, err := repository.db.Collection(_collectionName).InsertOne(ctx, &user, options.InsertOne())
	repository.dispatchDomainEvents(user, err)

	return err
}

func (repository userRepository) Update(ctx context.Context, user *users.User) error {
	_, err := repository.db.Collection(_collectionName).ReplaceOne(ctx, bson.M{"_id": user.Id}, &user)
	repository.dispatchDomainEvents(user, err)

	return err
}

func (repository userRepository) dispatchDomainEvents(user *users.User, err error) {
	if err == nil {
		repository.eventDispatcher.Dispatch(user.GetDomainEvents())
		user.ClearDomainEvents()
	}
}
