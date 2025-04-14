package domain

import (
	"context"

	domain "earnforglance/server/domain/customers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "customers"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

type UserRepository interface {
	GetByEmail(c context.Context, email string) (domain.Customer, error)
	GetPasw(c context.Context, email string) (domain.CustomerPassword, error)
	GetByID(c context.Context, id string) (User, error)
}
