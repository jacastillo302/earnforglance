package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultAddress struct {
	mock.Mock
}

func (m *MockSingleResultAddress) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Address); ok {
		*v.(*domain.Address) = *result
	}
	return args.Error(1)
}

var mockItemAddress = &domain.Address{
	ID:               primitive.NewObjectID(), // Existing ID of the record to update
	FirstName:        "Jane",
	LastName:         "Smith",
	Email:            "jane.smith@example.com",
	Company:          "Updated Corp",
	CountryID:        new(primitive.ObjectID),
	StateProvinceID:  new(primitive.ObjectID),
	County:           "Updated County",
	City:             "Updated City",
	Address1:         "456 Elm St",
	Address2:         "Suite 101",
	ZipPostalCode:    "67890",
	PhoneNumber:      "987-654-3210",
	FaxNumber:        "987-654-3211",
	CustomAttributes: "<custom><attribute>updated</attribute></custom>",
	CreatedOnUtc:     time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestAddressRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionAddress

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultAddress{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemAddress, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewAddressRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemAddress.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultAddress{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewAddressRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemAddress.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestAddressRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionAddress

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemAddress).Return(nil, nil).Once()

	repo := repository.NewAddressRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemAddress)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestAddressRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionAddress

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemAddress.ID}
	update := bson.M{"$set": mockItemAddress}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewAddressRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemAddress)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
