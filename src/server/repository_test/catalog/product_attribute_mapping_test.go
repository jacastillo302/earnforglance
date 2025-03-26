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

type MockSingleResultProductAttributeMapping struct {
	mock.Mock
}

func (m *MockSingleResultProductAttributeMapping) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductAttributeMapping); ok {
		*v.(*domain.ProductAttributeMapping) = *result
	}
	return args.Error(1)
}

var mockItem = &domain.ProductAttributeMapping{
	ID:                              primitive.NewObjectID(), // Existing ID of the record to update
	ProductID:                       primitive.NewObjectID(),
	ProductAttributeID:              primitive.NewObjectID(),
	TextPrompt:                      "Select a size",
	IsRequired:                      false,
	AttributeControlTypeID:          2,
	DisplayOrder:                    2,
	ValidationMinLength:             new(int),
	ValidationMaxLength:             new(int),
	ValidationFileAllowedExtensions: ".jpg,.png",
	ValidationFileMaximumSize:       new(int),
	DefaultValue:                    "Medium",
	ConditionAttributeXml:           "<attributes><size>medium</size></attributes>",
}

func TestProductAttributeMappingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductAttributeMapping

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAttributeMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAttributeMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAttributeMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAttributeMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductAttributeMappingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAttributeMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	collectionHelper.On("InsertOne", mock.Anything, mockItem).Return(nil, nil).Once()

	repo := repository.NewProductAttributeMappingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItem)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductAttributeMappingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAttributeMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	filter := bson.M{"_id": mockItem.ID}
	update := bson.M{"$set": mockItem}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductAttributeMappingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItem)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
