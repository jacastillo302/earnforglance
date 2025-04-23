package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/messages"
	repository "earnforglance/server/repository/messages"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultCampaign struct {
	mock.Mock
}

func (m *MockSingleResultCampaign) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Campaign); ok {
		*v.(*domain.Campaign) = *result
	}
	return args.Error(1)
}

var mockItemCampaign = &domain.Campaign{
	ID:                    bson.NewObjectID(), // Existing ID of the record to update
	Name:                  "Updated Holiday Sale",
	Subject:               "Updated Discounts for the Holidays!",
	Body:                  "Enjoy up to 60% off on selected items. Offer extended!",
	StoreID:               bson.NewObjectID(),
	CustomerRoleID:        bson.NewObjectID(),
	CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
	DontSendBeforeDateUtc: new(time.Time),
}

func TestCampaignRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCampaign

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCampaign{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCampaign, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCampaignRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCampaign.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCampaign{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCampaignRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCampaign.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCampaignRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCampaign

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCampaign).Return(nil, nil).Once()

	repo := repository.NewCampaignRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCampaign)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCampaignRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCampaign

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCampaign.ID}
	update := bson.M{"$set": mockItemCampaign}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCampaignRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCampaign)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
