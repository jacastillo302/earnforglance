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

func TestLanguageUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.LanguageRepository)
	timeout := time.Duration(10)
	usecase := test.NewLanguageUsecase(mockRepo, timeout)

	languageID := bson.NewObjectID().Hex()

	updatedLanguage := domain.Language{
		ID:                bson.NewObjectID(), // Existing ID of the record to update
		Name:              "Spanish",
		LanguageCulture:   "es-ES",
		UniqueSeoCode:     "es",
		FlagImageFileName: "es.png",
		Rtl:               false,
		LimitedToStores:   true,
		DefaultCurrencyID: bson.NewObjectID(),
		Published:         false,
		DisplayOrder:      2,
	}

	mockRepo.On("FetchByID", mock.Anything, languageID).Return(updatedLanguage, nil)

	result, err := usecase.FetchByID(context.Background(), languageID)

	assert.NoError(t, err)
	assert.Equal(t, updatedLanguage, result)
	mockRepo.AssertExpectations(t)
}

func TestLanguageUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.LanguageRepository)
	timeout := time.Duration(10)
	usecase := test.NewLanguageUsecase(mockRepo, timeout)

	newLanguage := &domain.Language{
		Name:              "English",
		LanguageCulture:   "en-US",
		UniqueSeoCode:     "en",
		FlagImageFileName: "us.png",
		Rtl:               false,
		LimitedToStores:   false,
		DefaultCurrencyID: bson.NewObjectID(),
		Published:         true,
		DisplayOrder:      1,
	}

	mockRepo.On("Create", mock.Anything, newLanguage).Return(nil)

	err := usecase.Create(context.Background(), newLanguage)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLanguageUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.LanguageRepository)
	timeout := time.Duration(10)
	usecase := test.NewLanguageUsecase(mockRepo, timeout)

	updatedLanguage := &domain.Language{
		ID:                bson.NewObjectID(), // Existing ID of the record to update
		Name:              "Spanish",
		LanguageCulture:   "es-ES",
		UniqueSeoCode:     "es",
		FlagImageFileName: "es.png",
		Rtl:               false,
		LimitedToStores:   true,
		DefaultCurrencyID: bson.NewObjectID(),
		Published:         false,
		DisplayOrder:      2,
	}

	mockRepo.On("Update", mock.Anything, updatedLanguage).Return(nil)

	err := usecase.Update(context.Background(), updatedLanguage)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLanguageUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.LanguageRepository)
	timeout := time.Duration(10)
	usecase := test.NewLanguageUsecase(mockRepo, timeout)

	languageID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, languageID).Return(nil)

	err := usecase.Delete(context.Background(), languageID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLanguageUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.LanguageRepository)
	timeout := time.Duration(10)
	usecase := test.NewLanguageUsecase(mockRepo, timeout)

	fetchedLanguages := []domain.Language{
		{
			ID:                bson.NewObjectID(),
			Name:              "English",
			LanguageCulture:   "en-US",
			UniqueSeoCode:     "en",
			FlagImageFileName: "us.png",
			Rtl:               false,
			LimitedToStores:   false,
			DefaultCurrencyID: bson.NewObjectID(),
			Published:         true,
			DisplayOrder:      1,
		},
		{
			ID:                bson.NewObjectID(),
			Name:              "Spanish",
			LanguageCulture:   "es-ES",
			UniqueSeoCode:     "es",
			FlagImageFileName: "es.png",
			Rtl:               false,
			LimitedToStores:   true,
			DefaultCurrencyID: bson.NewObjectID(),
			Published:         false,
			DisplayOrder:      2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedLanguages, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedLanguages, result)
	mockRepo.AssertExpectations(t)
}
