package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/catalog"
	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultSpecificationAttributeGroup struct {
	mock.Mock
}

func (m *MockSingleResultSpecificationAttributeGroup) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.SpecificationAttributeGroup); ok {
		*v.(*domain.SpecificationAttributeGroup) = *result
	}
	return args.Error(1)
}

var mockItemSpecificationAttributeGroup = &domain.SpecificationAttributeGroup{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	Name:         "Technical Specifications",
	DisplayOrder: 2,
}

func TestSpecificationAttributeGroupRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionSpecificationAttributeGroup

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSpecificationAttributeGroup{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemSpecificationAttributeGroup, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSpecificationAttributeGroupRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSpecificationAttributeGroup.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSpecificationAttributeGroup{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSpecificationAttributeGroupRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSpecificationAttributeGroup.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestSpecificationAttributeGroupRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSpecificationAttributeGroup

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemSpecificationAttributeGroup).Return(nil, nil).Once()

	repo := repository.NewSpecificationAttributeGroupRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemSpecificationAttributeGroup)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestSpecificationAttributeGroupRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSpecificationAttributeGroup

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemSpecificationAttributeGroup.ID}
	update := bson.M{"$set": mockItemSpecificationAttributeGroup}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewSpecificationAttributeGroupRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemSpecificationAttributeGroup)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
