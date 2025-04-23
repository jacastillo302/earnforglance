package usecase_test

import (
	"context"
	domian "earnforglance/server/domain/customers"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/customers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestCustomerPasswordUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CustomerPasswordRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerPasswordUsecase(mockRepo, timeout)

	customerID := bson.NewObjectID().Hex()

	updatedCustomerPassword := domian.CustomerPassword{
		CustomerID:       bson.NewObjectID(),  // Example Customer ID
		Password:         "hashedpassword123", // Example hashed password
		PasswordFormatID: 1,                   // Example password format ID (e.g., 1 for hashed)
		PasswordSalt:     "randomsaltvalue",   // Example password salt
		CreatedOnUTC:     time.Now().UTC(),
	}

	mockRepo.On("FetchByID", mock.Anything, customerID).Return(updatedCustomerPassword, nil)

	result, err := usecase.FetchByID(context.Background(), customerID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCustomerPassword, result)
	mockRepo.AssertExpectations(t)
}

func TestCustomerPasswordUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CustomerPasswordRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerPasswordUsecase(mockRepo, timeout)

	newCustomerPassword := &domian.CustomerPassword{
		ID:               bson.NewObjectID(),  // Generate a new MongoDB ObjectID
		CustomerID:       bson.NewObjectID(),  // Example Customer ID
		Password:         "hashedpassword123", // Example hashed password
		PasswordFormatID: 1,                   // Example password format ID (e.g., 1 for hashed)
		PasswordSalt:     "randomsaltvalue",   // Example password salt
		CreatedOnUTC:     time.Now().UTC(),
	}

	mockRepo.On("Create", mock.Anything, newCustomerPassword).Return(nil)

	err := usecase.Create(context.Background(), newCustomerPassword)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerPasswordUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CustomerPasswordRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerPasswordUsecase(mockRepo, timeout)

	updatedCustomerPassword := &domian.CustomerPassword{
		ID:               bson.NewObjectID(),  // Generate a new MongoDB ObjectID
		CustomerID:       bson.NewObjectID(),  // Example Customer ID
		Password:         "hashedpassword123", // Example hashed password
		PasswordFormatID: 1,                   // Example password format ID (e.g., 1 for hashed)
		PasswordSalt:     "randomsaltvalue",   // Example password salt
		CreatedOnUTC:     time.Now().UTC(),
	}

	mockRepo.On("Update", mock.Anything, updatedCustomerPassword).Return(nil)

	err := usecase.Update(context.Background(), updatedCustomerPassword)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerPasswordUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CustomerPasswordRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerPasswordUsecase(mockRepo, timeout)

	customerID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customerID).Return(nil)

	err := usecase.Delete(context.Background(), customerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerPasswordUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CustomerPasswordRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerPasswordUsecase(mockRepo, timeout)

	fetchedCustomerPasswords := []domian.CustomerPassword{
		{
			ID:               bson.NewObjectID(),
			CustomerID:       bson.NewObjectID(),
			Password:         "hashedpassword456",
			PasswordFormatID: 2, // Example password format ID (e.g., 2 for encrypted)
			PasswordSalt:     "anotherrandomsalt",
			CreatedOnUTC:     time.Now().UTC().Add(-24 * time.Hour), // Created 1 day ago
		},
		{
			ID:               bson.NewObjectID(),
			CustomerID:       bson.NewObjectID(),
			Password:         "hashedpassword789",
			PasswordFormatID: 1,
			PasswordSalt:     "yetanotherrandomsalt",
			CreatedOnUTC:     time.Now().UTC().Add(-48 * time.Hour), // Created 2 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCustomerPasswords, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCustomerPasswords, result)
	mockRepo.AssertExpectations(t)
}
