package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/customers"
	repository "earnforglance/server/repository/customers"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultCustomerSettings struct {
	mock.Mock
}

func (m *MockSingleResultCustomerSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CustomerSettings); ok {
		*v.(*domain.CustomerSettings) = *result
	}
	return args.Error(1)
}

var mockItemCustomerSettings = &domain.CustomerSettings{
	ID:                                         primitive.NewObjectID(), // Generate a new MongoDB ObjectID
	UsernamesEnabled:                           true,
	CheckUsernameAvailabilityEnabled:           true,
	AllowUsersToChangeUsernames:                false,
	UsernameValidationEnabled:                  true,
	UsernameValidationUseRegex:                 true,
	UsernameValidationRule:                     "^[a-zA-Z0-9]+$",
	PhoneNumberValidationEnabled:               true,
	PhoneNumberValidationUseRegex:              true,
	PhoneNumberValidationRule:                  "^[0-9]{10}$",
	DefaultPasswordFormat:                      "Hashed",
	HashedPasswordFormat:                       "SHA256",
	PasswordMinLength:                          8,
	PasswordMaxLength:                          20,
	PasswordRequireLowercase:                   true,
	PasswordRequireUppercase:                   true,
	PasswordRequireNonAlphanumeric:             true,
	PasswordRequireDigit:                       true,
	UnduplicatedPasswordsNumber:                5,
	PasswordRecoveryLinkDaysValid:              7,
	PasswordLifetime:                           90,
	FailedPasswordAllowedAttempts:              5,
	FailedPasswordLockoutMinutes:               15,
	RequiredReLoginAfterPasswordChange:         true,
	UserRegistrationType:                       "Standard",
	AllowCustomersToUploadAvatars:              true,
	AvatarMaximumSizeBytes:                     1048576,
	DefaultAvatarEnabled:                       true,
	ShowCustomersLocation:                      true,
	ShowCustomersJoinDate:                      true,
	AllowViewingProfiles:                       true,
	NotifyNewCustomerRegistration:              true,
	HideDownloadableProductsTab:                false,
	HideBackInStockSubscriptionsTab:            false,
	DownloadableProductsValidateUser:           true,
	CustomerNameFormat:                         "FullName",
	NewsletterEnabled:                          true,
	NewsletterTickedByDefault:                  false,
	HideNewsletterBlock:                        false,
	NewsletterBlockAllowToUnsubscribe:          true,
	OnlineCustomerMinutes:                      30,
	StoreLastVisitedPage:                       true,
	StoreIpAddresses:                           true,
	LastActivityMinutes:                        15,
	SuffixDeletedCustomers:                     true,
	EnteringEmailTwice:                         false,
	RequireRegistrationForDownloadableProducts: true,
	AllowCustomersToCheckGiftCardBalance:       true,
	DeleteGuestTaskOlderThanMinutes:            1440,
	FirstNameEnabled:                           true,
	FirstNameRequired:                          true,
	LastNameEnabled:                            true,
	LastNameRequired:                           true,
	GenderEnabled:                              true,
	NeutralGenderEnabled:                       false,
	DateOfBirthEnabled:                         true,
	DateOfBirthRequired:                        false,
	DateOfBirthMinimumAge:                      nil,
	CompanyEnabled:                             true,
	CompanyRequired:                            false,
	StreetAddressEnabled:                       true,
	StreetAddressRequired:                      true,
	StreetAddress2Enabled:                      false,
	StreetAddress2Required:                     false,
	ZipPostalCodeEnabled:                       true,
	ZipPostalCodeRequired:                      true,
	CityEnabled:                                true,
	CityRequired:                               true,
	CountyEnabled:                              false,
	CountyRequired:                             false,
	CountryEnabled:                             true,
	CountryRequired:                            true,
	DefaultCountryId:                           nil,
	StateProvinceEnabled:                       true,
	StateProvinceRequired:                      true,
	PhoneEnabled:                               true,
	PhoneRequired:                              true,
	FaxEnabled:                                 false,
	FaxRequired:                                false,
	AcceptPrivacyPolicyEnabled:                 true,
}

func TestCustomerSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCustomerSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCustomerSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCustomerSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCustomerSettings).Return(nil, nil).Once()

	repo := repository.NewCustomerSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCustomerSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCustomerSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCustomerSettings.ID}
	update := bson.M{"$set": mockItemCustomerSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCustomerSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCustomerSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
