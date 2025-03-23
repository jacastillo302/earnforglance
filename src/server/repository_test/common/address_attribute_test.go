package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultAddressAttribute struct {
	mock.Mock
}

func (m *MockSingleResultAddressAttribute) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.AddressAttribute); ok {
		*v.(*domain.AddressAttribute) = *result
	}
	return args.Error(1)
}

var mockItemAddressAttribute = &domain.AddressAttribute{
	ID:                              primitive.NewObjectID(), // Existing ID of the record to update
	Name:                            "City",
	IsRequired:                      false,
	AttributeControlTypeID:          2,
	DisplayOrder:                    2,
	DefaultValue:                    "New York",
	ValidationMinLength:             new(int),
	ValidationMaxLength:             new(int),
	ValidationFileAllowedExtensions: ".jpg,.png",
	ValidationFileMaximumSize:       new(int),
	ConditionAttributeXml:           "<conditions><required>false</required></conditions>",
}

func TestAddressAttributeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionAddressAttribute

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultAddressAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemAddressAttribute, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewAddressAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemAddressAttribute.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultAddressAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewAddressAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemAddressAttribute.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestAddressAttributeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionAddressAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemAddressAttribute).Return(nil, nil).Once()

	repo := repository.NewAddressAttributeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemAddressAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestAddressAttributeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionAddressAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemAddressAttribute.ID}
	update := bson.M{"$set": mockItemAddressAttribute}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewAddressAttributeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemAddressAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
