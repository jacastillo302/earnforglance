package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/attributes"
	repository "earnforglance/server/repository/attributes"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultPermisionRecordAttributeValue struct {
	mock.Mock
}

func (m *MockSingleResultPermisionRecordAttributeValue) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PermisionRecordAttributeValue); ok {
		*v.(*domain.PermisionRecordAttributeValue) = *result
	}
	return args.Error(1)
}

var mockItemPermisionRecordAttributeValue = &domain.PermisionRecordAttributeValue{
	ID:                         bson.NewObjectID(), // Existing ID of the record to update
	PermisionRecordAttributeID: bson.NewObjectID(),
	Value:                      "Preferred Currency",
	IsPreSelected:              false,
	DisplayOrder:               2,
}

func TestPermisionRecordAttributeValueRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPermisionRecordAttributeValue

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPermisionRecordAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPermisionRecordAttributeValue, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPermisionRecordAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPermisionRecordAttributeValue.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPermisionRecordAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPermisionRecordAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPermisionRecordAttributeValue.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPermisionRecordAttributeValueRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPermisionRecordAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPermisionRecordAttributeValue).Return(nil, nil).Once()

	repo := repository.NewPermisionRecordAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPermisionRecordAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPermisionRecordAttributeValueRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPermisionRecordAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPermisionRecordAttributeValue.ID}
	update := bson.M{"$set": mockItemPermisionRecordAttributeValue}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPermisionRecordAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPermisionRecordAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
