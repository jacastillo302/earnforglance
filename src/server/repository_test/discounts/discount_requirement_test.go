package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/discounts"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/discounts"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultDiscountRequirement struct {
	mock.Mock
}

func (m *MockSingleResultDiscountRequirement) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.DiscountRequirement); ok {
		*v.(*domain.DiscountRequirement) = *result
	}
	return args.Error(1)
}

var mockItemDiscountRequirement = &domain.DiscountRequirement{
	ID:                                primitive.NewObjectID(), // Existing ID of the record to update
	DiscountID:                        primitive.NewObjectID(),
	DiscountRequirementRuleSystemName: "CustomerRoleRequirement",
	ParentID:                          new(primitive.ObjectID),
	InteractionTypeID:                 new(int),
	IsGroup:                           true,
}

func TestDiscountRequirementRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionDiscountRequirement

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscountRequirement{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemDiscountRequirement, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountRequirementRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscountRequirement.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscountRequirement{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountRequirementRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscountRequirement.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestDiscountRequirementRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscountRequirement

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemDiscountRequirement).Return(nil, nil).Once()

	repo := repository.NewDiscountRequirementRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemDiscountRequirement)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestDiscountRequirementRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscountRequirement

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemDiscountRequirement.ID}
	update := bson.M{"$set": mockItemDiscountRequirement}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewDiscountRequirementRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemDiscountRequirement)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
