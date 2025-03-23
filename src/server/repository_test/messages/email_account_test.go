package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/messages"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultEmailAccount struct {
	mock.Mock
}

func (m *MockSingleResultEmailAccount) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.EmailAccount); ok {
		*v.(*domain.EmailAccount) = *result
	}
	return args.Error(1)
}

var mockItemEmailAccount = &domain.EmailAccount{
	ID:                          primitive.NewObjectID(), // Existing ID of the record to update
	Email:                       "updated@example.com",
	DisplayName:                 "Updated Account",
	Host:                        "smtp.updated.com",
	Port:                        465,
	Username:                    "updated_user",
	Password:                    "updatedpassword",
	EnableSsl:                   false,
	MaxNumberOfEmails:           200,
	EmailAuthenticationMethodID: 2,
	ClientID:                    "updated-client-id",
	ClientSecret:                "updated-client-secret",
	TenantID:                    "updated-tenant-id",
	EmailAuthenticationMethod:   1,
}

func TestEmailAccountRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionEmailAccount

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultEmailAccount{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemEmailAccount, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewEmailAccountRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemEmailAccount.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultEmailAccount{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewEmailAccountRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemEmailAccount.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestEmailAccountRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionEmailAccount

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemEmailAccount).Return(nil, nil).Once()

	repo := repository.NewEmailAccountRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemEmailAccount)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestEmailAccountRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionEmailAccount

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemEmailAccount.ID}
	update := bson.M{"$set": mockItemEmailAccount}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewEmailAccountRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemEmailAccount)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
