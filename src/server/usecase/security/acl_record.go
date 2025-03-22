package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/security"
)

type aclrecordUsecase struct {
	aclrecordRepository domain.AclRecordRepository
	contextTimeout      time.Duration
}

func NewAclRecordUsecase(aclrecordRepository domain.AclRecordRepository, timeout time.Duration) domain.AclRecordUsecase {
	return &aclrecordUsecase{
		aclrecordRepository: aclrecordRepository,
		contextTimeout:      timeout,
	}
}

func (tu *aclrecordUsecase) Create(c context.Context, aclrecord *domain.AclRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.aclrecordRepository.Create(ctx, aclrecord)
}

func (tu *aclrecordUsecase) Update(c context.Context, aclrecord *domain.AclRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.aclrecordRepository.Update(ctx, aclrecord)
}

func (tu *aclrecordUsecase) Delete(c context.Context, aclrecord string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.aclrecordRepository.Delete(ctx, aclrecord)
}

func (lu *aclrecordUsecase) FetchByID(c context.Context, aclrecordID string) (domain.AclRecord, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.aclrecordRepository.FetchByID(ctx, aclrecordID)
}

func (lu *aclrecordUsecase) Fetch(c context.Context) ([]domain.AclRecord, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.aclrecordRepository.Fetch(ctx)
}
