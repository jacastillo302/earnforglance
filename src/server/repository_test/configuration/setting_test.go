package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/configuration"
	repository "earnforglance/server/repository/configuration"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultSetting struct {
	mock.Mock
}

func (m *MockSingleResultSetting) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Setting); ok {
		*v.(*domain.Setting) = *result
	}
	return args.Error(1)
}

var mockItemSetting = &domain.Setting{
	ID:      primitive.NewObjectID(), // Existing ID of the record to update
	Name:    "SiteTitle",
	Value:   "Updated E-Commerce Store",
	StoreID: primitive.NewObjectID(),
}

func TestSettingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionSetting

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSetting{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemSetting, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSettingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSetting.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSetting{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSettingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSetting.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestSettingRepository_FetchByName(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionSetting

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSetting{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemSetting, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSettingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByName(context.Background(), mockItemSetting.Name)

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSetting{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSettingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByName(context.Background(), mockItemSetting.Name)

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestSettingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSetting

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemSetting).Return(nil, nil).Once()

	repo := repository.NewSettingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemSetting)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestSettingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSetting

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemSetting.ID}
	update := bson.M{"$set": mockItemSetting}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewSettingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemSetting)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestFetchByNames(t *testing.T) {
	// Setup in-memory MongoDB
	db := &mocks.Database{}
	collection := domain.CollectionSetting
	repo := repository.NewSettingRepository(db, collection)

	// Insert test data
	testSettings := []domain.Setting{
		{Name: "Setting1"},
		{Name: "Setting2"},
		{Name: "Setting3"},
	}
	for _, setting := range testSettings {
		_, _ = db.Collection(collection).InsertOne(context.Background(), setting)
	}

	// Test FetchByNames
	names := []string{"Setting1", "Setting3"}
	result, err := repo.FetchByNames(context.Background(), names)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Setting1", result[0].Name)
	assert.Equal(t, "Setting3", result[1].Name)
}
