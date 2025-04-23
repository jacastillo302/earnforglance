package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/seo"
	repository "earnforglance/server/repository/seo"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultUrlRecord struct {
	mock.Mock
}

func (m *MockSingleResultUrlRecord) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.UrlRecord); ok {
		*v.(*domain.UrlRecord) = *result
	}
	return args.Error(1)
}

var mockItemUrlRecord = &domain.UrlRecord{
	ID:                 bson.NewObjectID(), // Existing ID of the record to update
	PermissionRecordID: bson.NewObjectID(),
	Slug:               "example-category",
	IsActive:           false,
}

func TestUrlRecordRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionUrlRecord

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultUrlRecord{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemUrlRecord, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewUrlRecordRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemUrlRecord.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultUrlRecord{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewUrlRecordRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemUrlRecord.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestUrlRecordRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionUrlRecord

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemUrlRecord).Return(nil, nil).Once()

	repo := repository.NewUrlRecordRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemUrlRecord)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestUrlRecordRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionUrlRecord

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemUrlRecord.ID}
	update := bson.M{"$set": mockItemUrlRecord}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewUrlRecordRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemUrlRecord)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
