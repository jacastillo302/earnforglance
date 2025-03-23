package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/customers"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/customers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCustomerAttributeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CustomerAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCustomerAttributeUsecase(mockRepo, timeout)

	customerAttributeID := primitive.NewObjectID().Hex()

	updatedCustomerAttribute := domain.CustomerAttribute{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		Name:                            "Preferred Language",
		IsRequired:                      false,
		AttributeControlTypeID:          2,
		DisplayOrder:                    2,
		DefaultValue:                    "English",
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: ".txt,.pdf",
		ValidationFileMaximumSize:       new(int),
		ConditionAttributeXml:           "<conditions><required>false</required></conditions>",
	}

	mockRepo.On("FetchByID", mock.Anything, customerAttributeID).Return(updatedCustomerAttribute, nil)

	result, err := usecase.FetchByID(context.Background(), customerAttributeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCustomerAttribute, result)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAttributeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CustomerAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCustomerAttributeUsecase(mockRepo, timeout)

	newCustomerAttribute := &domain.CustomerAttribute{
		Name:                            "Date of Birth",
		IsRequired:                      true,
		AttributeControlTypeID:          1,
		DisplayOrder:                    1,
		DefaultValue:                    "01/01/2000",
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: "",
		ValidationFileMaximumSize:       nil,
		ConditionAttributeXml:           "<conditions><required>true</required></conditions>",
	}
	*newCustomerAttribute.ValidationMinLength = 10
	*newCustomerAttribute.ValidationMaxLength = 10

	mockRepo.On("Create", mock.Anything, newCustomerAttribute).Return(nil)

	err := usecase.Create(context.Background(), newCustomerAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAttributeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CustomerAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCustomerAttributeUsecase(mockRepo, timeout)

	updatedCustomerAttribute := &domain.CustomerAttribute{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		Name:                            "Preferred Language",
		IsRequired:                      false,
		AttributeControlTypeID:          2,
		DisplayOrder:                    2,
		DefaultValue:                    "English",
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: ".txt,.pdf",
		ValidationFileMaximumSize:       new(int),
		ConditionAttributeXml:           "<conditions><required>false</required></conditions>",
	}
	*updatedCustomerAttribute.ValidationMinLength = 3
	*updatedCustomerAttribute.ValidationMaxLength = 20
	*updatedCustomerAttribute.ValidationFileMaximumSize = 2048

	mockRepo.On("Update", mock.Anything, updatedCustomerAttribute).Return(nil)

	err := usecase.Update(context.Background(), updatedCustomerAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAttributeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CustomerAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCustomerAttributeUsecase(mockRepo, timeout)

	customerAttributeID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customerAttributeID).Return(nil)

	err := usecase.Delete(context.Background(), customerAttributeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAttributeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CustomerAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCustomerAttributeUsecase(mockRepo, timeout)

	fetchedCustomerAttributes := []domain.CustomerAttribute{
		{
			ID:                              primitive.NewObjectID(),
			Name:                            "Date of Birth",
			IsRequired:                      true,
			AttributeControlTypeID:          1,
			DisplayOrder:                    1,
			DefaultValue:                    "01/01/2000",
			ValidationMinLength:             new(int),
			ValidationMaxLength:             new(int),
			ValidationFileAllowedExtensions: "",
			ValidationFileMaximumSize:       nil,
			ConditionAttributeXml:           "<conditions><required>true</required></conditions>",
		},
		{
			ID:                              primitive.NewObjectID(),
			Name:                            "Preferred Language",
			IsRequired:                      false,
			AttributeControlTypeID:          2,
			DisplayOrder:                    2,
			DefaultValue:                    "English",
			ValidationMinLength:             new(int),
			ValidationMaxLength:             new(int),
			ValidationFileAllowedExtensions: ".txt,.pdf",
			ValidationFileMaximumSize:       new(int),
			ConditionAttributeXml:           "<conditions><required>false</required></conditions>",
		},
	}
	*fetchedCustomerAttributes[0].ValidationMinLength = 10
	*fetchedCustomerAttributes[0].ValidationMaxLength = 10
	*fetchedCustomerAttributes[1].ValidationMinLength = 3
	*fetchedCustomerAttributes[1].ValidationMaxLength = 20
	*fetchedCustomerAttributes[1].ValidationFileMaximumSize = 2048

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCustomerAttributes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCustomerAttributes, result)
	mockRepo.AssertExpectations(t)
}
