package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/orders"
	repository "earnforglance/server/repository/orders"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultBestSellersReportLine struct {
	mock.Mock
}

func (m *MockSingleResultBestSellersReportLine) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.BestSellersReportLine); ok {
		*v.(*domain.BestSellersReportLine) = *result
	}
	return args.Error(1)
}

var mockItemBestSellersReportLine = &domain.BestSellersReportLine{
	ProductID:     primitive.NewObjectID(), // Existing ProductID of the record to update
	ProductName:   "Noise-Cancelling Headphones",
	TotalAmount:   2000.50,
	TotalQuantity: 75,
}

func TestBestSellersReportLineRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionBestSellersReportLine

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBestSellersReportLine{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemBestSellersReportLine, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBestSellersReportLineRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), string(mockItemBestSellersReportLine.ProductID.Hex()))

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBestSellersReportLine{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBestSellersReportLineRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemBestSellersReportLine.ProductID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestBestSellersReportLineRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBestSellersReportLine

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemBestSellersReportLine).Return(nil, nil).Once()

	repo := repository.NewBestSellersReportLineRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemBestSellersReportLine)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestBestSellersReportLineRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBestSellersReportLine

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemBestSellersReportLine.ProductID}
	update := bson.M{"$set": mockItemBestSellersReportLine}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewBestSellersReportLineRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemBestSellersReportLine)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
