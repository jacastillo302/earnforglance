package domain

import (
	"context"
)

type Middleware struct {
	UserID        string
	Authorization bool
}

type MiddlewareUsecase interface {
	GetPermissionsCustumer(c context.Context, customerID string) ([]Middleware, error)
}

type MiddlewareRepository interface {
	GetPermissionsCustumer(c context.Context, customerID string) ([]Middleware, error)
}
