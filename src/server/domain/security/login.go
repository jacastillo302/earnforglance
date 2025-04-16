package domain

import (
	"context"
	settings "earnforglance/server/domain/configuration"
	domain "earnforglance/server/domain/customers"
	localization "earnforglance/server/domain/localization"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

const (
	CollectionUser = "customers"
)

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (domain.Customer, error)
	GetByUserName(c context.Context, email string) (domain.Customer, error)
	GetPasw(c context.Context, email string) (domain.CustomerPassword, error)
	GetSettingByName(c context.Context, name string) (settings.Setting, error)
	GetLocalebyName(c context.Context, name string, languageID string) (localization.LocaleStringResource, error)
	CreateAccessToken(user *domain.Customer, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *domain.Customer, secret string, expiry int) (refreshToken string, err error)
}

type LoginRepository interface {
	GetByEmail(c context.Context, email string) (domain.Customer, error)
	GetByUserName(c context.Context, email string) (domain.Customer, error)
	GetSettingByName(c context.Context, name string) (settings.Setting, error)
	GetPasw(c context.Context, email string) (domain.CustomerPassword, error)
	GetLocalebyName(c context.Context, name string, languageID string) (localization.LocaleStringResource, error)
	GetByID(c context.Context, id string) (domain.Customer, error)
}
