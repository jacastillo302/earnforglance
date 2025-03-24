package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type backinstocksubscriptionUsecase struct {
	backinstocksubscriptionRepository domain.BackInStockSubscriptionRepository
	contextTimeout                    time.Duration
}

func NewBackInStockSubscriptionUsecase(backinstocksubscriptionRepository domain.BackInStockSubscriptionRepository, timeout time.Duration) domain.BackInStockSubscriptionUsecase {
	return &backinstocksubscriptionUsecase{
		backinstocksubscriptionRepository: backinstocksubscriptionRepository,
		contextTimeout:                    timeout,
	}
}

func (tu *backinstocksubscriptionUsecase) CreateMany(c context.Context, items []domain.BackInStockSubscription) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.backinstocksubscriptionRepository.CreateMany(ctx, items)
}

func (tu *backinstocksubscriptionUsecase) Create(c context.Context, backinstocksubscription *domain.BackInStockSubscription) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.backinstocksubscriptionRepository.Create(ctx, backinstocksubscription)
}

func (tu *backinstocksubscriptionUsecase) Update(c context.Context, backinstocksubscription *domain.BackInStockSubscription) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.backinstocksubscriptionRepository.Update(ctx, backinstocksubscription)
}

func (tu *backinstocksubscriptionUsecase) Delete(c context.Context, backinstocksubscription string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.backinstocksubscriptionRepository.Delete(ctx, backinstocksubscription)
}

func (lu *backinstocksubscriptionUsecase) FetchByID(c context.Context, backinstocksubscriptionID string) (domain.BackInStockSubscription, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.backinstocksubscriptionRepository.FetchByID(ctx, backinstocksubscriptionID)
}

func (lu *backinstocksubscriptionUsecase) Fetch(c context.Context) ([]domain.BackInStockSubscription, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.backinstocksubscriptionRepository.Fetch(ctx)
}
