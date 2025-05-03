package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/attributes"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/attributes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestPermisionRecordAttributeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PermisionRecordAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermisionRecordAttributeUsecase(mockRepo, timeout)

	permisionRecordAttributeID := bson.NewObjectID().Hex()

	updatedPermisionRecordAttribute := domain.PermisionRecordAttribute{
		ID:                              bson.NewObjectID(), // Existing ID of the record to update
		Name:                            "Preferred Language",
		IsRequired:                      false,
		DisplayOrder:                    2,
		DefaultValue:                    "English",
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: ".txt,.pdf",
		ValidationFileMaximumSize:       new(int),
		ConditionAttributeXml:           "<conditions><required>false</required></conditions>",
	}

	mockRepo.On("FetchByID", mock.Anything, permisionRecordAttributeID).Return(updatedPermisionRecordAttribute, nil)

	result, err := usecase.FetchByID(context.Background(), permisionRecordAttributeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPermisionRecordAttribute, result)
	mockRepo.AssertExpectations(t)
}

func TestPermisionRecordAttributeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PermisionRecordAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermisionRecordAttributeUsecase(mockRepo, timeout)

	newPermisionRecordAttribute := &domain.PermisionRecordAttribute{
		Name:                            "Date of Birth",
		IsRequired:                      true,
		DisplayOrder:                    1,
		DefaultValue:                    "01/01/2000",
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: "",
		ValidationFileMaximumSize:       nil,
		ConditionAttributeXml:           "<conditions><required>true</required></conditions>",
	}
	*newPermisionRecordAttribute.ValidationMinLength = 10
	*newPermisionRecordAttribute.ValidationMaxLength = 10

	mockRepo.On("Create", mock.Anything, newPermisionRecordAttribute).Return(nil)

	err := usecase.Create(context.Background(), newPermisionRecordAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermisionRecordAttributeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PermisionRecordAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermisionRecordAttributeUsecase(mockRepo, timeout)

	updatedPermisionRecordAttribute := &domain.PermisionRecordAttribute{
		ID:                              bson.NewObjectID(), // Existing ID of the record to update
		Name:                            "Preferred Language",
		IsRequired:                      false,
		DisplayOrder:                    2,
		DefaultValue:                    "English",
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: ".txt,.pdf",
		ValidationFileMaximumSize:       new(int),
		ConditionAttributeXml:           "<conditions><required>false</required></conditions>",
	}
	*updatedPermisionRecordAttribute.ValidationMinLength = 3
	*updatedPermisionRecordAttribute.ValidationMaxLength = 20
	*updatedPermisionRecordAttribute.ValidationFileMaximumSize = 2048

	mockRepo.On("Update", mock.Anything, updatedPermisionRecordAttribute).Return(nil)

	err := usecase.Update(context.Background(), updatedPermisionRecordAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermisionRecordAttributeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PermisionRecordAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermisionRecordAttributeUsecase(mockRepo, timeout)

	permisionRecordAttributeID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, permisionRecordAttributeID).Return(nil)

	err := usecase.Delete(context.Background(), permisionRecordAttributeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermisionRecordAttributeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PermisionRecordAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermisionRecordAttributeUsecase(mockRepo, timeout)

	fetchedPermisionRecordAttributes := []domain.PermisionRecordAttribute{
		{
			ID:                              bson.NewObjectID(),
			Name:                            "Date of Birth",
			IsRequired:                      true,
			DisplayOrder:                    1,
			DefaultValue:                    "01/01/2000",
			ValidationMinLength:             new(int),
			ValidationMaxLength:             new(int),
			ValidationFileAllowedExtensions: "",
			ValidationFileMaximumSize:       nil,
			ConditionAttributeXml:           "<conditions><required>true</required></conditions>",
		},
		{
			ID:                              bson.NewObjectID(),
			Name:                            "Preferred Language",
			IsRequired:                      false,
			DisplayOrder:                    2,
			DefaultValue:                    "English",
			ValidationMinLength:             new(int),
			ValidationMaxLength:             new(int),
			ValidationFileAllowedExtensions: ".txt,.pdf",
			ValidationFileMaximumSize:       new(int),
			ConditionAttributeXml:           "<conditions><required>false</required></conditions>",
		},
	}
	*fetchedPermisionRecordAttributes[0].ValidationMinLength = 10
	*fetchedPermisionRecordAttributes[0].ValidationMaxLength = 10
	*fetchedPermisionRecordAttributes[1].ValidationMinLength = 3
	*fetchedPermisionRecordAttributes[1].ValidationMaxLength = 20
	*fetchedPermisionRecordAttributes[1].ValidationFileMaximumSize = 2048

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPermisionRecordAttributes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPermisionRecordAttributes, result)
	mockRepo.AssertExpectations(t)
}
