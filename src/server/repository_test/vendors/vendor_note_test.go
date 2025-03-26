package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/vendors"
	repository "earnforglance/server/repository/vendors"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultVendorNote struct {
	mock.Mock
}

func (m *MockSingleResultVendorNote) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.VendorNote); ok {
		*v.(*domain.VendorNote) = *result
	}
	return args.Error(1)
}

var mockItemVendorNote = &domain.VendorNote{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	VendorID:     primitive.NewObjectID(),
	Note:         "This is an updated note for the vendor.",
	CreatedOnUtc: time.Now().AddDate(0, 0, -1), // Created 1 day ago
}

func TestVendorNoteRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionVendorNote

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultVendorNote{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemVendorNote, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewVendorNoteRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemVendorNote.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultVendorNote{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewVendorNoteRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemVendorNote.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestVendorNoteRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionVendorNote

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemVendorNote).Return(nil, nil).Once()

	repo := repository.NewVendorNoteRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemVendorNote)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestVendorNoteRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionVendorNote

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemVendorNote.ID}
	update := bson.M{"$set": mockItemVendorNote}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewVendorNoteRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemVendorNote)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
