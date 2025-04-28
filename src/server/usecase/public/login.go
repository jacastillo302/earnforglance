package usecase

import (
	"context"
	"time"

	settings "earnforglance/server/domain/configuration"
	customers "earnforglance/server/domain/customers"
	localization "earnforglance/server/domain/localization"
	domain "earnforglance/server/domain/public"
	security "earnforglance/server/domain/security"
	"earnforglance/server/internal/tokenutil"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type loginUsecase struct {
	userRepository domain.LoginRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.LoginRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (customers.Customer, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) GetByUserName(c context.Context, usermame string) (customers.Customer, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByUserName(ctx, usermame)
}

func (lu *loginUsecase) GetPasw(c context.Context, CustumerID string) (customers.CustomerPassword, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetPasw(ctx, CustumerID)
}

func (lu *loginUsecase) GetSettingByName(c context.Context, name string) (settings.Setting, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetSettingByName(ctx, name)
}

func (lu *loginUsecase) GetLangugaByCode(c context.Context, lang string) (localization.Language, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetLangugaByCode(ctx, lang)
}

func (lu *loginUsecase) CreateAccessToken(user *customers.Customer, slugs []security.UrlRecord, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, slugs, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *customers.Customer, slugs []security.UrlRecord, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, slugs, secret, expiry)
}

func (lu *loginUsecase) GetLocalebyName(c context.Context, name string, languageID string) (localization.LocaleStringResource, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetLocalebyName(ctx, name, languageID)
}

func (lu *loginUsecase) AddActivityLog(c context.Context, customerID bson.ObjectID, systemKeyword string, comment string, ipAddress string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.AddActivityLog(ctx, customerID, systemKeyword, comment, ipAddress)
}
