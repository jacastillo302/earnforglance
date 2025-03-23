package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/customers"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/customers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBestCustomerReportLineUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.BestCustomerReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewBestCustomerReportLineUsecase(mockRepo, timeout)

	customerID := primitive.NewObjectID().Hex()

	updatedBestCustomerReportLine := domain.BestCustomerReportLine{
		ID:         primitive.NewObjectID(), // Existing ID of the record to update
		CustomerID: primitive.NewObjectID(),
		OrderTotal: 2000.50,
		OrderCount: 15,
	}

	mockRepo.On("FetchByID", mock.Anything, customerID).Return(updatedBestCustomerReportLine, nil)

	result, err := usecase.FetchByID(context.Background(), customerID)

	assert.NoError(t, err)
	assert.Equal(t, updatedBestCustomerReportLine, result)
	mockRepo.AssertExpectations(t)
}

func TestBestCustomerReportLineUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.BestCustomerReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewBestCustomerReportLineUsecase(mockRepo, timeout)

	newBestCustomerReportLine := &domain.BestCustomerReportLine{
		CustomerID: primitive.NewObjectID(),
		OrderTotal: 1500.75,
		OrderCount: 10,
	}

	mockRepo.On("Create", mock.Anything, newBestCustomerReportLine).Return(nil)

	err := usecase.Create(context.Background(), newBestCustomerReportLine)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBestCustomerReportLineUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.BestCustomerReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewBestCustomerReportLineUsecase(mockRepo, timeout)

	updatedBestCustomerReportLine := &domain.BestCustomerReportLine{
		ID:         primitive.NewObjectID(), // Existing ID of the record to update
		CustomerID: primitive.NewObjectID(),
		OrderTotal: 2000.50,
		OrderCount: 15,
	}

	mockRepo.On("Update", mock.Anything, updatedBestCustomerReportLine).Return(nil)

	err := usecase.Update(context.Background(), updatedBestCustomerReportLine)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBestCustomerReportLineUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.BestCustomerReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewBestCustomerReportLineUsecase(mockRepo, timeout)

	customerID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customerID).Return(nil)

	err := usecase.Delete(context.Background(), customerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBestCustomerReportLineUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.BestCustomerReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewBestCustomerReportLineUsecase(mockRepo, timeout)

	fetchedBestCustomerReportLines := []domain.BestCustomerReportLine{
		{
			ID:         primitive.NewObjectID(),
			CustomerID: primitive.NewObjectID(),
			OrderTotal: 1500.75,
			OrderCount: 10,
		},
		{
			ID:         primitive.NewObjectID(),
			CustomerID: primitive.NewObjectID(),
			OrderTotal: 2000.50,
			OrderCount: 15,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedBestCustomerReportLines, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedBestCustomerReportLines, result)
	mockRepo.AssertExpectations(t)
}
