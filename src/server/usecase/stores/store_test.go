package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/stores"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestStoreUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.StoreRepository)
	timeout := time.Duration(10)
	usecase := NewStoreUsecase(mockRepo, timeout)

	storeID := primitive.NewObjectID().Hex()

	updatedStore := domain.Store{
		ID:                     primitive.NewObjectID(), // Existing ID of the record to update
		Name:                   "Updated Store",
		DefaultMetaKeywords:    "updated, store, gadgets",
		DefaultMetaDescription: "An updated description for the store.",
		DefaultTitle:           "Updated Store - Gadgets & More",
		HomepageTitle:          "Welcome to Updated Store",
		HomepageDescription:    "Discover updated gadgets and more.",
		Url:                    "https://www.updatedstore.com",
		SslEnabled:             false,
		Hosts:                  "www.updatedstore.com",
		DefaultLanguageID:      primitive.NewObjectID(),
		DisplayOrder:           2,
		CompanyName:            "Updated Store LLC",
		CompanyAddress:         "456 Updated Avenue, Townsville",
		CompanyPhoneNumber:     "+1-800-987-6543",
		CompanyVat:             "VAT654321",
		Deleted:                true,
	}

	mockRepo.On("FetchByID", mock.Anything, storeID).Return(updatedStore, nil)

	result, err := usecase.FetchByID(context.Background(), storeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedStore, result)
	mockRepo.AssertExpectations(t)
}

func TestStoreUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.StoreRepository)
	timeout := time.Duration(10)
	usecase := NewStoreUsecase(mockRepo, timeout)

	newStore := &domain.Store{
		Name:                   "Main Store",
		DefaultMetaKeywords:    "electronics, gadgets, store",
		DefaultMetaDescription: "The best store for electronics and gadgets.",
		DefaultTitle:           "Main Store - Electronics & Gadgets",
		HomepageTitle:          "Welcome to Main Store",
		HomepageDescription:    "Find the latest electronics and gadgets here.",
		Url:                    "https://www.mainstore.com",
		SslEnabled:             true,
		Hosts:                  "www.mainstore.com",
		DefaultLanguageID:      primitive.NewObjectID(),
		DisplayOrder:           1,
		CompanyName:            "Main Store Inc.",
		CompanyAddress:         "123 Main Street, Cityville",
		CompanyPhoneNumber:     "+1-800-123-4567",
		CompanyVat:             "VAT123456",
		Deleted:                false,
	}

	mockRepo.On("Create", mock.Anything, newStore).Return(nil)

	err := usecase.Create(context.Background(), newStore)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStoreUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.StoreRepository)
	timeout := time.Duration(10)
	usecase := NewStoreUsecase(mockRepo, timeout)

	updatedStore := &domain.Store{
		ID:                     primitive.NewObjectID(), // Existing ID of the record to update
		Name:                   "Updated Store",
		DefaultMetaKeywords:    "updated, store, gadgets",
		DefaultMetaDescription: "An updated description for the store.",
		DefaultTitle:           "Updated Store - Gadgets & More",
		HomepageTitle:          "Welcome to Updated Store",
		HomepageDescription:    "Discover updated gadgets and more.",
		Url:                    "https://www.updatedstore.com",
		SslEnabled:             false,
		Hosts:                  "www.updatedstore.com",
		DefaultLanguageID:      primitive.NewObjectID(),
		DisplayOrder:           2,
		CompanyName:            "Updated Store LLC",
		CompanyAddress:         "456 Updated Avenue, Townsville",
		CompanyPhoneNumber:     "+1-800-987-6543",
		CompanyVat:             "VAT654321",
		Deleted:                true,
	}
	mockRepo.On("Update", mock.Anything, updatedStore).Return(nil)

	err := usecase.Update(context.Background(), updatedStore)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStoreUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.StoreRepository)
	timeout := time.Duration(10)
	usecase := NewStoreUsecase(mockRepo, timeout)

	storeID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, storeID).Return(nil)

	err := usecase.Delete(context.Background(), storeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStoreUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.StoreRepository)
	timeout := time.Duration(10)
	usecase := NewStoreUsecase(mockRepo, timeout)

	fetchedStores := []domain.Store{
		{
			ID:                     primitive.NewObjectID(),
			Name:                   "Main Store",
			DefaultMetaKeywords:    "electronics, gadgets, store",
			DefaultMetaDescription: "The best store for electronics and gadgets.",
			DefaultTitle:           "Main Store - Electronics & Gadgets",
			HomepageTitle:          "Welcome to Main Store",
			HomepageDescription:    "Find the latest electronics and gadgets here.",
			Url:                    "https://www.mainstore.com",
			SslEnabled:             true,
			Hosts:                  "www.mainstore.com",
			DefaultLanguageID:      primitive.NewObjectID(),
			DisplayOrder:           1,
			CompanyName:            "Main Store Inc.",
			CompanyAddress:         "123 Main Street, Cityville",
			CompanyPhoneNumber:     "+1-800-123-4567",
			CompanyVat:             "VAT123456",
			Deleted:                false,
		},
		{
			ID:                     primitive.NewObjectID(),
			Name:                   "Updated Store",
			DefaultMetaKeywords:    "updated, store, gadgets",
			DefaultMetaDescription: "An updated description for the store.",
			DefaultTitle:           "Updated Store - Gadgets & More",
			HomepageTitle:          "Welcome to Updated Store",
			HomepageDescription:    "Discover updated gadgets and more.",
			Url:                    "https://www.updatedstore.com",
			SslEnabled:             false,
			Hosts:                  "www.updatedstore.com",
			DefaultLanguageID:      primitive.NewObjectID(),
			DisplayOrder:           2,
			CompanyName:            "Updated Store LLC",
			CompanyAddress:         "456 Updated Avenue, Townsville",
			CompanyPhoneNumber:     "+1-800-987-6543",
			CompanyVat:             "VAT654321",
			Deleted:                true,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedStores, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedStores, result)
	mockRepo.AssertExpectations(t)
}
