package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/messages"
)

type emailaccountUsecase struct {
	emailaccountRepository domain.EmailAccountRepository
	contextTimeout         time.Duration
}

func NewEmailAccountUsecase(emailaccountRepository domain.EmailAccountRepository, timeout time.Duration) domain.EmailAccountUsecase {
	return &emailaccountUsecase{
		emailaccountRepository: emailaccountRepository,
		contextTimeout:         timeout,
	}
}

func (tu *emailaccountUsecase) CreateMany(c context.Context, items []domain.EmailAccount) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.emailaccountRepository.CreateMany(ctx, items)
}

func (tu *emailaccountUsecase) Create(c context.Context, emailaccount *domain.EmailAccount) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.emailaccountRepository.Create(ctx, emailaccount)
}

func (tu *emailaccountUsecase) Update(c context.Context, emailaccount *domain.EmailAccount) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.emailaccountRepository.Update(ctx, emailaccount)
}

func (tu *emailaccountUsecase) Delete(c context.Context, emailaccount string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.emailaccountRepository.Delete(ctx, emailaccount)
}

func (lu *emailaccountUsecase) FetchByID(c context.Context, emailaccountID string) (domain.EmailAccount, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.emailaccountRepository.FetchByID(ctx, emailaccountID)
}

func (lu *emailaccountUsecase) Fetch(c context.Context) ([]domain.EmailAccount, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.emailaccountRepository.Fetch(ctx)
}
