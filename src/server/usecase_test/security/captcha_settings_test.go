package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/security"
	test "earnforglance/server/usecase/security"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCaptchaSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CaptchaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCaptchaSettingsUsecase(mockRepo, timeout)

	securityID := primitive.NewObjectID().Hex()

	updatedCaptchaSettings := domain.CaptchaSettings{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		Enabled:                         false,
		CaptchaTypeID:                   2,
		ShowOnLoginPage:                 false,
		ShowOnRegistrationPage:          false,
		ShowOnContactUsPage:             true,
		ShowOnEmailWishlistToFriendPage: false,
		ShowOnEmailProductToFriendPage:  true,
		ShowOnBlogCommentPage:           false,
		ShowOnNewsCommentPage:           true,
		ShowOnNewsletterPage:            false,
		ShowOnProductReviewPage:         false,
		ShowOnApplyVendorPage:           true,
		ShowOnForgotPasswordPage:        false,
		ShowOnForum:                     true,
		ShowOnCheckoutPageForGuests:     false,
		ShowOnCheckGiftCardBalance:      true,
		ReCaptchaApiUrl:                 "https://www.google.com/recaptcha/api.js",
		ReCaptchaPublicKey:              "6Lc_aXkUAAAAAABC9876543210",
		ReCaptchaPrivateKey:             "6Lc_aXkUAAAAADEF9876543210",
		ReCaptchaV3ScoreThreshold:       0.7,
		ReCaptchaTheme:                  "dark",
		ReCaptchaRequestTimeout:         new(int),
		ReCaptchaDefaultLanguage:        "fr",
		AutomaticallyChooseLanguage:     false,
	}
	*updatedCaptchaSettings.ReCaptchaRequestTimeout = 10000

	mockRepo.On("FetchByID", mock.Anything, securityID).Return(updatedCaptchaSettings, nil)

	result, err := usecase.FetchByID(context.Background(), securityID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCaptchaSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestCaptchaSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CaptchaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCaptchaSettingsUsecase(mockRepo, timeout)

	newCaptchaSettings := &domain.CaptchaSettings{
		Enabled:                         true,
		CaptchaTypeID:                   3,
		ShowOnLoginPage:                 true,
		ShowOnRegistrationPage:          true,
		ShowOnContactUsPage:             false,
		ShowOnEmailWishlistToFriendPage: true,
		ShowOnEmailProductToFriendPage:  false,
		ShowOnBlogCommentPage:           true,
		ShowOnNewsCommentPage:           false,
		ShowOnNewsletterPage:            true,
		ShowOnProductReviewPage:         true,
		ShowOnApplyVendorPage:           false,
		ShowOnForgotPasswordPage:        true,
		ShowOnForum:                     false,
		ShowOnCheckoutPageForGuests:     true,
		ShowOnCheckGiftCardBalance:      false,
		ReCaptchaApiUrl:                 "https://www.google.com/recaptcha/api.js",
		ReCaptchaPublicKey:              "6Lc_aXkUAAAAAABC1234567890",
		ReCaptchaPrivateKey:             "6Lc_aXkUAAAAADEF1234567890",
		ReCaptchaV3ScoreThreshold:       0.5,
		ReCaptchaTheme:                  "light",
		ReCaptchaRequestTimeout:         new(int),
		ReCaptchaDefaultLanguage:        "en",
		AutomaticallyChooseLanguage:     true,
	}
	*newCaptchaSettings.ReCaptchaRequestTimeout = 5000

	mockRepo.On("Create", mock.Anything, newCaptchaSettings).Return(nil)

	err := usecase.Create(context.Background(), newCaptchaSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCaptchaSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CaptchaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCaptchaSettingsUsecase(mockRepo, timeout)

	updatedCaptchaSettings := &domain.CaptchaSettings{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		Enabled:                         false,
		CaptchaTypeID:                   1,
		ShowOnLoginPage:                 false,
		ShowOnRegistrationPage:          false,
		ShowOnContactUsPage:             true,
		ShowOnEmailWishlistToFriendPage: false,
		ShowOnEmailProductToFriendPage:  true,
		ShowOnBlogCommentPage:           false,
		ShowOnNewsCommentPage:           true,
		ShowOnNewsletterPage:            false,
		ShowOnProductReviewPage:         false,
		ShowOnApplyVendorPage:           true,
		ShowOnForgotPasswordPage:        false,
		ShowOnForum:                     true,
		ShowOnCheckoutPageForGuests:     false,
		ShowOnCheckGiftCardBalance:      true,
		ReCaptchaApiUrl:                 "https://www.google.com/recaptcha/api.js",
		ReCaptchaPublicKey:              "6Lc_aXkUAAAAAABC9876543210",
		ReCaptchaPrivateKey:             "6Lc_aXkUAAAAADEF9876543210",
		ReCaptchaV3ScoreThreshold:       0.7,
		ReCaptchaTheme:                  "dark",
		ReCaptchaRequestTimeout:         new(int),
		ReCaptchaDefaultLanguage:        "fr",
		AutomaticallyChooseLanguage:     false,
	}
	*updatedCaptchaSettings.ReCaptchaRequestTimeout = 10000

	mockRepo.On("Update", mock.Anything, updatedCaptchaSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedCaptchaSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCaptchaSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CaptchaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCaptchaSettingsUsecase(mockRepo, timeout)

	securityID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, securityID).Return(nil)

	err := usecase.Delete(context.Background(), securityID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCaptchaSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CaptchaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCaptchaSettingsUsecase(mockRepo, timeout)

	fetchedCaptchaSettings := []domain.CaptchaSettings{
		{
			ID:                              primitive.NewObjectID(),
			Enabled:                         true,
			CaptchaTypeID:                   3,
			ShowOnLoginPage:                 true,
			ShowOnRegistrationPage:          true,
			ShowOnContactUsPage:             false,
			ShowOnEmailWishlistToFriendPage: true,
			ShowOnEmailProductToFriendPage:  false,
			ShowOnBlogCommentPage:           true,
			ShowOnNewsCommentPage:           false,
			ShowOnNewsletterPage:            true,
			ShowOnProductReviewPage:         true,
			ShowOnApplyVendorPage:           false,
			ShowOnForgotPasswordPage:        true,
			ShowOnForum:                     false,
			ShowOnCheckoutPageForGuests:     true,
			ShowOnCheckGiftCardBalance:      false,
			ReCaptchaApiUrl:                 "https://www.google.com/recaptcha/api.js",
			ReCaptchaPublicKey:              "6Lc_aXkUAAAAAABC1234567890",
			ReCaptchaPrivateKey:             "6Lc_aXkUAAAAADEF1234567890",
			ReCaptchaV3ScoreThreshold:       0.5,
			ReCaptchaTheme:                  "light",
			ReCaptchaRequestTimeout:         new(int),
			ReCaptchaDefaultLanguage:        "en",
			AutomaticallyChooseLanguage:     true,
		},
		{
			ID:                              primitive.NewObjectID(),
			Enabled:                         false,
			CaptchaTypeID:                   2,
			ShowOnLoginPage:                 false,
			ShowOnRegistrationPage:          false,
			ShowOnContactUsPage:             true,
			ShowOnEmailWishlistToFriendPage: false,
			ShowOnEmailProductToFriendPage:  true,
			ShowOnBlogCommentPage:           false,
			ShowOnNewsCommentPage:           true,
			ShowOnNewsletterPage:            false,
			ShowOnProductReviewPage:         false,
			ShowOnApplyVendorPage:           true,
			ShowOnForgotPasswordPage:        false,
			ShowOnForum:                     true,
			ShowOnCheckoutPageForGuests:     false,
			ShowOnCheckGiftCardBalance:      true,
			ReCaptchaApiUrl:                 "https://www.google.com/recaptcha/api.js",
			ReCaptchaPublicKey:              "6Lc_aXkUAAAAAABC9876543210",
			ReCaptchaPrivateKey:             "6Lc_aXkUAAAAADEF9876543210",
			ReCaptchaV3ScoreThreshold:       0.7,
			ReCaptchaTheme:                  "dark",
			ReCaptchaRequestTimeout:         new(int),
			ReCaptchaDefaultLanguage:        "fr",
			AutomaticallyChooseLanguage:     false,
		},
	}
	*fetchedCaptchaSettings[0].ReCaptchaRequestTimeout = 5000
	*fetchedCaptchaSettings[1].ReCaptchaRequestTimeout = 10000

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCaptchaSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCaptchaSettings, result)
	mockRepo.AssertExpectations(t)
}
