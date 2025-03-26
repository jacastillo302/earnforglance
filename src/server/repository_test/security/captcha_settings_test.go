package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/security"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/security"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultCaptchaSettings struct {
	mock.Mock
}

func (m *MockSingleResultCaptchaSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CaptchaSettings); ok {
		*v.(*domain.CaptchaSettings) = *result
	}
	return args.Error(1)
}

var mockItemCaptchaSettings = &domain.CaptchaSettings{
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

func TestCaptchaSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCaptchaSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCaptchaSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCaptchaSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCaptchaSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCaptchaSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCaptchaSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCaptchaSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCaptchaSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCaptchaSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCaptchaSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCaptchaSettings).Return(nil, nil).Once()

	repo := repository.NewCaptchaSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCaptchaSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCaptchaSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCaptchaSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCaptchaSettings.ID}
	update := bson.M{"$set": mockItemCaptchaSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCaptchaSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCaptchaSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
