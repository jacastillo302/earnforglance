package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/customers"
	repository "earnforglance/server/repository/customers"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultCustomer struct {
	mock.Mock
}

func (m *MockSingleResultCustomer) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Customer); ok {
		*v.(*domain.Customer) = *result
	}
	return args.Error(1)
}

var mockItemCustomer = &domain.Customer{
	ID:                                 bson.NewObjectID(), // Existing ID of the record to update
	CustomerGuid:                       uuid.New().String(),
	Username:                           "janedoe",
	Email:                              "janedoe@example.com",
	FirstName:                          "Jane",
	LastName:                           "Doe",
	Gender:                             "Female",
	DateOfBirth:                        new(time.Time),
	Company:                            "Updated Corp",
	StreetAddress:                      "456 Elm St",
	StreetAddress2:                     "Suite 101",
	ZipPostalCode:                      "67890",
	City:                               "Updated City",
	County:                             "Updated County",
	CountryID:                          bson.NewObjectID(),
	StateProvinceID:                    bson.NewObjectID(),
	Phone:                              "987-654-3210",
	Fax:                                "987-654-3211",
	VatNumber:                          "VAT654321",
	VatNumberStatusID:                  "Valid",
	TimeZoneID:                         "PST",
	CustomPermisionRecordAttributesXML: "<attributes><attribute>updated</attribute></attributes>",
	CurrencyID:                         new(bson.ObjectID),
	LanguageID:                         new(bson.ObjectID),
	TaxDisplayTypeID:                   new(int),
	EmailToRevalidate:                  "janedoe@newdomain.com",
	AdminComment:                       "Updated customer",
	IsTaxExempt:                        true,
	AffiliateID:                        "Affiliate123",
	VendorID:                           "Vendor456",
	HasShoppingCartItems:               true,
	RequireReLogin:                     true,
	FailedLoginAttempts:                3,
	CannotLoginUntilDateUtc:            new(time.Time),
	Active:                             false,
	Deleted:                            true,
	IsSystemAccount:                    true,
	SystemName:                         "SystemAccount",
	LastIpAddress:                      "192.168.1.2",
	CreatedOnUtc:                       time.Now().AddDate(0, 0, -30), // Created 30 days ago
	LastLoginDateUtc:                   new(time.Time),
	LastActivityDateUtc:                time.Now(),
	RegisteredInStoreID:                "64f1b3d4e5f6789012345678",
	BillingAddressID:                   new(bson.ObjectID),
	MustChangePassword:                 true,
	ShippingAddressID:                  new(bson.ObjectID),
}

func TestCustomerRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCustomer

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomer{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCustomer, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomer.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomer{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomer.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCustomerRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomer

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCustomer).Return(nil, nil).Once()

	repo := repository.NewCustomerRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCustomer)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCustomerRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomer

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCustomer.ID}
	update := bson.M{"$set": mockItemCustomer}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCustomerRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCustomer)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
