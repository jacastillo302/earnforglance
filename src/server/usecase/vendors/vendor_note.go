package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/vendors"
)

type vendornoteUsecase struct {
	vendornoteRepository domain.VendorNoteRepository
	contextTimeout       time.Duration
}

func NewVendorNoteUsecase(vendornoteRepository domain.VendorNoteRepository, timeout time.Duration) domain.VendorNoteUsecase {
	return &vendornoteUsecase{
		vendornoteRepository: vendornoteRepository,
		contextTimeout:       timeout,
	}
}

func (tu *vendornoteUsecase) Create(c context.Context, vendornote *domain.VendorNote) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendornoteRepository.Create(ctx, vendornote)
}

func (tu *vendornoteUsecase) Update(c context.Context, vendornote *domain.VendorNote) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendornoteRepository.Update(ctx, vendornote)
}

func (tu *vendornoteUsecase) Delete(c context.Context, vendornote *domain.VendorNote) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendornoteRepository.Delete(ctx, vendornote)
}

func (lu *vendornoteUsecase) FetchByID(c context.Context, vendornoteID string) (domain.VendorNote, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.vendornoteRepository.FetchByID(ctx, vendornoteID)
}

func (lu *vendornoteUsecase) Fetch(c context.Context) ([]domain.VendorNote, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.vendornoteRepository.Fetch(ctx)
}
