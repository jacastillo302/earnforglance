package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/seo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUrlRecordUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.UrlRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewUrlRecordUsecase(mockRepo, timeout)

	urlRecordID := primitive.NewObjectID().Hex()

	updatedUrlRecord := domain.UrlRecord{
		ID:         primitive.NewObjectID(), // Existing ID of the record to update
		EntityID:   primitive.NewObjectID(),
		EntityName: "Category",
		Slug:       "example-category",
		IsActive:   false,
		LanguageID: primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, urlRecordID).Return(updatedUrlRecord, nil)

	result, err := usecase.FetchByID(context.Background(), urlRecordID)

	assert.NoError(t, err)
	assert.Equal(t, updatedUrlRecord, result)
	mockRepo.AssertExpectations(t)
}

func TestUrlRecordUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.UrlRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewUrlRecordUsecase(mockRepo, timeout)

	newUrlRecord := &domain.UrlRecord{
		EntityID:   primitive.NewObjectID(),
		EntityName: "Product",
		Slug:       "example-product",
		IsActive:   true,
		LanguageID: primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newUrlRecord).Return(nil)

	err := usecase.Create(context.Background(), newUrlRecord)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUrlRecordUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.UrlRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewUrlRecordUsecase(mockRepo, timeout)

	updatedUrlRecord := &domain.UrlRecord{
		ID:         primitive.NewObjectID(), // Existing ID of the record to update
		EntityID:   primitive.NewObjectID(),
		EntityName: "Category",
		Slug:       "example-category",
		IsActive:   false,
		LanguageID: primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedUrlRecord).Return(nil)

	err := usecase.Update(context.Background(), updatedUrlRecord)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUrlRecordUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.UrlRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewUrlRecordUsecase(mockRepo, timeout)

	urlRecordID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, urlRecordID).Return(nil)

	err := usecase.Delete(context.Background(), urlRecordID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUrlRecordUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.UrlRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewUrlRecordUsecase(mockRepo, timeout)

	fetchedUrlRecords := []domain.UrlRecord{
		{
			ID:         primitive.NewObjectID(),
			EntityID:   primitive.NewObjectID(),
			EntityName: "Product",
			Slug:       "example-product",
			IsActive:   true,
			LanguageID: primitive.NewObjectID(),
		},
		{
			ID:         primitive.NewObjectID(),
			EntityID:   primitive.NewObjectID(),
			EntityName: "Category",
			Slug:       "example-category",
			IsActive:   false,
			LanguageID: primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedUrlRecords, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedUrlRecords, result)
	mockRepo.AssertExpectations(t)
}
