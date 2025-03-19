package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/seo"
)

type urlrecordUsecase struct {
	urlrecordRepository domain.UrlRecordRepository
	contextTimeout      time.Duration
}

func NewUrlRecordUsecase(urlrecordRepository domain.UrlRecordRepository, timeout time.Duration) domain.UrlRecordUsecase {
	return &urlrecordUsecase{
		urlrecordRepository: urlrecordRepository,
		contextTimeout:      timeout,
	}
}

func (tu *urlrecordUsecase) Create(c context.Context, urlrecord *domain.UrlRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.urlrecordRepository.Create(ctx, urlrecord)
}

func (tu *urlrecordUsecase) Update(c context.Context, urlrecord *domain.UrlRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.urlrecordRepository.Update(ctx, urlrecord)
}

func (tu *urlrecordUsecase) Delete(c context.Context, urlrecord *domain.UrlRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.urlrecordRepository.Delete(ctx, urlrecord)
}

func (lu *urlrecordUsecase) FetchByID(c context.Context, urlrecordID string) (domain.UrlRecord, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.urlrecordRepository.FetchByID(ctx, urlrecordID)
}

func (lu *urlrecordUsecase) Fetch(c context.Context) ([]domain.UrlRecord, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.urlrecordRepository.Fetch(ctx)
}
