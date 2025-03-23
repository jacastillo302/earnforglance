package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/catalog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultSpecificationAttribute struct {
	mock.Mock
}

func (m *MockSingleResultSpecificationAttribute) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.SpecificationAttribute); ok {
		*v.(*domain.SpecificationAttribute) = *result
	}
	return args.Error(1)
}

var mockItemSpecificationAttribute = &domain.SpecificationAttribute{
	ID:                            primitive.NewObjectID(), // Existing ID of the record to update
	Name:                          "Color",
	DisplayOrder:                  2,
	SpecificationAttributeGroupID: new(int),
}

func TestSpecificationAttributeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionSpecificationAttribute

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSpecificationAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemSpecificationAttribute, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSpecificationAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSpecificationAttribute.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSpecificationAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSpecificationAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSpecificationAttribute.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestSpecificationAttributeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSpecificationAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemSpecificationAttribute).Return(nil, nil).Once()

	repo := repository.NewSpecificationAttributeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemSpecificationAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestSpecificationAttributeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSpecificationAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemSpecificationAttribute.ID}
	update := bson.M{"$set": mockItemSpecificationAttribute}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewSpecificationAttributeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemSpecificationAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
