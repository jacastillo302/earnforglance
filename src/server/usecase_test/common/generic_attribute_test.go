package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/common"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/common"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestGenericAttributeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.GenericAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGenericAttributeUsecase(mockRepo, timeout)

	genericAttributeID := bson.NewObjectID().Hex()

	updatedGenericAttribute := domain.GenericAttribute{
		ID:                      bson.NewObjectID(), // Existing ID of the record to update
		EntityID:                bson.NewObjectID(),
		KeyGroup:                "Customer",
		Key:                     "PreferredLanguage",
		Value:                   "English",
		StoreID:                 bson.NewObjectID(),
		CreatedOrUpdatedDateUTC: new(time.Time),
	}

	mockRepo.On("FetchByID", mock.Anything, genericAttributeID).Return(updatedGenericAttribute, nil)

	result, err := usecase.FetchByID(context.Background(), genericAttributeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedGenericAttribute, result)
	mockRepo.AssertExpectations(t)
}

func TestGenericAttributeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.GenericAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGenericAttributeUsecase(mockRepo, timeout)

	newGenericAttribute := &domain.GenericAttribute{
		EntityID:                bson.NewObjectID(),
		KeyGroup:                "Product",
		Key:                     "Color",
		Value:                   "Red",
		StoreID:                 bson.NewObjectID(),
		CreatedOrUpdatedDateUTC: nil,
	}

	mockRepo.On("Create", mock.Anything, newGenericAttribute).Return(nil)

	err := usecase.Create(context.Background(), newGenericAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGenericAttributeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.GenericAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGenericAttributeUsecase(mockRepo, timeout)

	updatedGenericAttribute := &domain.GenericAttribute{
		ID:                      bson.NewObjectID(), // Existing ID of the record to update
		EntityID:                bson.NewObjectID(),
		KeyGroup:                "Customer",
		Key:                     "PreferredLanguage",
		Value:                   "English",
		StoreID:                 bson.NewObjectID(),
		CreatedOrUpdatedDateUTC: new(time.Time),
	}
	*updatedGenericAttribute.CreatedOrUpdatedDateUTC = time.Now()

	mockRepo.On("Update", mock.Anything, updatedGenericAttribute).Return(nil)

	err := usecase.Update(context.Background(), updatedGenericAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGenericAttributeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.GenericAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGenericAttributeUsecase(mockRepo, timeout)

	genericAttributeID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, genericAttributeID).Return(nil)

	err := usecase.Delete(context.Background(), genericAttributeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGenericAttributeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.GenericAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGenericAttributeUsecase(mockRepo, timeout)

	fetchedGenericAttributes := []domain.GenericAttribute{
		{
			ID:                      bson.NewObjectID(),
			EntityID:                bson.NewObjectID(),
			KeyGroup:                "Product",
			Key:                     "Color",
			Value:                   "Red",
			StoreID:                 bson.NewObjectID(),
			CreatedOrUpdatedDateUTC: nil,
		},
		{
			ID:                      bson.NewObjectID(),
			EntityID:                bson.NewObjectID(),
			KeyGroup:                "Customer",
			Key:                     "PreferredLanguage",
			Value:                   "English",
			StoreID:                 bson.NewObjectID(),
			CreatedOrUpdatedDateUTC: new(time.Time),
		},
	}
	*fetchedGenericAttributes[1].CreatedOrUpdatedDateUTC = time.Now().AddDate(0, 0, -7) // 7 days ago

	mockRepo.On("Fetch", mock.Anything).Return(fetchedGenericAttributes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedGenericAttributes, result)
	mockRepo.AssertExpectations(t)
}
