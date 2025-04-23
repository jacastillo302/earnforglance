package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionCustomerSettings = "customer_settings"
)

// CustomerSettings represents customer settings in the system.
type CustomerSettings struct {
	ID                                         bson.ObjectID `bson:"_id,omitempty"`
	UsernamesEnabled                           bool          `bson:"usernames_enabled"`
	CheckUsernameAvailabilityEnabled           bool          `bson:"check_username_availability_enabled"`
	AllowUsersToChangeUsernames                bool          `bson:"allow_users_to_change_usernames"`
	UsernameValidationEnabled                  bool          `bson:"username_validation_enabled"`
	UsernameValidationUseRegex                 bool          `bson:"username_validation_use_regex"`
	UsernameValidationRule                     string        `bson:"username_validation_rule"`
	PhoneNumberValidationEnabled               bool          `bson:"phone_number_validation_enabled"`
	PhoneNumberValidationUseRegex              bool          `bson:"phone_number_validation_use_regex"`
	PhoneNumberValidationRule                  string        `bson:"phone_number_validation_rule"`
	DefaultPasswordFormat                      string        `bson:"default_password_format"`
	HashedPasswordFormat                       string        `bson:"hashed_password_format"`
	PasswordMinLength                          int           `bson:"password_min_length"`
	PasswordMaxLength                          int           `bson:"password_max_length"`
	PasswordRequireLowercase                   bool          `bson:"password_require_lowercase"`
	PasswordRequireUppercase                   bool          `bson:"password_require_uppercase"`
	PasswordRequireNonAlphanumeric             bool          `bson:"password_require_non_alphanumeric"`
	PasswordRequireDigit                       bool          `bson:"password_require_digit"`
	UnduplicatedPasswordsNumber                int           `bson:"unduplicated_passwords_number"`
	PasswordRecoveryLinkDaysValid              int           `bson:"password_recovery_link_days_valid"`
	PasswordLifetime                           int           `bson:"password_lifetime"`
	FailedPasswordAllowedAttempts              int           `bson:"failed_password_allowed_attempts"`
	FailedPasswordLockoutMinutes               int           `bson:"failed_password_lockout_minutes"`
	RequiredReLoginAfterPasswordChange         bool          `bson:"required_relogin_after_password_change"`
	UserRegistrationType                       string        `bson:"user_registration_type"`
	AllowCustomersToUploadAvatars              bool          `bson:"allow_customers_to_upload_avatars"`
	AvatarMaximumSizeBytes                     int           `bson:"avatar_maximum_size_bytes"`
	DefaultAvatarEnabled                       bool          `bson:"default_avatar_enabled"`
	ShowCustomersLocation                      bool          `bson:"show_customers_location"`
	ShowCustomersJoinDate                      bool          `bson:"show_customers_join_date"`
	AllowViewingProfiles                       bool          `bson:"allow_viewing_profiles"`
	NotifyNewCustomerRegistration              bool          `bson:"notify_new_customer_registration"`
	HideDownloadableProductsTab                bool          `bson:"hide_downloadable_products_tab"`
	HideBackInStockSubscriptionsTab            bool          `bson:"hide_back_in_stock_subscriptions_tab"`
	DownloadableProductsValidateUser           bool          `bson:"downloadable_products_validate_user"`
	CustomerNameFormat                         string        `bson:"customer_name_format"`
	NewsletterEnabled                          bool          `bson:"newsletter_enabled"`
	NewsletterTickedByDefault                  bool          `bson:"newsletter_ticked_by_default"`
	HideNewsletterBlock                        bool          `bson:"hide_newsletter_block"`
	NewsletterBlockAllowToUnsubscribe          bool          `bson:"newsletter_block_allow_to_unsubscribe"`
	OnlineCustomerMinutes                      int           `bson:"online_customer_minutes"`
	StoreLastVisitedPage                       bool          `bson:"store_last_visited_page"`
	StoreIpAddresses                           bool          `bson:"store_ip_addresses"`
	LastActivityMinutes                        int           `bson:"last_activity_minutes"`
	SuffixDeletedCustomers                     bool          `bson:"suffix_deleted_customers"`
	EnteringEmailTwice                         bool          `bson:"entering_email_twice"`
	RequireRegistrationForDownloadableProducts bool          `bson:"require_registration_for_downloadable_products"`
	AllowCustomersToCheckGiftCardBalance       bool          `bson:"allow_customers_to_check_gift_card_balance"`
	DeleteGuestTaskOlderThanMinutes            int           `bson:"delete_guest_task_older_than_minutes"`

	// Form fields
	FirstNameEnabled           bool `bson:"first_name_enabled"`
	FirstNameRequired          bool `bson:"first_name_required"`
	LastNameEnabled            bool `bson:"last_name_enabled"`
	LastNameRequired           bool `bson:"last_name_required"`
	GenderEnabled              bool `bson:"gender_enabled"`
	NeutralGenderEnabled       bool `bson:"neutral_gender_enabled"`
	DateOfBirthEnabled         bool `bson:"date_of_birth_enabled"`
	DateOfBirthRequired        bool `bson:"date_of_birth_required"`
	DateOfBirthMinimumAge      *int `bson:"date_of_birth_minimum_age"`
	CompanyEnabled             bool `bson:"company_enabled"`
	CompanyRequired            bool `bson:"company_required"`
	StreetAddressEnabled       bool `bson:"street_address_enabled"`
	StreetAddressRequired      bool `bson:"street_address_required"`
	StreetAddress2Enabled      bool `bson:"street_address2_enabled"`
	StreetAddress2Required     bool `bson:"street_address2_required"`
	ZipPostalCodeEnabled       bool `bson:"zip_postal_code_enabled"`
	ZipPostalCodeRequired      bool `bson:"zip_postal_code_required"`
	CityEnabled                bool `bson:"city_enabled"`
	CityRequired               bool `bson:"city_required"`
	CountyEnabled              bool `bson:"county_enabled"`
	CountyRequired             bool `bson:"county_required"`
	CountryEnabled             bool `bson:"country_enabled"`
	CountryRequired            bool `bson:"country_required"`
	DefaultCountryId           *int `bson:"default_country_id"`
	StateProvinceEnabled       bool `bson:"state_province_enabled"`
	StateProvinceRequired      bool `bson:"state_province_required"`
	PhoneEnabled               bool `bson:"phone_enabled"`
	PhoneRequired              bool `bson:"phone_required"`
	FaxEnabled                 bool `bson:"fax_enabled"`
	FaxRequired                bool `bson:"fax_required"`
	AcceptPrivacyPolicyEnabled bool `bson:"accept_privacy_policy_enabled"`
}

type CustomerSettingsRepository interface {
	CreateMany(c context.Context, items []CustomerSettings) error
	Create(c context.Context, permission_record_customer_role_mapping *CustomerSettings) error
	Update(c context.Context, permission_record_customer_role_mapping *CustomerSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerSettings, error)
	FetchByID(c context.Context, ID string) (CustomerSettings, error)
}

type CustomerSettingsUsecase interface {
	CreateMany(c context.Context, items []CustomerSettings) error
	FetchByID(c context.Context, ID string) (CustomerSettings, error)
	Create(c context.Context, permission_record_customer_role_mapping *CustomerSettings) error
	Update(c context.Context, permission_record_customer_role_mapping *CustomerSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerSettings, error)
}
