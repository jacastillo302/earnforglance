package usecase

import (
	"context"
	domain "earnforglance/server/domain/customers"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCustomerUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CustomerRepository)
	timeout := time.Duration(10)
	usecase := NewCustomerUsecase(mockRepo, timeout)

	customerID := primitive.NewObjectID().Hex()

	updatedCustomer := domain.Customer{
		ID:                          primitive.NewObjectID(), // Existing ID of the record to update
		CustomerGuid:                uuid.New(),
		Username:                    "janedoe",
		Email:                       "janedoe@example.com",
		FirstName:                   "Jane",
		LastName:                    "Doe",
		Gender:                      "Female",
		DateOfBirth:                 new(time.Time),
		Company:                     "Updated Corp",
		StreetAddress:               "456 Elm St",
		StreetAddress2:              "Suite 101",
		ZipPostalCode:               "67890",
		City:                        "Updated City",
		County:                      "Updated County",
		CountryID:                   primitive.NewObjectID(),
		StateProvinceID:             primitive.NewObjectID(),
		Phone:                       "987-654-3210",
		Fax:                         "987-654-3211",
		VatNumber:                   "VAT654321",
		VatNumberStatusID:           primitive.NewObjectID(),
		TimeZoneID:                  "PST",
		CustomCustomerAttributesXML: "<attributes><attribute>updated</attribute></attributes>",
		CurrencyID:                  new(primitive.ObjectID),
		LanguageID:                  new(primitive.ObjectID),
		TaxDisplayTypeID:            new(primitive.ObjectID),
		EmailToRevalidate:           "janedoe@newdomain.com",
		AdminComment:                "Updated customer",
		IsTaxExempt:                 true,
		AffiliateID:                 primitive.NewObjectID(),
		VendorID:                    primitive.NewObjectID(),
		HasShoppingCartItems:        true,
		RequireReLogin:              true,
		FailedLoginAttempts:         3,
		CannotLoginUntilDateUtc:     new(time.Time),
		Active:                      false,
		Deleted:                     true,
		IsSystemAccount:             true,
		SystemName:                  "SystemAccount",
		LastIpAddress:               "192.168.1.2",
		CreatedOnUtc:                time.Now().AddDate(0, 0, -30), // Created 30 days ago
		LastLoginDateUtc:            new(time.Time),
		LastActivityDateUtc:         time.Now(),
		RegisteredInStoreID:         2,
		BillingAddressID:            new(primitive.ObjectID),
		MustChangePassword:          true,
		ShippingAddressID:           new(primitive.ObjectID),
		VatNumberStatus:             2,
		TaxDisplayType:              new(int),
	}

	mockRepo.On("FetchByID", mock.Anything, customerID).Return(updatedCustomer, nil)

	result, err := usecase.FetchByID(context.Background(), customerID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCustomer, result)
	mockRepo.AssertExpectations(t)
}

func TestCustomerUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CustomerRepository)
	timeout := time.Duration(10)
	usecase := NewCustomerUsecase(mockRepo, timeout)

	newCustomer := &domain.Customer{
		CustomerGuid:                uuid.New(),
		Username:                    "johndoe",
		Email:                       "johndoe@example.com",
		FirstName:                   "John",
		LastName:                    "Doe",
		Gender:                      "Male",
		DateOfBirth:                 nil,
		Company:                     "Example Corp",
		StreetAddress:               "123 Main St",
		StreetAddress2:              "Apt 4B",
		ZipPostalCode:               "12345",
		City:                        "Example City",
		County:                      "Example County",
		CountryID:                   primitive.NewObjectID(),
		StateProvinceID:             primitive.NewObjectID(),
		Phone:                       "123-456-7890",
		Fax:                         "123-456-7891",
		VatNumber:                   "VAT123456",
		VatNumberStatusID:           primitive.NewObjectID(),
		TimeZoneID:                  "UTC",
		CustomCustomerAttributesXML: "<attributes><attribute>value</attribute></attributes>",
		CurrencyID:                  nil,
		LanguageID:                  nil,
		TaxDisplayTypeID:            nil,
		EmailToRevalidate:           "",
		AdminComment:                "New customer",
		IsTaxExempt:                 false,
		AffiliateID:                 primitive.NewObjectID(),
		VendorID:                    primitive.NewObjectID(),
		HasShoppingCartItems:        false,
		RequireReLogin:              false,
		FailedLoginAttempts:         0,
		CannotLoginUntilDateUtc:     nil,
		Active:                      true,
		Deleted:                     false,
		IsSystemAccount:             false,
		SystemName:                  "",
		LastIpAddress:               "192.168.1.1",
		CreatedOnUtc:                time.Now(),
		LastLoginDateUtc:            nil,
		LastActivityDateUtc:         time.Now(),
		RegisteredInStoreID:         1,
		BillingAddressID:            nil,
		MustChangePassword:          false,
		ShippingAddressID:           nil,
		VatNumberStatus:             1,
		TaxDisplayType:              nil,
	}

	mockRepo.On("Create", mock.Anything, newCustomer).Return(nil)

	err := usecase.Create(context.Background(), newCustomer)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CustomerRepository)
	timeout := time.Duration(10)
	usecase := NewCustomerUsecase(mockRepo, timeout)

	updatedCustomer := &domain.Customer{
		ID:                          primitive.NewObjectID(), // Existing ID of the record to update
		CustomerGuid:                uuid.New(),
		Username:                    "janedoe",
		Email:                       "janedoe@example.com",
		FirstName:                   "Jane",
		LastName:                    "Doe",
		Gender:                      "Female",
		DateOfBirth:                 new(time.Time),
		Company:                     "Updated Corp",
		StreetAddress:               "456 Elm St",
		StreetAddress2:              "Suite 101",
		ZipPostalCode:               "67890",
		City:                        "Updated City",
		County:                      "Updated County",
		CountryID:                   primitive.NewObjectID(),
		StateProvinceID:             primitive.NewObjectID(),
		Phone:                       "987-654-3210",
		Fax:                         "987-654-3211",
		VatNumber:                   "VAT654321",
		VatNumberStatusID:           primitive.NewObjectID(),
		TimeZoneID:                  "PST",
		CustomCustomerAttributesXML: "<attributes><attribute>updated</attribute></attributes>",
		CurrencyID:                  new(primitive.ObjectID),
		LanguageID:                  new(primitive.ObjectID),
		TaxDisplayTypeID:            new(primitive.ObjectID),
		EmailToRevalidate:           "janedoe@newdomain.com",
		AdminComment:                "Updated customer",
		IsTaxExempt:                 true,
		AffiliateID:                 primitive.NewObjectID(),
		VendorID:                    primitive.NewObjectID(),
		HasShoppingCartItems:        true,
		RequireReLogin:              true,
		FailedLoginAttempts:         3,
		CannotLoginUntilDateUtc:     new(time.Time),
		Active:                      false,
		Deleted:                     true,
		IsSystemAccount:             true,
		SystemName:                  "SystemAccount",
		LastIpAddress:               "192.168.1.2",
		CreatedOnUtc:                time.Now().AddDate(0, 0, -30), // Created 30 days ago
		LastLoginDateUtc:            new(time.Time),
		LastActivityDateUtc:         time.Now(),
		RegisteredInStoreID:         2,
		BillingAddressID:            new(primitive.ObjectID),
		MustChangePassword:          true,
		ShippingAddressID:           new(primitive.ObjectID),
		VatNumberStatus:             2,
		TaxDisplayType:              new(int),
	}
	*updatedCustomer.DateOfBirth = time.Now().AddDate(-30, 0, 0)           // 30 years ago
	*updatedCustomer.CannotLoginUntilDateUtc = time.Now().AddDate(0, 0, 7) // 7 days from now
	*updatedCustomer.LastLoginDateUtc = time.Now().AddDate(0, 0, -1)       // 1 day ago
	*updatedCustomer.CurrencyID = primitive.NewObjectID()
	*updatedCustomer.LanguageID = primitive.NewObjectID()
	*updatedCustomer.TaxDisplayTypeID = primitive.NewObjectID()
	*updatedCustomer.BillingAddressID = primitive.NewObjectID()
	*updatedCustomer.ShippingAddressID = primitive.NewObjectID()
	*updatedCustomer.TaxDisplayType = 1

	mockRepo.On("Update", mock.Anything, updatedCustomer).Return(nil)

	err := usecase.Update(context.Background(), updatedCustomer)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CustomerRepository)
	timeout := time.Duration(10)
	usecase := NewCustomerUsecase(mockRepo, timeout)

	customerID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customerID).Return(nil)

	err := usecase.Delete(context.Background(), customerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CustomerRepository)
	timeout := time.Duration(10)
	usecase := NewCustomerUsecase(mockRepo, timeout)

	fetchedCustomers := []domain.Customer{
		{
			ID:                          primitive.NewObjectID(),
			CustomerGuid:                uuid.New(),
			Username:                    "johndoe",
			Email:                       "johndoe@example.com",
			FirstName:                   "John",
			LastName:                    "Doe",
			Gender:                      "Male",
			DateOfBirth:                 nil,
			Company:                     "Example Corp",
			StreetAddress:               "123 Main St",
			StreetAddress2:              "Apt 4B",
			ZipPostalCode:               "12345",
			City:                        "Example City",
			County:                      "Example County",
			CountryID:                   primitive.NewObjectID(),
			StateProvinceID:             primitive.NewObjectID(),
			Phone:                       "123-456-7890",
			Fax:                         "123-456-7891",
			VatNumber:                   "VAT123456",
			VatNumberStatusID:           primitive.NewObjectID(),
			TimeZoneID:                  "UTC",
			CustomCustomerAttributesXML: "<attributes><attribute>value</attribute></attributes>",
			CurrencyID:                  nil,
			LanguageID:                  nil,
			TaxDisplayTypeID:            nil,
			EmailToRevalidate:           "",
			AdminComment:                "New customer",
			IsTaxExempt:                 false,
			AffiliateID:                 primitive.NewObjectID(),
			VendorID:                    primitive.NewObjectID(),
			HasShoppingCartItems:        false,
			RequireReLogin:              false,
			FailedLoginAttempts:         0,
			CannotLoginUntilDateUtc:     nil,
			Active:                      true,
			Deleted:                     false,
			IsSystemAccount:             false,
			SystemName:                  "",
			LastIpAddress:               "192.168.1.1",
			CreatedOnUtc:                time.Now().AddDate(0, 0, -10), // Created 10 days ago
			LastLoginDateUtc:            nil,
			LastActivityDateUtc:         time.Now(),
			RegisteredInStoreID:         1,
			BillingAddressID:            nil,
			MustChangePassword:          false,
			ShippingAddressID:           nil,
			VatNumberStatus:             1,
			TaxDisplayType:              nil,
		},
		{
			ID:                          primitive.NewObjectID(),
			CustomerGuid:                uuid.New(),
			Username:                    "janedoe",
			Email:                       "janedoe@example.com",
			FirstName:                   "Jane",
			LastName:                    "Doe",
			Gender:                      "Female",
			DateOfBirth:                 new(time.Time),
			Company:                     "Updated Corp",
			StreetAddress:               "456 Elm St",
			StreetAddress2:              "Suite 101",
			ZipPostalCode:               "67890",
			City:                        "Updated City",
			County:                      "Updated County",
			CountryID:                   primitive.NewObjectID(),
			StateProvinceID:             primitive.NewObjectID(),
			Phone:                       "987-654-3210",
			Fax:                         "987-654-3211",
			VatNumber:                   "VAT654321",
			VatNumberStatusID:           primitive.NewObjectID(),
			TimeZoneID:                  "PST",
			CustomCustomerAttributesXML: "<attributes><attribute>updated</attribute></attributes>",
			CurrencyID:                  new(primitive.ObjectID),
			LanguageID:                  new(primitive.ObjectID),
			TaxDisplayTypeID:            new(primitive.ObjectID),
			EmailToRevalidate:           "janedoe@newdomain.com",
			AdminComment:                "Updated customer",
			IsTaxExempt:                 true,
			AffiliateID:                 primitive.NewObjectID(),
			VendorID:                    primitive.NewObjectID(),
			HasShoppingCartItems:        true,
			RequireReLogin:              true,
			FailedLoginAttempts:         3,
			CannotLoginUntilDateUtc:     new(time.Time),
			Active:                      false,
			Deleted:                     true,
			IsSystemAccount:             true,
			SystemName:                  "SystemAccount",
			LastIpAddress:               "192.168.1.2",
			CreatedOnUtc:                time.Now().AddDate(0, 0, -30), // Created 30 days ago
			LastLoginDateUtc:            new(time.Time),
			LastActivityDateUtc:         time.Now(),
			RegisteredInStoreID:         2,
			BillingAddressID:            new(primitive.ObjectID),
			MustChangePassword:          true,
			ShippingAddressID:           new(primitive.ObjectID),
			VatNumberStatus:             2,
			TaxDisplayType:              new(int),
		},
	}
	*fetchedCustomers[1].DateOfBirth = time.Now().AddDate(-30, 0, 0)           // 30 years ago
	*fetchedCustomers[1].CannotLoginUntilDateUtc = time.Now().AddDate(0, 0, 7) // 7 days from now
	*fetchedCustomers[1].LastLoginDateUtc = time.Now().AddDate(0, 0, -1)       // 1 day ago
	*fetchedCustomers[1].CurrencyID = primitive.NewObjectID()
	*fetchedCustomers[1].LanguageID = primitive.NewObjectID()
	*fetchedCustomers[1].TaxDisplayTypeID = primitive.NewObjectID()
	*fetchedCustomers[1].BillingAddressID = primitive.NewObjectID()
	*fetchedCustomers[1].ShippingAddressID = primitive.NewObjectID()
	*fetchedCustomers[1].TaxDisplayType = 1

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCustomers, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCustomers, result)
	mockRepo.AssertExpectations(t)
}
