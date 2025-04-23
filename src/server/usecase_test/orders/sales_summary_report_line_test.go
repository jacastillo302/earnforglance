package usecase_test

import (
	"context"
	"testing"
	"time"

	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	test "earnforglance/server/usecase/orders"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestSalesSummaryReportLineUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.SalesSummaryReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewSalesSummaryReportLineUsecase(mockRepo, timeout)

	orderID := bson.NewObjectID().Hex()

	updatedSalesSummaryReportLine := domain.SalesSummaryReportLine{
		ID:             bson.NewObjectID(), // Existing ID of the record to update
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

	mockRepo.On("FetchByID", mock.Anything, orderID).Return(updatedSalesSummaryReportLine, nil)

	result, err := usecase.FetchByID(context.Background(), orderID)

	assert.NoError(t, err)
	assert.Equal(t, updatedSalesSummaryReportLine, result)
	mockRepo.AssertExpectations(t)
}

func TestSalesSummaryReportLineUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.SalesSummaryReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewSalesSummaryReportLineUsecase(mockRepo, timeout)

	newSalesSummaryReportLine := &domain.SalesSummaryReportLine{
		Summary:        "Daily Sales Summary",
		SummaryDate:    time.Now(),
		NumberOfOrders: 50,
		Profit:         5000.00,
		ProfitStr:      "$5000.00",
		Shipping:       "$200.00",
		Tax:            "$300.00",
		OrderTotal:     "$5500.00",
		SummaryType:    1,
	}

	mockRepo.On("Create", mock.Anything, newSalesSummaryReportLine).Return(nil)

	err := usecase.Create(context.Background(), newSalesSummaryReportLine)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSalesSummaryReportLineUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.SalesSummaryReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewSalesSummaryReportLineUsecase(mockRepo, timeout)
	updatedSalesSummaryReportLine := &domain.SalesSummaryReportLine{
		ID:             bson.NewObjectID(), // Existing ID of the record to update
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

	mockRepo.On("Update", mock.Anything, updatedSalesSummaryReportLine).Return(nil)

	err := usecase.Update(context.Background(), updatedSalesSummaryReportLine)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSalesSummaryReportLineUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.SalesSummaryReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewSalesSummaryReportLineUsecase(mockRepo, timeout)

	orderID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, orderID).Return(nil)

	err := usecase.Delete(context.Background(), orderID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSalesSummaryReportLineUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.SalesSummaryReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewSalesSummaryReportLineUsecase(mockRepo, timeout)

	fetchedSalesSummaryReportLines := []domain.SalesSummaryReportLine{
		{
			ID:             bson.NewObjectID(),
			Summary:        "Daily Sales Summary",
			SummaryDate:    time.Now().AddDate(0, 0, -1), // Summary from yesterday
			NumberOfOrders: 50,
			Profit:         5000.00,
			ProfitStr:      "$5000.00",
			Shipping:       "$200.00",
			Tax:            "$300.00",
			OrderTotal:     "$5500.00",
			SummaryType:    1,
		},
		{
			ID:             bson.NewObjectID(),
			Summary:        "Weekly Sales Summary",
			SummaryDate:    time.Now().AddDate(0, 0, -7), // Summary from 7 days ago
			NumberOfOrders: 100,
			Profit:         10000.00,
			ProfitStr:      "$10000.00",
			Shipping:       "$400.00",
			Tax:            "$600.00",
			OrderTotal:     "$11000.00",
			SummaryType:    2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedSalesSummaryReportLines, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedSalesSummaryReportLines, result)
	mockRepo.AssertExpectations(t)
}
