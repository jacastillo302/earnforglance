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

func TestReturnRequestUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestRepository)
	timeout := time.Duration(10)
	usecase := NewReturnRequestUsecase(mockRepo, timeout)

	returnRequestID := primitive.NewObjectID().Hex()

	updatedReturnRequest := domain.ReturnRequest{
		ID:                    primitive.NewObjectID(), // Existing ID of the record to update
		CustomNumber:          "RR67890",
		StoreID:               primitive.NewObjectID(),
		OrderItemID:           primitive.NewObjectID(),
		CustomerID:            primitive.NewObjectID(),
		Quantity:              1,
		ReturnedQuantity:      1,
		ReasonForReturn:       "Wrong Item Delivered",
		RequestedAction:       "Refund Item",
		CustomerComments:      "Received the wrong item.",
		UploadedFileID:        primitive.NewObjectID(),
		StaffNotes:            "Process refund immediately.",
		ReturnRequestStatusID: primitive.NewObjectID(),
		CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
		UpdatedOnUtc:          time.Now(),
		ReturnRequestStatus:   3,
	}

	mockRepo.On("FetchByID", mock.Anything, returnRequestID).Return(updatedReturnRequest, nil)

	result, err := usecase.FetchByID(context.Background(), returnRequestID)

	assert.NoError(t, err)
	assert.Equal(t, updatedReturnRequest, result)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestRepository)
	timeout := time.Duration(10)
	usecase := NewReturnRequestUsecase(mockRepo, timeout)

	newReturnRequest := &domain.ReturnRequest{
		CustomNumber:          "RR12345",
		StoreID:               primitive.NewObjectID(),
		OrderItemID:           primitive.NewObjectID(),
		CustomerID:            primitive.NewObjectID(),
		Quantity:              2,
		ReturnedQuantity:      0,
		ReasonForReturn:       "Defective Item",
		RequestedAction:       "Replace Item",
		CustomerComments:      "The item is not working as expected.",
		UploadedFileID:        primitive.NewObjectID(),
		StaffNotes:            "Inspect the item upon return.",
		ReturnRequestStatusID: primitive.NewObjectID(),
		CreatedOnUtc:          time.Now(),
		UpdatedOnUtc:          time.Now(),
		ReturnRequestStatus:   2,
	}

	mockRepo.On("Create", mock.Anything, newReturnRequest).Return(nil)

	err := usecase.Create(context.Background(), newReturnRequest)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestRepository)
	timeout := time.Duration(10)
	usecase := NewReturnRequestUsecase(mockRepo, timeout)

	updatedReturnRequest := &domain.ReturnRequest{
		ID:                    primitive.NewObjectID(), // Existing ID of the record to update
		CustomNumber:          "RR67890",
		StoreID:               primitive.NewObjectID(),
		OrderItemID:           primitive.NewObjectID(),
		CustomerID:            primitive.NewObjectID(),
		Quantity:              1,
		ReturnedQuantity:      1,
		ReasonForReturn:       "Wrong Item Delivered",
		RequestedAction:       "Refund Item",
		CustomerComments:      "Received the wrong item.",
		UploadedFileID:        primitive.NewObjectID(),
		StaffNotes:            "Process refund immediately.",
		ReturnRequestStatusID: primitive.NewObjectID(),
		CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
		UpdatedOnUtc:          time.Now(),
		ReturnRequestStatus:   3,
	}

	mockRepo.On("Update", mock.Anything, updatedReturnRequest).Return(nil)

	err := usecase.Update(context.Background(), updatedReturnRequest)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestRepository)
	timeout := time.Duration(10)
	usecase := NewReturnRequestUsecase(mockRepo, timeout)

	returnRequestID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, returnRequestID).Return(nil)

	err := usecase.Delete(context.Background(), returnRequestID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestRepository)
	timeout := time.Duration(10)
	usecase := NewReturnRequestUsecase(mockRepo, timeout)

	fetchedReturnRequests := []domain.ReturnRequest{
		{
			ID:                    primitive.NewObjectID(),
			CustomNumber:          "RR12345",
			StoreID:               primitive.NewObjectID(),
			OrderItemID:           primitive.NewObjectID(),
			CustomerID:            primitive.NewObjectID(),
			Quantity:              2,
			ReturnedQuantity:      0,
			ReasonForReturn:       "Defective Item",
			RequestedAction:       "Replace Item",
			CustomerComments:      "The item is not working as expected.",
			UploadedFileID:        primitive.NewObjectID(),
			StaffNotes:            "Inspect the item upon return.",
			ReturnRequestStatusID: primitive.NewObjectID(),
			CreatedOnUtc:          time.Now().AddDate(0, 0, -10), // Created 10 days ago
			UpdatedOnUtc:          time.Now().AddDate(0, 0, -5),  // Updated 5 days ago
			ReturnRequestStatus:   2,
		},
		{
			ID:                    primitive.NewObjectID(),
			CustomNumber:          "RR67890",
			StoreID:               primitive.NewObjectID(),
			OrderItemID:           primitive.NewObjectID(),
			CustomerID:            primitive.NewObjectID(),
			Quantity:              1,
			ReturnedQuantity:      1,
			ReasonForReturn:       "Wrong Item Delivered",
			RequestedAction:       "Refund Item",
			CustomerComments:      "Received the wrong item.",
			UploadedFileID:        primitive.NewObjectID(),
			StaffNotes:            "Process refund immediately.",
			ReturnRequestStatusID: primitive.NewObjectID(),
			CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
			UpdatedOnUtc:          time.Now(),
			ReturnRequestStatus:   1,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedReturnRequests, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedReturnRequests, result)
	mockRepo.AssertExpectations(t)
}
