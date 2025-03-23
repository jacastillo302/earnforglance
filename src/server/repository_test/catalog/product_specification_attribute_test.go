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

type MockSingleResultProductSpecificationAttribute struct {
	mock.Mock
}

func (m *MockSingleResultProductSpecificationAttribute) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductSpecificationAttribute); ok {
		*v.(*domain.ProductSpecificationAttribute) = *result
	}
	return args.Error(1)
}

var mockItemProductSpecificationAttribute = &domain.ProductSpecificationAttribute{
	ID:                             primitive.NewObjectID(), // Existing ID of the record to update
	ProductID:                      primitive.NewObjectID(),
	AttributeTypeID:                primitive.NewObjectID(),
	SpecificationAttributeOptionID: primitive.NewObjectID(),
	CustomValue:                    "Updated Custom Value",
	AllowFiltering:                 false,
	ShowOnProductPage:              false,
	DisplayOrder:                   2,
	AttributeType:                  3,
}

func TestProductSpecificationAttributeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductSpecificationAttribute

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductSpecificationAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductSpecificationAttribute, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductSpecificationAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductSpecificationAttribute.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductSpecificationAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductSpecificationAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductSpecificationAttribute.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductSpecificationAttributeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductSpecificationAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductSpecificationAttribute).Return(nil, nil).Once()

	repo := repository.NewProductSpecificationAttributeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductSpecificationAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductSpecificationAttributeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductSpecificationAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductSpecificationAttribute.ID}
	update := bson.M{"$set": mockItemProductSpecificationAttribute}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductSpecificationAttributeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductSpecificationAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
