package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/common"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/common"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddressUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.AddressRepository)
	timeout := time.Duration(10)
	usecase := test.NewAddressUsecase(mockRepo, timeout)

	addressID := primitive.NewObjectID().Hex()

	updatedAddress := domain.Address{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		FirstName:        "Jane",
		LastName:         "Smith",
		Email:            "jane.smith@example.com",
		Company:          "Updated Corp",
		CountryID:        new(primitive.ObjectID),
		StateProvinceID:  new(primitive.ObjectID),
		County:           "Updated County",
		City:             "Updated City",
		Address1:         "456 Elm St",
		Address2:         "Suite 101",
		ZipPostalCode:    "67890",
		PhoneNumber:      "987-654-3210",
		FaxNumber:        "987-654-3211",
		CustomAttributes: "<custom><attribute>updated</attribute></custom>",
		CreatedOnUtc:     time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, addressID).Return(updatedAddress, nil)

	result, err := usecase.FetchByID(context.Background(), addressID)

	assert.NoError(t, err)
	assert.Equal(t, updatedAddress, result)
	mockRepo.AssertExpectations(t)
}

func TestAddressUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.AddressRepository)
	timeout := time.Duration(10)
	usecase := test.NewAddressUsecase(mockRepo, timeout)

	newAddress := &domain.Address{
		FirstName:        "John",
		LastName:         "Doe",
		Email:            "john.doe@example.com",
		Company:          "Example Corp",
		CountryID:        nil,
		StateProvinceID:  nil,
		County:           "Example County",
		City:             "Example City",
		Address1:         "123 Main St",
		Address2:         "Apt 4B",
		ZipPostalCode:    "12345",
		PhoneNumber:      "123-456-7890",
		FaxNumber:        "123-456-7891",
		CustomAttributes: "<custom><attribute>value</attribute></custom>",
		CreatedOnUtc:     time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newAddress).Return(nil)

	err := usecase.Create(context.Background(), newAddress)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddressUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.AddressRepository)
	timeout := time.Duration(10)
	usecase := test.NewAddressUsecase(mockRepo, timeout)

	updatedAddress := &domain.Address{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		FirstName:        "Jane",
		LastName:         "Smith",
		Email:            "jane.smith@example.com",
		Company:          "Updated Corp",
		CountryID:        new(primitive.ObjectID),
		StateProvinceID:  new(primitive.ObjectID),
		County:           "Updated County",
		City:             "Updated City",
		Address1:         "456 Elm St",
		Address2:         "Suite 101",
		ZipPostalCode:    "67890",
		PhoneNumber:      "987-654-3210",
		FaxNumber:        "987-654-3211",
		CustomAttributes: "<custom><attribute>updated</attribute></custom>",
		CreatedOnUtc:     time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}
	*updatedAddress.CountryID = primitive.NewObjectID()
	*updatedAddress.StateProvinceID = primitive.NewObjectID()

	mockRepo.On("Update", mock.Anything, updatedAddress).Return(nil)

	err := usecase.Update(context.Background(), updatedAddress)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddressUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.AddressRepository)
	timeout := time.Duration(10)
	usecase := test.NewAddressUsecase(mockRepo, timeout)

	addressID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, addressID).Return(nil)

	err := usecase.Delete(context.Background(), addressID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddressUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.AddressRepository)
	timeout := time.Duration(10)
	usecase := test.NewAddressUsecase(mockRepo, timeout)

	fetchedAddresses := []domain.Address{
		{
			ID:               primitive.NewObjectID(),
			FirstName:        "John",
			LastName:         "Doe",
			Email:            "john.doe@example.com",
			Company:          "Example Corp",
			CountryID:        nil,
			StateProvinceID:  nil,
			County:           "Example County",
			City:             "Example City",
			Address1:         "123 Main St",
			Address2:         "Apt 4B",
			ZipPostalCode:    "12345",
			PhoneNumber:      "123-456-7890",
			FaxNumber:        "123-456-7891",
			CustomAttributes: "<custom><attribute>value</attribute></custom>",
			CreatedOnUtc:     time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:               primitive.NewObjectID(),
			FirstName:        "Jane",
			LastName:         "Smith",
			Email:            "jane.smith@example.com",
			Company:          "Updated Corp",
			CountryID:        new(primitive.ObjectID),
			StateProvinceID:  new(primitive.ObjectID),
			County:           "Updated County",
			City:             "Updated City",
			Address1:         "456 Elm St",
			Address2:         "Suite 101",
			ZipPostalCode:    "67890",
			PhoneNumber:      "987-654-3210",
			FaxNumber:        "987-654-3211",
			CustomAttributes: "<custom><attribute>updated</attribute></custom>",
			CreatedOnUtc:     time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}
	*fetchedAddresses[1].CountryID = primitive.NewObjectID()
	*fetchedAddresses[1].StateProvinceID = primitive.NewObjectID()

	mockRepo.On("Fetch", mock.Anything).Return(fetchedAddresses, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedAddresses, result)
	mockRepo.AssertExpectations(t)
}
