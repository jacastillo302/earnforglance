package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/vendors"
	repository "earnforglance/server/repository/vendors"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultVendorSettings struct {
	mock.Mock
}

func (m *MockSingleResultVendorSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.VendorSettings); ok {
		*v.(*domain.VendorSettings) = *result
	}
	return args.Error(1)
}

var mockItemVendorSettings = &domain.VendorSettings{
	ID:                                           bson.NewObjectID(), // Existing ID of the record to update
	DefaultVendorPageSizeOptions:                 "5,15,25",
	VendorsBlockItemsToDisplay:                   15,
	ShowVendorOnProductDetailsPage:               false,
	ShowVendorOnOrderDetailsPage:                 false,
	AllowCustomersToContactVendors:               false,
	AllowCustomersToApplyForVendorAccount:        false,
	TermsOfServiceEnabled:                        true,
	AllowSearchByVendor:                          false,
	AllowVendorsToEditInfo:                       false,
	NotifyStoreOwnerAboutVendorInformationChange: false,
	MaximumProductNumber:                         50,
	AllowVendorsToImportProducts:                 true,
	MaximumProductPicturesNumber:                 10,
}

func TestVendorSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionVendorSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultVendorSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemVendorSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewVendorSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemVendorSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultVendorSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewVendorSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemVendorSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestVendorSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionVendorSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemVendorSettings).Return(nil, nil).Once()

	repo := repository.NewVendorSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemVendorSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestVendorSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionVendorSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemVendorSettings.ID}
	update := bson.M{"$set": mockItemVendorSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewVendorSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemVendorSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
