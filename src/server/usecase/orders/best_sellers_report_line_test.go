package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBestSellersReportLineUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.BestSellersReportLineRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewBestSellersReportLineUsecase(mockRepo, timeout)

	reportLineID := primitive.NewObjectID().Hex()

	updatedBestSellersReportLine := domain.BestSellersReportLine{
		ProductID:     primitive.NewObjectID(), // Existing ProductID of the record to update
		ProductName:   "Noise-Cancelling Headphones",
		TotalAmount:   2000.50,
		TotalQuantity: 75,
	}

	mockRepo.On("FetchByID", mock.Anything, reportLineID).Return(updatedBestSellersReportLine, nil)

	result, err := usecase.FetchByID(context.Background(), reportLineID)

	assert.NoError(t, err)
	assert.Equal(t, updatedBestSellersReportLine, result)
	mockRepo.AssertExpectations(t)
}

func TestBestSellersReportLineUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.BestSellersReportLineRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewBestSellersReportLineUsecase(mockRepo, timeout)

	newBestSellersReportLine := &domain.BestSellersReportLine{
		ProductID:     primitive.NewObjectID(),
		ProductName:   "Wireless Headphones",
		TotalAmount:   1500.75,
		TotalQuantity: 50,
	}

	mockRepo.On("Create", mock.Anything, newBestSellersReportLine).Return(nil)

	err := usecase.Create(context.Background(), newBestSellersReportLine)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBestSellersReportLineUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.BestSellersReportLineRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewBestSellersReportLineUsecase(mockRepo, timeout)

	updatedBestSellersReportLine := &domain.BestSellersReportLine{
		ProductID:     primitive.NewObjectID(), // Existing ProductID of the record to update
		ProductName:   "Noise-Cancelling Headphones",
		TotalAmount:   2000.50,
		TotalQuantity: 75,
	}

	mockRepo.On("Update", mock.Anything, updatedBestSellersReportLine).Return(nil)

	err := usecase.Update(context.Background(), updatedBestSellersReportLine)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBestSellersReportLineUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.BestSellersReportLineRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewBestSellersReportLineUsecase(mockRepo, timeout)

	reportLineID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, reportLineID).Return(nil)

	err := usecase.Delete(context.Background(), reportLineID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBestSellersReportLineUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.BestSellersReportLineRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewBestSellersReportLineUsecase(mockRepo, timeout)

	fetchedBestSellersReportLines := []domain.BestSellersReportLine{
		{
			ProductID:     primitive.NewObjectID(),
			ProductName:   "Wireless Headphones",
			TotalAmount:   1500.75,
			TotalQuantity: 50,
		},
		{
			ProductID:     primitive.NewObjectID(),
			ProductName:   "Bluetooth Speaker",
			TotalAmount:   1200.00,
			TotalQuantity: 30,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedBestSellersReportLines, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedBestSellersReportLines, result)
	mockRepo.AssertExpectations(t)
}
