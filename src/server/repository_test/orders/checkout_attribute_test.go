package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/orders"
	repository "earnforglance/server/repository/orders"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultCheckoutAttribute struct {
	mock.Mock
}

func (m *MockSingleResultCheckoutAttribute) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CheckoutAttribute); ok {
		*v.(*domain.CheckoutAttribute) = *result
	}
	return args.Error(1)
}

var mockItemCheckoutAttribute = &domain.CheckoutAttribute{
	ID:                              bson.NewObjectID(), // Existing ID of the record to update
	TextPrompt:                      "Update your custom message",
	ShippableProductRequired:        false,
	IsTaxExempt:                     true,
	TaxCategoryID:                   bson.NewObjectID(),
	LimitedToStores:                 true,
	ValidationMinLength:             new(int),
	ValidationMaxLength:             new(int),
	ValidationFileAllowedExtensions: ".pdf,.docx",
	ValidationFileMaximumSize:       new(int),
	DefaultValue:                    "Updated Default Message",
	ConditionAttributeXml:           "<UpdatedConditions></UpdatedConditions>",
}

func TestCheckoutAttributeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCheckoutAttribute

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCheckoutAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCheckoutAttribute, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCheckoutAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCheckoutAttribute.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCheckoutAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCheckoutAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCheckoutAttribute.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCheckoutAttributeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCheckoutAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCheckoutAttribute).Return(nil, nil).Once()

	repo := repository.NewCheckoutAttributeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCheckoutAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCheckoutAttributeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCheckoutAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCheckoutAttribute.ID}
	update := bson.M{"$set": mockItemCheckoutAttribute}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCheckoutAttributeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCheckoutAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
