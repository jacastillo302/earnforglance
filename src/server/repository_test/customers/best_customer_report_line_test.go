package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/customers"
	repository "earnforglance/server/repository/customers"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultBestCustomerReportLine struct {
	mock.Mock
}

func (m *MockSingleResultBestCustomerReportLine) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.BestCustomerReportLine); ok {
		*v.(*domain.BestCustomerReportLine) = *result
	}
	return args.Error(1)
}

var mockItemBestCustomerReportLine = &domain.BestCustomerReportLine{
	ID:         primitive.NewObjectID(), // Existing ID of the record to update
	CustomerID: primitive.NewObjectID(),
	OrderTotal: 2000.50,
	OrderCount: 15,
}

func TestBestCustomerReportLineRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionBestCustomerReportLine

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBestCustomerReportLine{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemBestCustomerReportLine, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBestCustomerReportLineRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemBestCustomerReportLine.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBestCustomerReportLine{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBestCustomerReportLineRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemBestCustomerReportLine.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestBestCustomerReportLineRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBestCustomerReportLine

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemBestCustomerReportLine).Return(nil, nil).Once()

	repo := repository.NewBestCustomerReportLineRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemBestCustomerReportLine)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestBestCustomerReportLineRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBestCustomerReportLine

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemBestCustomerReportLine.ID}
	update := bson.M{"$set": mockItemBestCustomerReportLine}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewBestCustomerReportLineRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemBestCustomerReportLine)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
