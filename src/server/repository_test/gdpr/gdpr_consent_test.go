package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/gdpr"
	repository "earnforglance/server/repository/gdpr"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultGdprConsent struct {
	mock.Mock
}

func (m *MockSingleResultGdprConsent) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.GdprConsent); ok {
		*v.(*domain.GdprConsent) = *result
	}
	return args.Error(1)
}

var mockItemGdprConsent = &domain.GdprConsent{
	ID:                        primitive.NewObjectID(), // Existing ID of the record to update
	Message:                   "Updated GDPR consent message.",
	IsRequired:                false,
	RequiredMessage:           "Consent is optional.",
	DisplayDuringRegistration: false,
	DisplayOnCustomerInfoPage: true,
	DisplayOrder:              2,
}

func TestGdprConsentRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionGdprConsent

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultGdprConsent{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemGdprConsent, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewGdprConsentRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemGdprConsent.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultGdprConsent{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewGdprConsentRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemGdprConsent.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestGdprConsentRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionGdprConsent

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemGdprConsent).Return(nil, nil).Once()

	repo := repository.NewGdprConsentRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemGdprConsent)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestGdprConsentRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionGdprConsent

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemGdprConsent.ID}
	update := bson.M{"$set": mockItemGdprConsent}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewGdprConsentRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemGdprConsent)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
