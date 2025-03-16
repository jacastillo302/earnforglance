package domain

// MessageTemplateSystemNames represents message template system names
var MessageTemplateSystemNames = struct {
	// Customer
	CustomerRegisteredStoreOwnerNotification    string
	CustomerWelcomeMessage                      string
	CustomerEmailValidationMessage              string
	CustomerEmailRevalidationMessage            string
	CustomerPasswordRecoveryMessage             string
	DeleteCustomerRequestStoreOwnerNotification string

	// Order
	OrderPlacedVendorNotification                   string
	OrderPlacedStoreOwnerNotification               string
	OrderPlacedAffiliateNotification                string
	OrderPaidStoreOwnerNotification                 string
	OrderPaidCustomerNotification                   string
	OrderPaidVendorNotification                     string
	OrderPaidAffiliateNotification                  string
	OrderPlacedCustomerNotification                 string
	ShipmentSentCustomerNotification                string
	ShipmentReadyForPickupCustomerNotification      string
	ShipmentDeliveredCustomerNotification           string
	OrderProcessingCustomerNotification             string
	OrderCompletedCustomerNotification              string
	OrderCompletedStoreOwnerNotification            string
	OrderCancelledCustomerNotification              string
	OrderCancelledVendorNotification                string
	OrderRefundedStoreOwnerNotification             string
	OrderRefundedCustomerNotification               string
	NewOrderNoteAddedCustomerNotification           string
	RecurringPaymentCancelledStoreOwnerNotification string
	RecurringPaymentCancelledCustomerNotification   string
	RecurringPaymentFailedCustomerNotification      string

	// Newsletter
	NewsletterSubscriptionActivationMessage   string
	NewsletterSubscriptionDeactivationMessage string

	// To friend
	EmailAFriendMessage     string
	WishlistToFriendMessage string

	// Return requests
	NewReturnRequestStoreOwnerNotification         string
	NewReturnRequestCustomerNotification           string
	ReturnRequestStatusChangedCustomerNotification string

	// Forum
	NewForumTopicMessage       string
	NewForumPostMessage        string
	PrivateMessageNotification string

	// Misc
	NewVendorAccountApplyStoreOwnerNotification             string
	VendorInformationChangeStoreOwnerNotification           string
	GiftCardNotification                                    string
	ProductReviewStoreOwnerNotification                     string
	ProductReviewReplyCustomerNotification                  string
	QuantityBelowStoreOwnerNotification                     string
	QuantityBelowAttributeCombinationStoreOwnerNotification string
	QuantityBelowVendorNotification                         string
	QuantityBelowAttributeCombinationVendorNotification     string
	NewVATSubmittedStoreOwnerNotification                   string
	BlogCommentStoreOwnerNotification                       string
	NewsCommentStoreOwnerNotification                       string
	BackInStockNotification                                 string
	ContactUsMessage                                        string
	ContactVendorMessage                                    string
}{
	// Customer
	CustomerRegisteredStoreOwnerNotification:    "NewCustomer.Notification",
	CustomerWelcomeMessage:                      "Customer.WelcomeMessage",
	CustomerEmailValidationMessage:              "Customer.EmailValidationMessage",
	CustomerEmailRevalidationMessage:            "Customer.EmailRevalidationMessage",
	CustomerPasswordRecoveryMessage:             "Customer.PasswordRecovery",
	DeleteCustomerRequestStoreOwnerNotification: "Customer.Gdpr.DeleteRequest",

	// Order
	OrderPlacedVendorNotification:                   "OrderPlaced.VendorNotification",
	OrderPlacedStoreOwnerNotification:               "OrderPlaced.StoreOwnerNotification",
	OrderPlacedAffiliateNotification:                "OrderPlaced.AffiliateNotification",
	OrderPaidStoreOwnerNotification:                 "OrderPaid.StoreOwnerNotification",
	OrderPaidCustomerNotification:                   "OrderPaid.CustomerNotification",
	OrderPaidVendorNotification:                     "OrderPaid.VendorNotification",
	OrderPaidAffiliateNotification:                  "OrderPaid.AffiliateNotification",
	OrderPlacedCustomerNotification:                 "OrderPlaced.CustomerNotification",
	ShipmentSentCustomerNotification:                "ShipmentSent.CustomerNotification",
	ShipmentReadyForPickupCustomerNotification:      "ShipmentReadyForPickup.CustomerNotification",
	ShipmentDeliveredCustomerNotification:           "ShipmentDelivered.CustomerNotification",
	OrderProcessingCustomerNotification:             "OrderProcessing.CustomerNotification",
	OrderCompletedCustomerNotification:              "OrderCompleted.CustomerNotification",
	OrderCompletedStoreOwnerNotification:            "OrderCompleted.StoreOwnerNotification",
	OrderCancelledCustomerNotification:              "OrderCancelled.CustomerNotification",
	OrderCancelledVendorNotification:                "OrderCancelled.VendorNotification",
	OrderRefundedStoreOwnerNotification:             "OrderRefunded.StoreOwnerNotification",
	OrderRefundedCustomerNotification:               "OrderRefunded.CustomerNotification",
	NewOrderNoteAddedCustomerNotification:           "Customer.NewOrderNote",
	RecurringPaymentCancelledStoreOwnerNotification: "RecurringPaymentCancelled.StoreOwnerNotification",
	RecurringPaymentCancelledCustomerNotification:   "RecurringPaymentCancelled.CustomerNotification",
	RecurringPaymentFailedCustomerNotification:      "RecurringPaymentFailed.CustomerNotification",

	// Newsletter
	NewsletterSubscriptionActivationMessage:   "NewsLetterSubscription.ActivationMessage",
	NewsletterSubscriptionDeactivationMessage: "NewsLetterSubscription.DeactivationMessage",

	// To friend
	EmailAFriendMessage:     "Service.EmailAFriend",
	WishlistToFriendMessage: "Wishlist.EmailAFriend",

	// Return requests
	NewReturnRequestStoreOwnerNotification:         "NewReturnRequest.StoreOwnerNotification",
	NewReturnRequestCustomerNotification:           "NewReturnRequest.CustomerNotification",
	ReturnRequestStatusChangedCustomerNotification: "ReturnRequestStatusChanged.CustomerNotification",

	// Forum
	NewForumTopicMessage:       "Forums.NewForumTopic",
	NewForumPostMessage:        "Forums.NewForumPost",
	PrivateMessageNotification: "Customer.NewPM",

	// Misc
	NewVendorAccountApplyStoreOwnerNotification:             "VendorAccountApply.StoreOwnerNotification",
	VendorInformationChangeStoreOwnerNotification:           "VendorInformationChange.StoreOwnerNotification",
	GiftCardNotification:                                    "GiftCard.Notification",
	ProductReviewStoreOwnerNotification:                     "Product.ProductReview",
	ProductReviewReplyCustomerNotification:                  "ProductReview.Reply.CustomerNotification",
	QuantityBelowStoreOwnerNotification:                     "QuantityBelow.StoreOwnerNotification",
	QuantityBelowAttributeCombinationStoreOwnerNotification: "QuantityBelow.AttributeCombination.StoreOwnerNotification",
	QuantityBelowVendorNotification:                         "QuantityBelow.VendorNotification",
	QuantityBelowAttributeCombinationVendorNotification:     "QuantityBelow.AttributeCombination.VendorNotification",
	NewVATSubmittedStoreOwnerNotification:                   "NewVATSubmitted.StoreOwnerNotification",
	BlogCommentStoreOwnerNotification:                       "Blog.BlogComment",
	NewsCommentStoreOwnerNotification:                       "News.NewsComment",
	BackInStockNotification:                                 "Customer.BackInStock",
	ContactUsMessage:                                        "Service.ContactUs",
	ContactVendorMessage:                                    "Service.ContactVendor",
}
