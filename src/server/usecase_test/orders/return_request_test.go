package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	test "earnforglance/server/usecase/orders"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestReturnRequestUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestRepository)
	timeout := time.Duration(10)
	usecase := test.NewReturnRequestUsecase(mockRepo, timeout)

	returnRequestID := bson.NewObjectID().Hex()

	updatedReturnRequest := domain.ReturnRequest{
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
		ReturnRequestStatusID: 10,
		CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
		UpdatedOnUtc:          time.Now(),
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
	usecase := test.NewReturnRequestUsecase(mockRepo, timeout)

	newReturnRequest := &domain.ReturnRequest{
		CustomNumber:          "RR12345",
		StoreID:               bson.NewObjectID(),
		OrderItemID:           bson.NewObjectID(),
		CustomerID:            bson.NewObjectID(),
		Quantity:              2,
		ReturnedQuantity:      0,
		ReasonForReturn:       "Defective Item",
		RequestedAction:       "Replace Item",
		CustomerComments:      "The item is not working as expected.",
		UploadedFileID:        bson.NewObjectID(),
		StaffNotes:            "Inspect the item upon return.",
		ReturnRequestStatusID: 40,
		CreatedOnUtc:          time.Now(),
		UpdatedOnUtc:          time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newReturnRequest).Return(nil)

	err := usecase.Create(context.Background(), newReturnRequest)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestRepository)
	timeout := time.Duration(10)
	usecase := test.NewReturnRequestUsecase(mockRepo, timeout)

	updatedReturnRequest := &domain.ReturnRequest{
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
		ReturnRequestStatusID: 30,
		CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
		UpdatedOnUtc:          time.Now(),
	}

	mockRepo.On("Update", mock.Anything, updatedReturnRequest).Return(nil)

	err := usecase.Update(context.Background(), updatedReturnRequest)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestRepository)
	timeout := time.Duration(10)
	usecase := test.NewReturnRequestUsecase(mockRepo, timeout)

	returnRequestID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, returnRequestID).Return(nil)

	err := usecase.Delete(context.Background(), returnRequestID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestRepository)
	timeout := time.Duration(10)
	usecase := test.NewReturnRequestUsecase(mockRepo, timeout)

	fetchedReturnRequests := []domain.ReturnRequest{
		{
			ID:                    bson.NewObjectID(),
			CustomNumber:          "RR12345",
			StoreID:               bson.NewObjectID(),
			OrderItemID:           bson.NewObjectID(),
			CustomerID:            bson.NewObjectID(),
			Quantity:              2,
			ReturnedQuantity:      0,
			ReasonForReturn:       "Defective Item",
			RequestedAction:       "Replace Item",
			CustomerComments:      "The item is not working as expected.",
			UploadedFileID:        bson.NewObjectID(),
			StaffNotes:            "Inspect the item upon return.",
			ReturnRequestStatusID: 20,
			CreatedOnUtc:          time.Now().AddDate(0, 0, -10), // Created 10 days ago
			UpdatedOnUtc:          time.Now().AddDate(0, 0, -5),  // Updated 5 days ago
		},
		{
			ID:                    bson.NewObjectID(),
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
			ReturnRequestStatusID: 10,
			CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
			UpdatedOnUtc:          time.Now(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedReturnRequests, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedReturnRequests, result)
	mockRepo.AssertExpectations(t)
}
