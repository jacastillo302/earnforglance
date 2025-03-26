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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultSalesSummaryReportLine struct {
	mock.Mock
}

func (m *MockSingleResultSalesSummaryReportLine) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.SalesSummaryReportLine); ok {
		*v.(*domain.SalesSummaryReportLine) = *result
	}
	return args.Error(1)
}

var mockItemSalesSummaryReportLine = &domain.SalesSummaryReportLine{
	ID:             primitive.NewObjectID(), // Existing ID of the record to update
	Summary:        "Updated Weekly Sales Summary",
	SummaryDate:    time.Now().AddDate(0, 0, -7), // Summary from 7 days ago
	NumberOfOrders: 100,
	Profit:         10000.00,
	ProfitStr:      "$10000.00",
	Shipping:       "$400.00",
	Tax:            "$600.00",
	OrderTotal:     "$11000.00",
	SummaryType:    2,
}

func TestSalesSummaryReportLineRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionSalesSummaryReportLine

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSalesSummaryReportLine{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemSalesSummaryReportLine, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSalesSummaryReportLineRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSalesSummaryReportLine.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSalesSummaryReportLine{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSalesSummaryReportLineRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSalesSummaryReportLine.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestSalesSummaryReportLineRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSalesSummaryReportLine

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemSalesSummaryReportLine).Return(nil, nil).Once()

	repo := repository.NewSalesSummaryReportLineRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemSalesSummaryReportLine)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestSalesSummaryReportLineRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSalesSummaryReportLine

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemSalesSummaryReportLine.ID}
	update := bson.M{"$set": mockItemSalesSummaryReportLine}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewSalesSummaryReportLineRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemSalesSummaryReportLine)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
