package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/logging"
)

type logUsecase struct {
	logRepository  domain.LogRepository
	contextTimeout time.Duration
}

func NewLogUsecase(logRepository domain.LogRepository, timeout time.Duration) domain.LogUsecase {
	return &logUsecase{
		logRepository:  logRepository,
		contextTimeout: timeout,
	}
}

func (lu *logUsecase) Create(c context.Context, log *domain.Log) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.logRepository.Create(ctx, log)
}

func (lu *logUsecase) Update(c context.Context, log *domain.Log) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.logRepository.Update(ctx, log)
}

func (lu *logUsecase) Delete(c context.Context, log *domain.Log) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.logRepository.Delete(ctx, log)
}

func (lu *logUsecase) FetchByID(c context.Context, logID string) (domain.Log, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.logRepository.FetchByID(ctx, logID)
}

func (lu *logUsecase) Fetch(c context.Context) ([]domain.Log, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.logRepository.Fetch(ctx)
}
