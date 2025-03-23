package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/orders"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultOrderNote struct {
	mock.Mock
}

func (m *MockSingleResultOrderNote) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.OrderNote); ok {
		*v.(*domain.OrderNote) = *result
	}
	return args.Error(1)
}

var mockItemOrderNote = &domain.OrderNote{
	ID:                primitive.NewObjectID(), // Existing ID of the record to update
	OrderID:           primitive.NewObjectID(),
	Note:              "This is an updated note for the order.",
	DownloadID:        primitive.NewObjectID(),
	DisplayToCustomer: false,
	CreatedOnUtc:      time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestOrderNoteRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionOrderNote

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultOrderNote{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemOrderNote, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewOrderNoteRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemOrderNote.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultOrderNote{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewOrderNoteRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemOrderNote.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestOrderNoteRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionOrderNote

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemOrderNote).Return(nil, nil).Once()

	repo := repository.NewOrderNoteRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemOrderNote)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestOrderNoteRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionOrderNote

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemOrderNote.ID}
	update := bson.M{"$set": mockItemOrderNote}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewOrderNoteRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemOrderNote)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
