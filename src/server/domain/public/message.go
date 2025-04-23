package domain

import (
	"context"
)

type NewsLetterRequest struct {
	Email     string
	StoreID   string
	IpAddress string
	Lang      string
}

type NewsLetterResponse struct {
	Result  bool
	Message string
}

type NewsLetterRepository interface {
	NewsLetterSubscription(c context.Context, filter NewsLetterRequest, IpAdress string) (NewsLetterResponse, error)
}

type NewsLetterUsecase interface {
	NewsLetterSubscription(c context.Context, filter NewsLetterRequest, IpAdress string) (NewsLetterResponse, error)
}
