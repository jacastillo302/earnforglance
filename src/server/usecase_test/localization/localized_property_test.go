package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/localization"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/localization"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestLocalizedPropertyUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.LocalizedPropertyRepository)
	timeout := time.Duration(10)
	usecase := test.NewLocalizedPropertyUsecase(mockRepo, timeout)

	localizationID := bson.NewObjectID().Hex()

	updatedLocalizedProperty := domain.LocalizedProperty{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		PermissionRecordID: bson.NewObjectID(),
		LanguageID:         bson.NewObjectID(),
		LocaleKeyGroup:     "Category",
		LocaleKey:          "Description",
		LocaleValue:        "Electronics and Gadgets",
	}

	mockRepo.On("FetchByID", mock.Anything, localizationID).Return(updatedLocalizedProperty, nil)

	result, err := usecase.FetchByID(context.Background(), localizationID)

	assert.NoError(t, err)
	assert.Equal(t, updatedLocalizedProperty, result)
	mockRepo.AssertExpectations(t)
}

func TestLocalizedPropertyUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.LocalizedPropertyRepository)
	timeout := time.Duration(10)
	usecase := test.NewLocalizedPropertyUsecase(mockRepo, timeout)
	newLocalizedProperty := &domain.LocalizedProperty{
		PermissionRecordID: bson.NewObjectID(),
		LanguageID:         bson.NewObjectID(),
		LocaleKeyGroup:     "Product",
		LocaleKey:          "Name",
		LocaleValue:        "Laptop",
	}

	mockRepo.On("Create", mock.Anything, newLocalizedProperty).Return(nil)

	err := usecase.Create(context.Background(), newLocalizedProperty)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLocalizedPropertyUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.LocalizedPropertyRepository)
	timeout := time.Duration(10)
	usecase := test.NewLocalizedPropertyUsecase(mockRepo, timeout)

	updatedLocalizedProperty := &domain.LocalizedProperty{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		PermissionRecordID: bson.NewObjectID(),
		LanguageID:         bson.NewObjectID(),
		LocaleKeyGroup:     "Category",
		LocaleKey:          "Description",
		LocaleValue:        "Electronics and Gadgets",
	}

	mockRepo.On("Update", mock.Anything, updatedLocalizedProperty).Return(nil)

	err := usecase.Update(context.Background(), updatedLocalizedProperty)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLocalizedPropertyUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.LocalizedPropertyRepository)
	timeout := time.Duration(10)
	usecase := test.NewLocalizedPropertyUsecase(mockRepo, timeout)

	localizationID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, localizationID).Return(nil)

	err := usecase.Delete(context.Background(), localizationID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLocalizedPropertyUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.LocalizedPropertyRepository)
	timeout := time.Duration(10)
	usecase := test.NewLocalizedPropertyUsecase(mockRepo, timeout)

	fetchedLocalizedProperties := []domain.LocalizedProperty{
		{
			ID:                 bson.NewObjectID(),
			PermissionRecordID: bson.NewObjectID(),
			LanguageID:         bson.NewObjectID(),
			LocaleKeyGroup:     "Product",
			LocaleKey:          "Name",
			LocaleValue:        "Laptop",
		},
		{
			ID:                 bson.NewObjectID(),
			PermissionRecordID: bson.NewObjectID(),
			LanguageID:         bson.NewObjectID(),
			LocaleKeyGroup:     "Category",
			LocaleKey:          "Description",
			LocaleValue:        "Electronics and Gadgets",
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedLocalizedProperties, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedLocalizedProperties, result)
	mockRepo.AssertExpectations(t)
}
