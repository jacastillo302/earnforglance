package usecase

import (
	"context"
	"time"

	customers "earnforglance/server/domain/customers"
	domain "earnforglance/server/domain/public"
	security "earnforglance/server/domain/security"
	"earnforglance/server/internal/tokenutil"
)

type refreshTokenUsecase struct {
	userRepository domain.LoginRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.LoginRepository, timeout time.Duration) security.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(c context.Context, email string) (customers.Customer, error) {
	ctx, cancel := context.WithTimeout(c, rtu.contextTimeout)
	defer cancel()
	return rtu.userRepository.GetByID(ctx, email)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *customers.Customer, slugs []security.UrlRecord, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, slugs, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *customers.Customer, slugs []security.UrlRecord, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, slugs, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}
