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

type MockSingleResultPermisionRecordAttribute struct {
	mock.Mock
}

func (m *MockSingleResultPermisionRecordAttribute) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PermisionRecordAttribute); ok {
		*v.(*domain.PermisionRecordAttribute) = *result
	}
	return args.Error(1)
}

var mockItemPermisionRecordAttribute = &domain.PermisionRecordAttribute{
	ID:                              bson.NewObjectID(), // Existing ID of the record to update
	Name:                            "Preferred Language",
	IsRequired:                      false,
	DisplayOrder:                    2,
	DefaultValue:                    "English",
	ValidationMinLength:             new(int),
	ValidationMaxLength:             new(int),
	ValidationFileAllowedExtensions: ".txt,.pdf",
	ValidationFileMaximumSize:       new(int),
	ConditionAttributeXml:           "<conditions><required>false</required></conditions>",
}

func TestPermisionRecordAttributeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPermisionRecordAttribute

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPermisionRecordAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPermisionRecordAttribute, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPermisionRecordAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPermisionRecordAttribute.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPermisionRecordAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPermisionRecordAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPermisionRecordAttribute.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPermisionRecordAttributeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPermisionRecordAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPermisionRecordAttribute).Return(nil, nil).Once()

	repo := repository.NewPermisionRecordAttributeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPermisionRecordAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPermisionRecordAttributeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPermisionRecordAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPermisionRecordAttribute.ID}
	update := bson.M{"$set": mockItemPermisionRecordAttribute}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPermisionRecordAttributeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPermisionRecordAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
