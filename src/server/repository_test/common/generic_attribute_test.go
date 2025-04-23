package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/common"
	repository "earnforglance/server/repository/common"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultGenericAttribute struct {
	mock.Mock
}

func (m *MockSingleResultGenericAttribute) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.GenericAttribute); ok {
		*v.(*domain.GenericAttribute) = *result
	}
	return args.Error(1)
}

var mockItemGenericAttribute = &domain.GenericAttribute{
	ID:                      bson.NewObjectID(), // Existing ID of the record to update
	EntityID:                bson.NewObjectID(),
	KeyGroup:                "Customer",
	Key:                     "PreferredLanguage",
	Value:                   "English",
	StoreID:                 bson.NewObjectID(),
	CreatedOrUpdatedDateUTC: new(time.Time),
}

func TestGenericAttributeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionGenericAttribute

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultGenericAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemGenericAttribute, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewGenericAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemGenericAttribute.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultGenericAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewGenericAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemGenericAttribute.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestGenericAttributeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionGenericAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemGenericAttribute).Return(nil, nil).Once()

	repo := repository.NewGenericAttributeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemGenericAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestGenericAttributeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionGenericAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemGenericAttribute.ID}
	update := bson.M{"$set": mockItemGenericAttribute}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewGenericAttributeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemGenericAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
