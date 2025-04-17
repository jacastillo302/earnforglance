package domain

import (
	"context"
	settings "earnforglance/server/domain/configuration"
	domain "earnforglance/server/domain/customers"
	localization "earnforglance/server/domain/localization"
)

type LoginRequest struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Lang     string `bson:"language"`
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
	GetLangugaByCode(c context.Context, lang string) (localization.Language, error)
	GetLocalebyName(c context.Context, name string, languageID string) (localization.LocaleStringResource, error)
	CreateAccessToken(user *domain.Customer, slugs []UrlRecord, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *domain.Customer, slugs []UrlRecord, secret string, expiry int) (refreshToken string, err error)
}

type LoginRepository interface {
	GetByEmail(c context.Context, email string) (domain.Customer, error)
	GetByUserName(c context.Context, email string) (domain.Customer, error)
	GetSettingByName(c context.Context, name string) (settings.Setting, error)
	GetPasw(c context.Context, email string) (domain.CustomerPassword, error)
	GetLangugaByCode(c context.Context, lang string) (localization.Language, error)
	GetLocalebyName(c context.Context, name string, languageID string) (localization.LocaleStringResource, error)
	GetByID(c context.Context, id string) (domain.Customer, error)
}
