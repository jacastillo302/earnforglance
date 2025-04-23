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

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultExternalAuthenticationRecord struct {
	mock.Mock
}

func (m *MockSingleResultExternalAuthenticationRecord) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ExternalAuthenticationRecord); ok {
		*v.(*domain.ExternalAuthenticationRecord) = *result
	}
	return args.Error(1)
}

var mockItemExternalAuthenticationRecord = &domain.ExternalAuthenticationRecord{
	ID:                        bson.NewObjectID(),
	CustomerID:                bson.NewObjectID(),
	Email:                     "user1@example.com",
	ExternalIdentifier:        "external-id-67890",
	ExternalDisplayIdentifier: "ExternalUser1",
	OAuthToken:                "oauth-token-def456",
	OAuthAccessToken:          "oauth-access-token-uvw",
	ProviderSystemName:        "Facebook",
}

func TestExternalAuthenticationRecordRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionExternalAuthenticationRecord

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultExternalAuthenticationRecord{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemExternalAuthenticationRecord, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewExternalAuthenticationRecordRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemExternalAuthenticationRecord.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultExternalAuthenticationRecord{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewExternalAuthenticationRecordRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemExternalAuthenticationRecord.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestExternalAuthenticationRecordRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionExternalAuthenticationRecord

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemExternalAuthenticationRecord).Return(nil, nil).Once()

	repo := repository.NewExternalAuthenticationRecordRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemExternalAuthenticationRecord)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestExternalAuthenticationRecordRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionExternalAuthenticationRecord

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemExternalAuthenticationRecord.ID}
	update := bson.M{"$set": mockItemExternalAuthenticationRecord}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewExternalAuthenticationRecordRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemExternalAuthenticationRecord)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
