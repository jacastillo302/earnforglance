package domain

// NopCustomerDefaults represents default values related to customers data
var CustomerDefaults = struct {
	// System customer roles
	AdministratorsRoleName  string
	ForumModeratorsRoleName string
	RegisteredRoleName      string
	GuestsRoleName          string
	VendorsRoleName         string

	// System customers
	SearchEngineCustomerName   string
	BackgroundTaskCustomerName string

	// Customer attributes
	DiscountCouponCodeAttribute                        string
	GiftCardCouponCodesAttribute                       string
	AvatarPictureIdAttribute                           string
	ForumPostCountAttribute                            string
	SignatureAttribute                                 string
	PasswordRecoveryTokenAttribute                     string
	PasswordRecoveryTokenDateGeneratedAttribute        string
	AccountActivationTokenAttribute                    string
	EmailRevalidationTokenAttribute                    string
	LastVisitedPageAttribute                           string
	ImpersonatedCustomerIdAttribute                    string
	AdminAreaStoreScopeConfigurationAttribute          string
	SelectedPaymentMethodAttribute                     string
	SelectedShippingOptionAttribute                    string
	SelectedPickupPointAttribute                       string
	CheckoutAttributes                                 string
	OfferedShippingOptionsAttribute                    string
	LastContinueShoppingPageAttribute                  string
	NotifiedAboutNewPrivateMessagesAttribute           string
	WorkingThemeNameAttribute                          string
	UseRewardPointsDuringCheckoutAttribute             string
	EuCookieLawAcceptedAttribute                       string
	SelectedMultiFactorAuthenticationProviderAttribute string
	CustomerMultiFactorAuthenticationInfo              string
	HideConfigurationStepsAttribute                    string
	CloseConfigurationStepsAttribute                   string
}{
	// System customer roles
	AdministratorsRoleName:  "Administrators",
	ForumModeratorsRoleName: "ForumModerators",
	RegisteredRoleName:      "Registered",
	GuestsRoleName:          "Guests",
	VendorsRoleName:         "Vendors",

	// System customers
	SearchEngineCustomerName:   "SearchEngine",
	BackgroundTaskCustomerName: "BackgroundTask",

	// Customer attributes
	DiscountCouponCodeAttribute:                        "DiscountCouponCode",
	GiftCardCouponCodesAttribute:                       "GiftCardCouponCodes",
	AvatarPictureIdAttribute:                           "AvatarPictureId",
	ForumPostCountAttribute:                            "ForumPostCount",
	SignatureAttribute:                                 "Signature",
	PasswordRecoveryTokenAttribute:                     "PasswordRecoveryToken",
	PasswordRecoveryTokenDateGeneratedAttribute:        "PasswordRecoveryTokenDateGenerated",
	AccountActivationTokenAttribute:                    "AccountActivationToken",
	EmailRevalidationTokenAttribute:                    "EmailRevalidationToken",
	LastVisitedPageAttribute:                           "LastVisitedPage",
	ImpersonatedCustomerIdAttribute:                    "ImpersonatedCustomerId",
	AdminAreaStoreScopeConfigurationAttribute:          "AdminAreaStoreScopeConfiguration",
	SelectedPaymentMethodAttribute:                     "SelectedPaymentMethod",
	SelectedShippingOptionAttribute:                    "SelectedShippingOption",
	SelectedPickupPointAttribute:                       "SelectedPickupPoint",
	CheckoutAttributes:                                 "CheckoutAttributes",
	OfferedShippingOptionsAttribute:                    "OfferedShippingOptions",
	LastContinueShoppingPageAttribute:                  "LastContinueShoppingPage",
	NotifiedAboutNewPrivateMessagesAttribute:           "NotifiedAboutNewPrivateMessages",
	WorkingThemeNameAttribute:                          "WorkingThemeName",
	UseRewardPointsDuringCheckoutAttribute:             "UseRewardPointsDuringCheckout",
	EuCookieLawAcceptedAttribute:                       "EuCookieLaw.Accepted",
	SelectedMultiFactorAuthenticationProviderAttribute: "SelectedMultiFactorAuthProvider",
	CustomerMultiFactorAuthenticationInfo:              "CustomerMultiFactorAuthenticationInfo",
	HideConfigurationStepsAttribute:                    "HideConfigurationSteps",
	CloseConfigurationStepsAttribute:                   "CloseConfigurationSteps",
}
