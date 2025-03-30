package usecase_test

import (
	"context"
	domian "earnforglance/server/domain/customers"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/customers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCustomerRoleUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CustomerRoleRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerRoleUsecase(mockRepo, timeout)

	customerID := primitive.NewObjectID().Hex()

	updatedCustomerRole := domian.CustomerRole{
		ID:                      primitive.NewObjectID(),
		Name:                    "Registered",
		FreeShipping:            false,
		TaxExempt:               false,
		Active:                  true,
		IsSystemRole:            true,
		SystemName:              "Registered",
		EnablePasswordLifetime:  false,
		OverrideTaxDisplayType:  false,
		DefaultTaxDisplayTypeID: 2,
		PurchasedWithProductId:  0,
	}

	mockRepo.On("FetchByID", mock.Anything, customerID).Return(updatedCustomerRole, nil)

	result, err := usecase.FetchByID(context.Background(), customerID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCustomerRole, result)
	mockRepo.AssertExpectations(t)
}

func TestCustomerRoleUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CustomerRoleRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerRoleUsecase(mockRepo, timeout)

	newCustomerRole := &domian.CustomerRole{
		Name:                    "Administrator", // Example role name
		FreeShipping:            true,            // Indicates free shipping for this role
		TaxExempt:               true,            // Indicates tax exemption for this role
		Active:                  true,            // Indicates the role is active
		IsSystemRole:            true,            // Indicates this is a system role
		SystemName:              "Admin",         // Example system name
		EnablePasswordLifetime:  true,            // Indicates password lifetime enforcement
		OverrideTaxDisplayType:  false,           // Indicates no custom tax display type
		DefaultTaxDisplayTypeID: 1,               // Example default tax display type ID
		PurchasedWithProductId:  0,               // No product required for this role

	}

	mockRepo.On("Create", mock.Anything, newCustomerRole).Return(nil)

	err := usecase.Create(context.Background(), newCustomerRole)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerRoleUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CustomerRoleRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerRoleUsecase(mockRepo, timeout)

	updatedCustomerRole := &domian.CustomerRole{
		ID:                      primitive.NewObjectID(), // Generate a new MongoDB ObjectID
		Name:                    "Administrator",         // Example role name
		FreeShipping:            true,                    // Indicates free shipping for this role
		TaxExempt:               true,                    // Indicates tax exemption for this role
		Active:                  true,                    // Indicates the role is active
		IsSystemRole:            true,                    // Indicates this is a system role
		SystemName:              "Admin",                 // Example system name
		EnablePasswordLifetime:  true,                    // Indicates password lifetime enforcement
		OverrideTaxDisplayType:  false,                   // Indicates no custom tax display type
		DefaultTaxDisplayTypeID: 1,                       // Example default tax display type ID
		PurchasedWithProductId:  0,                       // No product required for this role

	}

	mockRepo.On("Update", mock.Anything, updatedCustomerRole).Return(nil)

	err := usecase.Update(context.Background(), updatedCustomerRole)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerRoleUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CustomerRoleRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerRoleUsecase(mockRepo, timeout)

	customerID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customerID).Return(nil)

	err := usecase.Delete(context.Background(), customerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerRoleUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CustomerRoleRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerRoleUsecase(mockRepo, timeout)

	fetchedCustomerRoles := []domian.CustomerRole{
		{
			ID:                      primitive.NewObjectID(),
			Name:                    "Registered",
			FreeShipping:            false,
			TaxExempt:               false,
			Active:                  true,
			IsSystemRole:            true,
			SystemName:              "Registered",
			EnablePasswordLifetime:  false,
			OverrideTaxDisplayType:  false,
			DefaultTaxDisplayTypeID: 2,
			PurchasedWithProductId:  0,
		},
		{
			ID:                      primitive.NewObjectID(),
			Name:                    "Guest",
			FreeShipping:            false,
			TaxExempt:               false,
			Active:                  true,
			IsSystemRole:            true,
			SystemName:              "Guest",
			EnablePasswordLifetime:  false,
			OverrideTaxDisplayType:  false,
			DefaultTaxDisplayTypeID: 0,
			PurchasedWithProductId:  0,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCustomerRoles, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCustomerRoles, result)
	mockRepo.AssertExpectations(t)
}
