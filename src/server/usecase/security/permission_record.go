package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/security"
)

type permissionrecordUsecase struct {
	permissionrecordRepository domain.PermissionRecordRepository
	contextTimeout             time.Duration
}

func NewPermissionRecordUsecase(permissionrecordRepository domain.PermissionRecordRepository, timeout time.Duration) domain.PermissionRecordUsecase {
	return &permissionrecordUsecase{
		permissionrecordRepository: permissionrecordRepository,
		contextTimeout:             timeout,
	}
}

func (tu *permissionrecordUsecase) Create(c context.Context, permissionrecord *domain.PermissionRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.permissionrecordRepository.Create(ctx, permissionrecord)
}

func (tu *permissionrecordUsecase) Update(c context.Context, permissionrecord *domain.PermissionRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.permissionrecordRepository.Update(ctx, permissionrecord)
}

func (tu *permissionrecordUsecase) Delete(c context.Context, permissionrecord *domain.PermissionRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.permissionrecordRepository.Delete(ctx, permissionrecord)
}

func (lu *permissionrecordUsecase) FetchByID(c context.Context, permissionrecordID string) (domain.PermissionRecord, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.permissionrecordRepository.FetchByID(ctx, permissionrecordID)
}

func (lu *permissionrecordUsecase) Fetch(c context.Context) ([]domain.PermissionRecord, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.permissionrecordRepository.Fetch(ctx)
}
