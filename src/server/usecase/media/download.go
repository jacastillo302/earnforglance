package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/media"
)

type downloadUsecase struct {
	downloadRepository domain.DownloadRepository
	contextTimeout     time.Duration
}

func NewDownloadUsecase(downloadRepository domain.DownloadRepository, timeout time.Duration) domain.DownloadUsecase {
	return &downloadUsecase{
		downloadRepository: downloadRepository,
		contextTimeout:     timeout,
	}
}

func (tu *downloadUsecase) Create(c context.Context, download *domain.Download) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.downloadRepository.Create(ctx, download)
}

func (tu *downloadUsecase) Update(c context.Context, download *domain.Download) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.downloadRepository.Update(ctx, download)
}

func (tu *downloadUsecase) Delete(c context.Context, download *domain.Download) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.downloadRepository.Delete(ctx, download)
}

func (lu *downloadUsecase) FetchByID(c context.Context, downloadID string) (domain.Download, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.downloadRepository.FetchByID(ctx, downloadID)
}

func (lu *downloadUsecase) Fetch(c context.Context) ([]domain.Download, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.downloadRepository.Fetch(ctx)
}
