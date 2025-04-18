package domain

import (
	"context"
	customers "earnforglance/server/domain/customers"
	// Removed to resolve import cycle
)

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUsecase interface {
	GetUserByID(c context.Context, id string) (customers.Customer, error)
	CreateAccessToken(user *customers.Customer, slugs []UrlRecord, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *customers.Customer, slugs []UrlRecord, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
}
