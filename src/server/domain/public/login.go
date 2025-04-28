package domain

import (
	"context"
	settings "earnforglance/server/domain/configuration"
	domain "earnforglance/server/domain/customers"
	localization "earnforglance/server/domain/localization"
	security "earnforglance/server/domain/security"

	"go.mongodb.org/mongo-driver/v2/bson"
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
	CreateAccessToken(user *domain.Customer, slugs []security.UrlRecord, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *domain.Customer, slugs []security.UrlRecord, secret string, expiry int) (refreshToken string, err error)
	AddActivityLog(c context.Context, customerID bson.ObjectID, systemKeyword string, comment string, ipAddress string) (bool, error)
}

type LoginRepository interface {
	GetByEmail(c context.Context, email string) (domain.Customer, error)
	GetByUserName(c context.Context, email string) (domain.Customer, error)
	GetSettingByName(c context.Context, name string) (settings.Setting, error)
	GetPasw(c context.Context, email string) (domain.CustomerPassword, error)
	GetLangugaByCode(c context.Context, lang string) (localization.Language, error)
	GetLocalebyName(c context.Context, name string, languageID string) (localization.LocaleStringResource, error)
	GetByID(c context.Context, id string) (domain.Customer, error)
	AddActivityLog(c context.Context, customerID bson.ObjectID, systemKeyword string, comment string, ipAddress string) (bool, error)
}
