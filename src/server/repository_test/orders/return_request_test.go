package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/orders"
	repository "earnforglance/server/repository/orders"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultReturnRequest struct {
	mock.Mock
}

func (m *MockSingleResultReturnRequest) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ReturnRequest); ok {
		*v.(*domain.ReturnRequest) = *result
	}
	return args.Error(1)
}

var mockItemReturnRequest = &domain.ReturnRequest{
	ID:                    bson.NewObjectID(), // Existing ID of the record to update
	CustomNumber:          "RR67890",
	StoreID:               bson.NewObjectID(),
	OrderItemID:           bson.NewObjectID(),
	CustomerID:            bson.NewObjectID(),
	Quantity:              1,
	ReturnedQuantity:      1,
	ReasonForReturn:       "Wrong Item Delivered",
	RequestedAction:       "Refund Item",
	CustomerComments:      "Received the wrong item.",
	UploadedFileID:        bson.NewObjectID(),
	StaffNotes:            "Process refund immediately.",
	ReturnRequestStatusID: 0,
	CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
	UpdatedOnUtc:          time.Now(),
}

func TestReturnRequestRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionReturnRequest

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultReturnRequest{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemReturnRequest, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewReturnRequestRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemReturnRequest.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultReturnRequest{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewReturnRequestRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemReturnRequest.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestReturnRequestRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionReturnRequest

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemReturnRequest).Return(nil, nil).Once()

	repo := repository.NewReturnRequestRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemReturnRequest)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestReturnRequestRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionReturnRequest

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemReturnRequest.ID}
	update := bson.M{"$set": mockItemReturnRequest}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewReturnRequestRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemReturnRequest)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
