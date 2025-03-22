package usecase

import (
	"context"
	domain "earnforglance/server/domain/common"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddressAttributeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.AddressAttributeRepository)
	timeout := time.Duration(10)
	usecase := NewAddressAttributeUsecase(mockRepo, timeout)

	addressAttributeID := primitive.NewObjectID().Hex()

	updatedAddressAttribute := domain.AddressAttribute{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		Name:                            "City",
		IsRequired:                      false,
		AttributeControlTypeID:          2,
		DisplayOrder:                    2,
		DefaultValue:                    "New York",
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: ".jpg,.png",
		ValidationFileMaximumSize:       new(int),
		ConditionAttributeXml:           "<conditions><required>false</required></conditions>",
	}

	mockRepo.On("FetchByID", mock.Anything, addressAttributeID).Return(updatedAddressAttribute, nil)

	result, err := usecase.FetchByID(context.Background(), addressAttributeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedAddressAttribute, result)
	mockRepo.AssertExpectations(t)
}

func TestAddressAttributeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.AddressAttributeRepository)
	timeout := time.Duration(10)
	usecase := NewAddressAttributeUsecase(mockRepo, timeout)

	newAddressAttribute := &domain.AddressAttribute{
		Name:                            "Street Address",
		IsRequired:                      true,
		AttributeControlTypeID:          1,
		DisplayOrder:                    1,
		DefaultValue:                    "123 Main St",
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: "",
		ValidationFileMaximumSize:       nil,
		ConditionAttributeXml:           "<conditions><required>true</required></conditions>",
	}
	*newAddressAttribute.ValidationMinLength = 5
	*newAddressAttribute.ValidationMaxLength = 100

	mockRepo.On("Create", mock.Anything, newAddressAttribute).Return(nil)

	err := usecase.Create(context.Background(), newAddressAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddressAttributeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.AddressAttributeRepository)
	timeout := time.Duration(10)
	usecase := NewAddressAttributeUsecase(mockRepo, timeout)

	updatedAddressAttribute := &domain.AddressAttribute{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		Name:                            "City",
		IsRequired:                      false,
		AttributeControlTypeID:          2,
		DisplayOrder:                    2,
		DefaultValue:                    "New York",
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: ".jpg,.png",
		ValidationFileMaximumSize:       new(int),
		ConditionAttributeXml:           "<conditions><required>false</required></conditions>",
	}
	*updatedAddressAttribute.ValidationMinLength = 3
	*updatedAddressAttribute.ValidationMaxLength = 50
	*updatedAddressAttribute.ValidationFileMaximumSize = 2048

	mockRepo.On("Update", mock.Anything, updatedAddressAttribute).Return(nil)

	err := usecase.Update(context.Background(), updatedAddressAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddressAttributeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.AddressAttributeRepository)
	timeout := time.Duration(10)
	usecase := NewAddressAttributeUsecase(mockRepo, timeout)

	addressAttributeID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, addressAttributeID).Return(nil)

	err := usecase.Delete(context.Background(), addressAttributeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddressAttributeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.AddressAttributeRepository)
	timeout := time.Duration(10)
	usecase := NewAddressAttributeUsecase(mockRepo, timeout)

	fetchedAddressAttributes := []domain.AddressAttribute{
		{
			ID:                              primitive.NewObjectID(),
			Name:                            "Street Address",
			IsRequired:                      true,
			AttributeControlTypeID:          1,
			DisplayOrder:                    1,
			DefaultValue:                    "123 Main St",
			ValidationMinLength:             new(int),
			ValidationMaxLength:             new(int),
			ValidationFileAllowedExtensions: "",
			ValidationFileMaximumSize:       nil,
			ConditionAttributeXml:           "<conditions><required>true</required></conditions>",
		},
		{
			ID:                              primitive.NewObjectID(),
			Name:                            "City",
			IsRequired:                      false,
			AttributeControlTypeID:          2,
			DisplayOrder:                    2,
			DefaultValue:                    "New York",
			ValidationMinLength:             new(int),
			ValidationMaxLength:             new(int),
			ValidationFileAllowedExtensions: ".jpg,.png",
			ValidationFileMaximumSize:       new(int),
			ConditionAttributeXml:           "<conditions><required>false</required></conditions>",
		},
	}
	*fetchedAddressAttributes[0].ValidationMinLength = 5
	*fetchedAddressAttributes[0].ValidationMaxLength = 100
	*fetchedAddressAttributes[1].ValidationMinLength = 3
	*fetchedAddressAttributes[1].ValidationMaxLength = 50
	*fetchedAddressAttributes[1].ValidationFileMaximumSize = 2048

	mockRepo.On("Fetch", mock.Anything).Return(fetchedAddressAttributes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedAddressAttributes, result)
	mockRepo.AssertExpectations(t)
}
