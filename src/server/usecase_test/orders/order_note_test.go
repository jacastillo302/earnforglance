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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestOrderNoteUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.OrderNoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderNoteUsecase(mockRepo, timeout)

	orderNoteID := primitive.NewObjectID().Hex()

	updatedOrderNote := domain.OrderNote{
		ID:                primitive.NewObjectID(), // Existing ID of the record to update
		OrderID:           primitive.NewObjectID(),
		Note:              "This is an updated note for the order.",
		DownloadID:        primitive.NewObjectID(),
		DisplayToCustomer: false,
		CreatedOnUtc:      time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, orderNoteID).Return(updatedOrderNote, nil)

	result, err := usecase.FetchByID(context.Background(), orderNoteID)

	assert.NoError(t, err)
	assert.Equal(t, updatedOrderNote, result)
	mockRepo.AssertExpectations(t)
}

func TestOrderNoteUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.OrderNoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderNoteUsecase(mockRepo, timeout)

	newOrderNote := &domain.OrderNote{
		OrderID:           primitive.NewObjectID(),
		Note:              "This is a note for the order.",
		DownloadID:        primitive.NewObjectID(),
		DisplayToCustomer: true,
		CreatedOnUtc:      time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newOrderNote).Return(nil)

	err := usecase.Create(context.Background(), newOrderNote)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderNoteUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.OrderNoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderNoteUsecase(mockRepo, timeout)

	updatedOrderNote := &domain.OrderNote{
		ID:                primitive.NewObjectID(), // Existing ID of the record to update
		OrderID:           primitive.NewObjectID(),
		Note:              "This is an updated note for the order.",
		DownloadID:        primitive.NewObjectID(),
		DisplayToCustomer: false,
		CreatedOnUtc:      time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("Update", mock.Anything, updatedOrderNote).Return(nil)

	err := usecase.Update(context.Background(), updatedOrderNote)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderNoteUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.OrderNoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderNoteUsecase(mockRepo, timeout)

	orderNoteID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, orderNoteID).Return(nil)

	err := usecase.Delete(context.Background(), orderNoteID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderNoteUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.OrderNoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderNoteUsecase(mockRepo, timeout)

	fetchedOrderNotes := []domain.OrderNote{
		{
			ID:                primitive.NewObjectID(),
			OrderID:           primitive.NewObjectID(),
			Note:              "This is a note for the order.",
			DownloadID:        primitive.NewObjectID(),
			DisplayToCustomer: true,
			CreatedOnUtc:      time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:                primitive.NewObjectID(),
			OrderID:           primitive.NewObjectID(),
			Note:              "This is another note for the order.",
			DownloadID:        primitive.NewObjectID(),
			DisplayToCustomer: false,
			CreatedOnUtc:      time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedOrderNotes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedOrderNotes, result)
	mockRepo.AssertExpectations(t)
}
