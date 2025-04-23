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

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultManufacturerTemplate struct {
	mock.Mock
}

func (m *MockSingleResultManufacturerTemplate) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ManufacturerTemplate); ok {
		*v.(*domain.ManufacturerTemplate) = *result
	}
	return args.Error(1)
}

func TestManufacturerTemplateRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionManufacturerTemplate

	mockItem := domain.ManufacturerTemplate{ID: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, Name: "", ViewPath: "", DisplayOrder: 0}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultManufacturerTemplate{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewManufacturerTemplateRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultManufacturerTemplate{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewManufacturerTemplateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestManufacturerTemplateRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionManufacturerTemplate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockManufacturerTemplate := &domain.ManufacturerTemplate{
		ID:           bson.NewObjectID(),
		Name:         "Updated Manufacturer Template",
		ViewPath:     "/Views/Manufacturer/Updated.cshtml",
		DisplayOrder: 2,
	}

	collectionHelper.On("InsertOne", mock.Anything, mockManufacturerTemplate).Return(nil, nil).Once()

	repo := repository.NewManufacturerTemplateRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockManufacturerTemplate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestManufacturerTemplateRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionManufacturerTemplate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockManufacturerTemplate := &domain.ManufacturerTemplate{
		ID:           bson.NewObjectID(),
		Name:         "Updated Manufacturer Template",
		ViewPath:     "/Views/Manufacturer/Updated.cshtml",
		DisplayOrder: 2,
	}

	filter := bson.M{"_id": mockManufacturerTemplate.ID}
	update := bson.M{"$set": mockManufacturerTemplate}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewManufacturerTemplateRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockManufacturerTemplate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
