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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultBaseAttribute struct {
	mock.Mock
}

func (m *MockSingleResultBaseAttribute) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.BaseAttribute); ok {
		*v.(*domain.BaseAttribute) = *result
	}
	return args.Error(1)
}

func TestBaseAttributeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionBaseAttribute

	mockItem := domain.BaseAttribute{
		ID:                     primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		Name:                   "",
		IsRequired:             false,
		AttributeControlTypeId: 0,
		DisplayOrder:           0,
	}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBaseAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBaseAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		//assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBaseAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBaseAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestBaseAttributeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBaseAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBaseAttribute := &domain.BaseAttribute{
		ID:                     primitive.NewObjectID(),
		Name:                   "Test Attribute",
		IsRequired:             true,
		AttributeControlTypeId: 1, // Example: TextBox
		DisplayOrder:           5,
	}

	collectionHelper.On("InsertOne", mock.Anything, mockBaseAttribute).Return(nil, nil).Once()

	repo := repository.NewBaseAttributeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockBaseAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestBaseAttributeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBaseAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBaseAttribute := &domain.BaseAttribute{
		ID:                     primitive.NewObjectID(),
		Name:                   "Test Attribute",
		IsRequired:             true,
		AttributeControlTypeId: 1, // Example: TextBox
		DisplayOrder:           5,
	}

	filter := bson.M{"_id": mockBaseAttribute.ID}
	update := bson.M{"$set": mockBaseAttribute}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewBaseAttributeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockBaseAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
