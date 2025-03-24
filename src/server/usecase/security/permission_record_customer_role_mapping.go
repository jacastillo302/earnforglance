package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/security"
)

type permissionrecordcustomerrolemappingUsecase struct {
	permissionrecordcustomerrolemappingRepository domain.PermissionRecordCustomerRoleMappingRepository
	contextTimeout                                time.Duration
}

func NewPermissionRecordCustomerRoleMappingUsecase(permissionrecordcustomerrolemappingRepository domain.PermissionRecordCustomerRoleMappingRepository, timeout time.Duration) domain.PermissionRecordCustomerRoleMappingUsecase {
	return &permissionrecordcustomerrolemappingUsecase{
		permissionrecordcustomerrolemappingRepository: permissionrecordcustomerrolemappingRepository,
		contextTimeout: timeout,
	}
}

func (tu *permissionrecordcustomerrolemappingUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.PermissionRecordCustomerRoleMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.permissionrecordcustomerrolemappingRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *permissionrecordcustomerrolemappingUsecase) Create(c context.Context, permissionrecordcustomerrolemapping *domain.PermissionRecordCustomerRoleMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.permissionrecordcustomerrolemappingRepository.Create(ctx, permissionrecordcustomerrolemapping)
}

func (tu *permissionrecordcustomerrolemappingUsecase) Update(c context.Context, permissionrecordcustomerrolemapping *domain.PermissionRecordCustomerRoleMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.permissionrecordcustomerrolemappingRepository.Update(ctx, permissionrecordcustomerrolemapping)
}

func (tu *permissionrecordcustomerrolemappingUsecase) Delete(c context.Context, permissionrecordcustomerrolemapping string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.permissionrecordcustomerrolemappingRepository.Delete(ctx, permissionrecordcustomerrolemapping)
}

func (lu *permissionrecordcustomerrolemappingUsecase) FetchByID(c context.Context, permissionrecordcustomerrolemappingID string) (domain.PermissionRecordCustomerRoleMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.permissionrecordcustomerrolemappingRepository.FetchByID(ctx, permissionrecordcustomerrolemappingID)
}

func (lu *permissionrecordcustomerrolemappingUsecase) Fetch(c context.Context) ([]domain.PermissionRecordCustomerRoleMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.permissionrecordcustomerrolemappingRepository.Fetch(ctx)
}
