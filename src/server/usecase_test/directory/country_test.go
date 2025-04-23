package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/directory"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/directory"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestCountryUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CountryRepository)
	timeout := time.Duration(10)
	usecase := test.NewCountryUsecase(mockRepo, timeout)

	countryID := bson.NewObjectID().Hex()

	updatedCountry := domain.Country{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		Name:               "Canada",
		AllowsBilling:      true,
		AllowsShipping:     true,
		TwoLetterIsoCode:   "CA",
		ThreeLetterIsoCode: "CAN",
		NumericIsoCode:     124,
		SubjectToVat:       true,
		Published:          false,
		DisplayOrder:       2,
		LimitedToStores:    true,
	}

	mockRepo.On("FetchByID", mock.Anything, countryID).Return(updatedCountry, nil)

	result, err := usecase.FetchByID(context.Background(), countryID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCountry, result)
	mockRepo.AssertExpectations(t)
}

func TestCountryUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CountryRepository)
	timeout := time.Duration(10)
	usecase := test.NewCountryUsecase(mockRepo, timeout)

	newCountry := &domain.Country{
		Name:               "United States",
		AllowsBilling:      true,
		AllowsShipping:     true,
		TwoLetterIsoCode:   "US",
		ThreeLetterIsoCode: "USA",
		NumericIsoCode:     840,
		SubjectToVat:       false,
		Published:          true,
		DisplayOrder:       1,
		LimitedToStores:    false,
	}

	mockRepo.On("Create", mock.Anything, newCountry).Return(nil)

	err := usecase.Create(context.Background(), newCountry)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCountryUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CountryRepository)
	timeout := time.Duration(10)
	usecase := test.NewCountryUsecase(mockRepo, timeout)

	updatedCountry := &domain.Country{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		Name:               "Canada",
		AllowsBilling:      true,
		AllowsShipping:     true,
		TwoLetterIsoCode:   "CA",
		ThreeLetterIsoCode: "CAN",
		NumericIsoCode:     124,
		SubjectToVat:       true,
		Published:          false,
		DisplayOrder:       2,
		LimitedToStores:    true,
	}

	mockRepo.On("Update", mock.Anything, updatedCountry).Return(nil)

	err := usecase.Update(context.Background(), updatedCountry)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCountryUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CountryRepository)
	timeout := time.Duration(10)
	usecase := test.NewCountryUsecase(mockRepo, timeout)

	countryID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, countryID).Return(nil)

	err := usecase.Delete(context.Background(), countryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCountryUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CountryRepository)
	timeout := time.Duration(10)
	usecase := test.NewCountryUsecase(mockRepo, timeout)

	fetchedCountries := []domain.Country{
		{
			ID:                 bson.NewObjectID(),
			Name:               "United States",
			AllowsBilling:      true,
			AllowsShipping:     true,
			TwoLetterIsoCode:   "US",
			ThreeLetterIsoCode: "USA",
			NumericIsoCode:     840,
			SubjectToVat:       false,
			Published:          true,
			DisplayOrder:       1,
			LimitedToStores:    false,
		},
		{
			ID:                 bson.NewObjectID(),
			Name:               "Canada",
			AllowsBilling:      true,
			AllowsShipping:     true,
			TwoLetterIsoCode:   "CA",
			ThreeLetterIsoCode: "CAN",
			NumericIsoCode:     124,
			SubjectToVat:       true,
			Published:          false,
			DisplayOrder:       2,
			LimitedToStores:    true,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCountries, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCountries, result)
	mockRepo.AssertExpectations(t)
}
