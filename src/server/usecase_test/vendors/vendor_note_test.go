package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/vendors"
	test "earnforglance/server/usecase/vendors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestVendorNoteUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.VendorNoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewVendorNoteUsecase(mockRepo, timeout)

	vendorNoteID := bson.NewObjectID().Hex()

	updatedVendorNote := domain.VendorNote{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		VendorID:     bson.NewObjectID(),
		Note:         "This is an updated note for the vendor.",
		CreatedOnUtc: time.Now().AddDate(0, 0, -1), // Created 1 day ago
	}

	mockRepo.On("FetchByID", mock.Anything, vendorNoteID).Return(updatedVendorNote, nil)

	result, err := usecase.FetchByID(context.Background(), vendorNoteID)

	assert.NoError(t, err)
	assert.Equal(t, updatedVendorNote, result)
	mockRepo.AssertExpectations(t)
}

func TestVendorNoteUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.VendorNoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewVendorNoteUsecase(mockRepo, timeout)

	newVendorNote := &domain.VendorNote{
		VendorID:     bson.NewObjectID(),
		Note:         "This is a note for the vendor.",
		CreatedOnUtc: time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newVendorNote).Return(nil)

	err := usecase.Create(context.Background(), newVendorNote)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorNoteUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.VendorNoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewVendorNoteUsecase(mockRepo, timeout)

	updatedVendorNote := &domain.VendorNote{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		VendorID:     bson.NewObjectID(),
		Note:         "This is an updated note for the vendor.",
		CreatedOnUtc: time.Now().AddDate(0, 0, -1), // Created 1 day ago
	}

	mockRepo.On("Update", mock.Anything, updatedVendorNote).Return(nil)

	err := usecase.Update(context.Background(), updatedVendorNote)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorNoteUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.VendorNoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewVendorNoteUsecase(mockRepo, timeout)

	vendorNoteID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, vendorNoteID).Return(nil)

	err := usecase.Delete(context.Background(), vendorNoteID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorNoteUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.VendorNoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewVendorNoteUsecase(mockRepo, timeout)

	fetchedVendorNotes := []domain.VendorNote{
		{
			ID:           bson.NewObjectID(),
			VendorID:     bson.NewObjectID(),
			Note:         "This is a note for the vendor.",
			CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
		},
		{
			ID:           bson.NewObjectID(),
			VendorID:     bson.NewObjectID(),
			Note:         "This is another note for the vendor.",
			CreatedOnUtc: time.Now().AddDate(0, 0, -3), // Created 3 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedVendorNotes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedVendorNotes, result)
	mockRepo.AssertExpectations(t)
}
