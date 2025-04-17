package middleware

import (
	"context"
	middleware "earnforglance/server/domain/middleware"
	"time"
)

type MiddlewareUsecase struct {
	MiddlewareRepository middleware.MiddlewareRepository
	contextTimeout       time.Duration
}

func NewMiddlewareUsecase(MiddlewareRepository middleware.MiddlewareRepository, timeout time.Duration) middleware.MiddlewareUsecase {
	return &MiddlewareUsecase{
		MiddlewareRepository: MiddlewareRepository,
		contextTimeout:       timeout,
	}
}

func (tu *MiddlewareUsecase) GetPermissionsCustumer(c context.Context, customerID string) ([]middleware.Middleware, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.MiddlewareRepository.GetPermissionsCustumer(ctx, customerID)
}
