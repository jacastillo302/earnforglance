package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/common"
	repository "earnforglance/server/repository/common"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultAddressSettings struct {
	mock.Mock
}

func (m *MockSingleResultAddressSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.AddressSettings); ok {
		*v.(*domain.AddressSettings) = *result
	}
	return args.Error(1)
}

var mockItemAddressSettings = &domain.AddressSettings{
	ID:                        bson.NewObjectID(), // Existing ID of the record to update
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
	DefaultCountryID:          new(bson.ObjectID),
	StateProvinceEnabled:      false,
	PhoneEnabled:              true,
	PhoneRequired:             false,
	FaxEnabled:                true,
	FaxRequired:               true,
	PreselectCountryIfOnlyOne: false,
}

func TestAddressSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionAddressSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultAddressSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemAddressSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewAddressSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemAddressSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultAddressSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewAddressSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemAddressSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestAddressSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionAddressSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemAddressSettings).Return(nil, nil).Once()

	repo := repository.NewAddressSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemAddressSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestAddressSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionAddressSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemAddressSettings.ID}
	update := bson.M{"$set": mockItemAddressSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewAddressSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemAddressSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
