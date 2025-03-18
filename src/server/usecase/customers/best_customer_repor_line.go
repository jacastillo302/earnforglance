package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type bestcustomerreporlineUsecase struct {
	bestcustomerreporlineRepository domain.BestCustomerReporLineRepository
	contextTimeout                  time.Duration
}

func NewBestCustomerReporLineUsecase(bestcustomerreporlineRepository domain.BestCustomerReporLineRepository, timeout time.Duration) domain.BestCustomerReporLineUsecase {
	// ...existing code...
}

func (tu *bestcustomerreporlineUsecase) Create(c context.Context, bestcustomerreporline *domain.BestCustomerReporLine) error {
	// ...existing code...
}

func (tu *bestcustomerreporlineUsecase) Update(c context.Context, bestcustomerreporline *domain.BestCustomerReporLine) error {
	// ...existing code...
}

func (tu *bestcustomerreporlineUsecase) Delete(c context.Context, bestcustomerreporline *domain.BestCustomerReporLine) error {
	// ...existing code...
}

func (lu *bestcustomerreporlineUsecase) FetchByID(c context.Context, bestcustomerreporlineID string) (domain.BestCustomerReporLine, error) {
	// ...existing code...
}

func (lu *bestcustomerreporlineUsecase) Fetch(c context.Context) ([]domain.BestCustomerReporLine, error) {
	// ...existing code...
}
