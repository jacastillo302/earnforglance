package domain

import (
	"context"
	domain "earnforglance/server/domain/customers"
)

type SingInRequest struct {
	Customer   domain.Customer
	Password   string
	Attributes map[string]string
	News       bool
	IpAddress  string
	Lang       string
}

type SingInResponse struct {
	Result  bool
	Message string
}

type CustomerRepository interface {
	SingIn(c context.Context, sigin SingInRequest) (SingInResponse, error)
	GetSlugs(c context.Context, record string) ([]string, error)
}

type CustomerUsecase interface {
	SingIn(c context.Context, sigin SingInRequest) (SingInResponse, error)
	GetSlugs(c context.Context, record string) ([]string, error)
}
