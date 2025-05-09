package domain

import (
	"context"
	customers "earnforglance/server/domain/customers"
)

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(c context.Context, user *customers.Customer) error
	CreateAccessToken(user *customers.Customer, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *customers.Customer, secret string, expiry int) (refreshToken string, err error)
}
