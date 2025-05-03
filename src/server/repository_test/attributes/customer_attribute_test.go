package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/attributes"
	repository "earnforglance/server/repository/attributes"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultCustomerAttribute struct {
	mock.Mock
}

func (m *MockSingleResultCustomerAttribute) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CustomerAttribute); ok {
		*v.(*domain.CustomerAttribute) = *result
	}
	return args.Error(1)
}

var mockItemCustomerAttribute = &domain.CustomerAttribute{
	ID:                              bson.NewObjectID(), // Existing ID of the record to update
	Name:                            "Preferred Language",
	IsRequired:                      false,
	AttributeControlTypeID:          2,
	DisplayOrder:                    2,
	DefaultValue:                    "English",
	ValidationMinLength:             new(int),
	ValidationMaxLength:             new(int),
	ValidationFileAllowedExtensions: ".txt,.pdf",
	ValidationFileMaximumSize:       new(int),
	ConditionAttributeXml:           "<conditions><required>false</required></conditions>",
}

func TestCustomerAttributeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCustomerAttribute

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCustomerAttribute, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerAttribute.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerAttribute.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCustomerAttributeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCustomerAttribute).Return(nil, nil).Once()

	repo := repository.NewCustomerAttributeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCustomerAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCustomerAttributeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCustomerAttribute.ID}
	update := bson.M{"$set": mockItemCustomerAttribute}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCustomerAttributeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCustomerAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
