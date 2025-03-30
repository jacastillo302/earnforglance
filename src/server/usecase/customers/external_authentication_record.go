package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type externalAuthenticationRecordUsecase struct {
	externalAuthenticationRecordRepository domain.ExternalAuthenticationRecordRepository
	contextTimeout                         time.Duration
}

func NewExternalAuthenticationRecordUsecase(externalAuthenticationRecordRepository domain.ExternalAuthenticationRecordRepository, timeout time.Duration) domain.ExternalAuthenticationRecordUsecase {
	return &externalAuthenticationRecordUsecase{
		externalAuthenticationRecordRepository: externalAuthenticationRecordRepository,
		contextTimeout:                         timeout,
	}
}

func (tu *externalAuthenticationRecordUsecase) CreateMany(c context.Context, items []domain.ExternalAuthenticationRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.externalAuthenticationRecordRepository.CreateMany(ctx, items)
}

func (tu *externalAuthenticationRecordUsecase) Create(c context.Context, externalAuthenticationRecord *domain.ExternalAuthenticationRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.externalAuthenticationRecordRepository.Create(ctx, externalAuthenticationRecord)
}

func (tu *externalAuthenticationRecordUsecase) Update(c context.Context, externalAuthenticationRecord *domain.ExternalAuthenticationRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.externalAuthenticationRecordRepository.Update(ctx, externalAuthenticationRecord)
}

func (tu *externalAuthenticationRecordUsecase) Delete(c context.Context, externalAuthenticationRecord string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.externalAuthenticationRecordRepository.Delete(ctx, externalAuthenticationRecord)
}

func (lu *externalAuthenticationRecordUsecase) FetchByID(c context.Context, externalAuthenticationRecordID string) (domain.ExternalAuthenticationRecord, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.externalAuthenticationRecordRepository.FetchByID(ctx, externalAuthenticationRecordID)
}

func (lu *externalAuthenticationRecordUsecase) Fetch(c context.Context) ([]domain.ExternalAuthenticationRecord, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.externalAuthenticationRecordRepository.Fetch(ctx)
}
