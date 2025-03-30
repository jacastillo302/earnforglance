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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCustomerSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CustomerSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerSettingsUsecase(mockRepo, timeout)

	customersID := primitive.NewObjectID().Hex()

	updatedCustomerSettings := domian.CustomerSettings{
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

	mockRepo.On("FetchByID", mock.Anything, customersID).Return(updatedCustomerSettings, nil)

	result, err := usecase.FetchByID(context.Background(), customersID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCustomerSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestCustomerSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CustomerSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerSettingsUsecase(mockRepo, timeout)

	newCustomerSettings := &domian.CustomerSettings{
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

	mockRepo.On("Create", mock.Anything, newCustomerSettings).Return(nil)

	err := usecase.Create(context.Background(), newCustomerSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CustomerSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerSettingsUsecase(mockRepo, timeout)

	updatedCustomerSettings := &domian.CustomerSettings{
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

	mockRepo.On("Update", mock.Anything, updatedCustomerSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedCustomerSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CustomerSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerSettingsUsecase(mockRepo, timeout)

	customersID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customersID).Return(nil)

	err := usecase.Delete(context.Background(), customersID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CustomerSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerSettingsUsecase(mockRepo, timeout)

	fetchedCustomerSettings := []domian.CustomerSettings{
		{
			ID:                                         primitive.NewObjectID(),
			UsernamesEnabled:                           false,
			CheckUsernameAvailabilityEnabled:           false,
			AllowUsersToChangeUsernames:                true,
			UsernameValidationEnabled:                  false,
			UsernameValidationUseRegex:                 false,
			UsernameValidationRule:                     "",
			PhoneNumberValidationEnabled:               false,
			PhoneNumberValidationUseRegex:              false,
			PhoneNumberValidationRule:                  "",
			DefaultPasswordFormat:                      "Clear",
			HashedPasswordFormat:                       "",
			PasswordMinLength:                          6,
			PasswordMaxLength:                          12,
			PasswordRequireLowercase:                   false,
			PasswordRequireUppercase:                   false,
			PasswordRequireNonAlphanumeric:             false,
			PasswordRequireDigit:                       false,
			UnduplicatedPasswordsNumber:                0,
			PasswordRecoveryLinkDaysValid:              3,
			PasswordLifetime:                           0,
			FailedPasswordAllowedAttempts:              10,
			FailedPasswordLockoutMinutes:               0,
			RequiredReLoginAfterPasswordChange:         false,
			UserRegistrationType:                       "EmailValidation",
			AllowCustomersToUploadAvatars:              false,
			AvatarMaximumSizeBytes:                     0,
			DefaultAvatarEnabled:                       false,
			ShowCustomersLocation:                      false,
			ShowCustomersJoinDate:                      false,
			AllowViewingProfiles:                       false,
			NotifyNewCustomerRegistration:              false,
			HideDownloadableProductsTab:                true,
			HideBackInStockSubscriptionsTab:            true,
			DownloadableProductsValidateUser:           false,
			CustomerNameFormat:                         "Username",
			NewsletterEnabled:                          false,
			NewsletterTickedByDefault:                  false,
			HideNewsletterBlock:                        true,
			NewsletterBlockAllowToUnsubscribe:          false,
			OnlineCustomerMinutes:                      0,
			StoreLastVisitedPage:                       false,
			StoreIpAddresses:                           false,
			LastActivityMinutes:                        0,
			SuffixDeletedCustomers:                     false,
			EnteringEmailTwice:                         true,
			RequireRegistrationForDownloadableProducts: false,
			AllowCustomersToCheckGiftCardBalance:       false,
			DeleteGuestTaskOlderThanMinutes:            0,
			FirstNameEnabled:                           false,
			FirstNameRequired:                          false,
			LastNameEnabled:                            false,
			LastNameRequired:                           false,
			GenderEnabled:                              false,
			NeutralGenderEnabled:                       false,
			DateOfBirthEnabled:                         false,
			DateOfBirthRequired:                        false,
			DateOfBirthMinimumAge:                      nil,
			CompanyEnabled:                             false,
			CompanyRequired:                            false,
			StreetAddressEnabled:                       false,
			StreetAddressRequired:                      false,
			StreetAddress2Enabled:                      false,
			StreetAddress2Required:                     false,
			ZipPostalCodeEnabled:                       false,
			ZipPostalCodeRequired:                      false,
			CityEnabled:                                false,
			CityRequired:                               false,
			CountyEnabled:                              false,
			CountyRequired:                             false,
			CountryEnabled:                             false,
			CountryRequired:                            false,
			DefaultCountryId:                           nil,
			StateProvinceEnabled:                       false,
			StateProvinceRequired:                      false,
			PhoneEnabled:                               false,
			PhoneRequired:                              false,
			FaxEnabled:                                 false,
			FaxRequired:                                false,
			AcceptPrivacyPolicyEnabled:                 false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCustomerSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCustomerSettings, result)
	mockRepo.AssertExpectations(t)
}
