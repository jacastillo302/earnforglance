package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/seo"
	test "earnforglance/server/usecase/seo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestUrlRecordUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.UrlRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewUrlRecordUsecase(mockRepo, timeout)

	urlRecordID := bson.NewObjectID().Hex()

	updatedUrlRecord := domain.UrlRecord{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		PermissionRecordID: bson.NewObjectID(),
		Slug:               "example-category",
		IsActive:           false,
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
	usecase := test.NewUrlRecordUsecase(mockRepo, timeout)

	newUrlRecord := &domain.UrlRecord{
		PermissionRecordID: bson.NewObjectID(),
		Slug:               "example-product",
		IsActive:           true,
	}

	mockRepo.On("Create", mock.Anything, newUrlRecord).Return(nil)

	err := usecase.Create(context.Background(), newUrlRecord)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUrlRecordUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.UrlRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewUrlRecordUsecase(mockRepo, timeout)

	updatedUrlRecord := &domain.UrlRecord{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		PermissionRecordID: bson.NewObjectID(),
		Slug:               "example-category",
		IsActive:           false,
	}

	mockRepo.On("Update", mock.Anything, updatedUrlRecord).Return(nil)

	err := usecase.Update(context.Background(), updatedUrlRecord)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUrlRecordUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.UrlRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewUrlRecordUsecase(mockRepo, timeout)

	urlRecordID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, urlRecordID).Return(nil)

	err := usecase.Delete(context.Background(), urlRecordID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUrlRecordUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.UrlRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewUrlRecordUsecase(mockRepo, timeout)

	fetchedUrlRecords := []domain.UrlRecord{
		{
			ID:                 bson.NewObjectID(),
			PermissionRecordID: bson.NewObjectID(),
			Slug:               "example-product",
			IsActive:           true,
		},
		{
			ID:                 bson.NewObjectID(),
			PermissionRecordID: bson.NewObjectID(),
			Slug:               "example-category",
			IsActive:           false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedUrlRecords, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedUrlRecords, result)
	mockRepo.AssertExpectations(t)
}
