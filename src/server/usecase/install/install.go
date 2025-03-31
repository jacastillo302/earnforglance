package install

import (
	"context"
	"time"

	domain "earnforglance/server/domain/install"
)

type InstallUsecase struct {
	InstallRepository domain.InstallRepository
	contextTimeout    time.Duration
}

func NewInstallUsecase(IsntallRepository domain.InstallRepository, timeout time.Duration) domain.InstallLogUsecase {
	return &InstallUsecase{
		InstallRepository: IsntallRepository,
		contextTimeout:    timeout,
	}
}

func (tu *InstallUsecase) PingDatabase(c context.Context) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.PingDatabase(ctx)
}
