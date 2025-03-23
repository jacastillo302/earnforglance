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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddressSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.AddressSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewAddressSettingsUsecase(mockRepo, timeout)

	addressSettingsID := primitive.NewObjectID().Hex()

	updatedAddressSettings := domain.AddressSettings{
		ID:                        primitive.NewObjectID(), // Existing ID of the record to update
		CompanyEnabled:            false,
		CompanyRequired:           false,
		StreetAddressEnabled:      true,
		StreetAddressRequired:     false,
		StreetAddress2Enabled:     true,
		StreetAddress2Required:    true,
		ZipPostalCodeEnabled:      true,
		ZipPostalCodeRequired:     false,
		CityEnabled:               true,
		CityRequired:              false,
		CountyEnabled:             true,
		CountyRequired:            true,
		CountryEnabled:            true,
		DefaultCountryID:          new(primitive.ObjectID),
		StateProvinceEnabled:      false,
		PhoneEnabled:              true,
		PhoneRequired:             false,
		FaxEnabled:                true,
		FaxRequired:               true,
		PreselectCountryIfOnlyOne: false,
	}

	mockRepo.On("FetchByID", mock.Anything, addressSettingsID).Return(updatedAddressSettings, nil)

	result, err := usecase.FetchByID(context.Background(), addressSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedAddressSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestAddressSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.AddressSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewAddressSettingsUsecase(mockRepo, timeout)

	newAddressSettings := &domain.AddressSettings{
		CompanyEnabled:            true,
		CompanyRequired:           false,
		StreetAddressEnabled:      true,
		StreetAddressRequired:     true,
		StreetAddress2Enabled:     false,
		StreetAddress2Required:    false,
		ZipPostalCodeEnabled:      true,
		ZipPostalCodeRequired:     true,
		CityEnabled:               true,
		CityRequired:              true,
		CountyEnabled:             false,
		CountyRequired:            false,
		CountryEnabled:            true,
		DefaultCountryID:          nil,
		StateProvinceEnabled:      true,
		PhoneEnabled:              true,
		PhoneRequired:             true,
		FaxEnabled:                false,
		FaxRequired:               false,
		PreselectCountryIfOnlyOne: true,
	}

	mockRepo.On("Create", mock.Anything, newAddressSettings).Return(nil)

	err := usecase.Create(context.Background(), newAddressSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddressSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.AddressSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewAddressSettingsUsecase(mockRepo, timeout)

	updatedAddressSettings := &domain.AddressSettings{
		ID:                        primitive.NewObjectID(), // Existing ID of the record to update
		CompanyEnabled:            false,
		CompanyRequired:           false,
		StreetAddressEnabled:      true,
		StreetAddressRequired:     false,
		StreetAddress2Enabled:     true,
		StreetAddress2Required:    true,
		ZipPostalCodeEnabled:      true,
		ZipPostalCodeRequired:     false,
		CityEnabled:               true,
		CityRequired:              false,
		CountyEnabled:             true,
		CountyRequired:            true,
		CountryEnabled:            true,
		DefaultCountryID:          new(primitive.ObjectID),
		StateProvinceEnabled:      false,
		PhoneEnabled:              true,
		PhoneRequired:             false,
		FaxEnabled:                true,
		FaxRequired:               true,
		PreselectCountryIfOnlyOne: false,
	}
	*updatedAddressSettings.DefaultCountryID = primitive.NewObjectID()

	mockRepo.On("Update", mock.Anything, updatedAddressSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedAddressSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddressSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.AddressSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewAddressSettingsUsecase(mockRepo, timeout)

	addressSettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, addressSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), addressSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddressSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.AddressSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewAddressSettingsUsecase(mockRepo, timeout)

	fetchedAddressSettings := []domain.AddressSettings{
		{
			ID:                        primitive.NewObjectID(),
			CompanyEnabled:            true,
			CompanyRequired:           false,
			StreetAddressEnabled:      true,
			StreetAddressRequired:     true,
			StreetAddress2Enabled:     false,
			StreetAddress2Required:    false,
			ZipPostalCodeEnabled:      true,
			ZipPostalCodeRequired:     true,
			CityEnabled:               true,
			CityRequired:              true,
			CountyEnabled:             false,
			CountyRequired:            false,
			CountryEnabled:            true,
			DefaultCountryID:          nil,
			StateProvinceEnabled:      true,
			PhoneEnabled:              true,
			PhoneRequired:             true,
			FaxEnabled:                false,
			FaxRequired:               false,
			PreselectCountryIfOnlyOne: true,
		},
		{
			ID:                        primitive.NewObjectID(),
			CompanyEnabled:            false,
			CompanyRequired:           false,
			StreetAddressEnabled:      true,
			StreetAddressRequired:     false,
			StreetAddress2Enabled:     true,
			StreetAddress2Required:    true,
			ZipPostalCodeEnabled:      true,
			ZipPostalCodeRequired:     false,
			CityEnabled:               true,
			CityRequired:              false,
			CountyEnabled:             true,
			CountyRequired:            true,
			CountryEnabled:            true,
			DefaultCountryID:          new(primitive.ObjectID),
			StateProvinceEnabled:      false,
			PhoneEnabled:              true,
			PhoneRequired:             false,
			FaxEnabled:                true,
			FaxRequired:               true,
			PreselectCountryIfOnlyOne: false,
		},
	}
	*fetchedAddressSettings[1].DefaultCountryID = primitive.NewObjectID()

	mockRepo.On("Fetch", mock.Anything).Return(fetchedAddressSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedAddressSettings, result)
	mockRepo.AssertExpectations(t)
}
