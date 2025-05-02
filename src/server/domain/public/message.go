package domain

import (
	"context"
)

type ContactUsRequest struct {
	Email     string
	Subject   string
	Enquiry   string
	FullName  string
	StoreID   []string
	News      bool
	IpAddress string
	Lang      string
}

type NewsLetterRequest struct {
	Email     string
	StoreID   []string
	IpAddress string
	Lang      string
}

type NewsLetterResponse struct {
	Result  bool
	Message string
}

type NewsLetterRepository interface {
	NewsLetterSubscription(c context.Context, news NewsLetterRequest, IpAdress string) (NewsLetterResponse, error)
	NewsLetterUnSubscribe(c context.Context, news NewsLetterRequest) (NewsLetterResponse, error)
	NewsLetterActivation(c context.Context, Guid string) (NewsLetterResponse, error)
	NewsLetterInactivate(c context.Context, Guid string) (NewsLetterResponse, error)
	ContactUs(c context.Context, contact ContactUsRequest) (NewsLetterResponse, error)
	GetSlugs(c context.Context, record string) ([]string, error)
}

type NewsLetterUsecase interface {
	NewsLetterSubscription(c context.Context, filter NewsLetterRequest, IpAdress string) (NewsLetterResponse, error)
	NewsLetterUnSubscribe(c context.Context, news NewsLetterRequest) (NewsLetterResponse, error)
	NewsLetterActivation(c context.Context, Guid string) (NewsLetterResponse, error)
	NewsLetterInactivate(c context.Context, Guid string) (NewsLetterResponse, error)
	ContactUs(c context.Context, contact ContactUsRequest) (NewsLetterResponse, error)
	GetSlugs(c context.Context, record string) ([]string, error)
}
