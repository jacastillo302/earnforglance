package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionCaptchaSettings = "captcha_settings"
)

// CaptchaSettings represents CAPTCHA settings
type CaptchaSettings struct {
	ID                              primitive.ObjectID `bson:"_id,omitempty"`
	Enabled                         bool               `bson:"enabled"`
	CaptchaType                     CaptchaType        `bson:"captcha_type"`
	ShowOnLoginPage                 bool               `bson:"show_on_login_page"`
	ShowOnRegistrationPage          bool               `bson:"show_on_registration_page"`
	ShowOnContactUsPage             bool               `bson:"show_on_contact_us_page"`
	ShowOnEmailWishlistToFriendPage bool               `bson:"show_on_email_wishlist_to_friend_page"`
	ShowOnEmailProductToFriendPage  bool               `bson:"show_on_email_product_to_friend_page"`
	ShowOnBlogCommentPage           bool               `bson:"show_on_blog_comment_page"`
	ShowOnNewsCommentPage           bool               `bson:"show_on_news_comment_page"`
	ShowOnNewsletterPage            bool               `bson:"show_on_newsletter_page"`
	ShowOnProductReviewPage         bool               `bson:"show_on_product_review_page"`
	ShowOnApplyVendorPage           bool               `bson:"show_on_apply_vendor_page"`
	ShowOnForgotPasswordPage        bool               `bson:"show_on_forgot_password_page"`
	ShowOnForum                     bool               `bson:"show_on_forum"`
	ShowOnCheckoutPageForGuests     bool               `bson:"show_on_checkout_page_for_guests"`
	ShowOnCheckGiftCardBalance      bool               `bson:"show_on_check_gift_card_balance"`
	ReCaptchaApiUrl                 string             `bson:"re_captcha_api_url"`
	ReCaptchaPublicKey              string             `bson:"re_captcha_public_key"`
	ReCaptchaPrivateKey             string             `bson:"re_captcha_private_key"`
	ReCaptchaV3ScoreThreshold       float64            `bson:"re_captcha_v3_score_threshold"`
	ReCaptchaTheme                  string             `bson:"re_captcha_theme"`
	ReCaptchaRequestTimeout         *int               `bson:"re_captcha_request_timeout,omitempty"`
	ReCaptchaDefaultLanguage        string             `bson:"re_captcha_default_language"`
	AutomaticallyChooseLanguage     bool               `bson:"automatically_choose_language"`
}
