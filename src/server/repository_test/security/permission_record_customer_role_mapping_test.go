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

type MockSingleResultPermissionRecordCustomerRoleMapping struct {
	mock.Mock
}

func (m *MockSingleResultPermissionRecordCustomerRoleMapping) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PermissionRecordCustomerRoleMapping); ok {
		*v.(*domain.PermissionRecordCustomerRoleMapping) = *result
	}
	return args.Error(1)
}

var mockItemPermissionRecordCustomerRoleMapping = &domain.PermissionRecordCustomerRoleMapping{
	ID:                 primitive.NewObjectID(), // Existing ID of the record to update
	PermissionRecordID: primitive.NewObjectID(),
	CustomerRoleID:     primitive.NewObjectID(),
}

func TestPermissionRecordCustomerRoleMappingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPermissionRecordCustomerRoleMapping

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPermissionRecordCustomerRoleMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPermissionRecordCustomerRoleMapping, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPermissionRecordCustomerRoleMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPermissionRecordCustomerRoleMapping.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPermissionRecordCustomerRoleMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPermissionRecordCustomerRoleMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPermissionRecordCustomerRoleMapping.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPermissionRecordCustomerRoleMappingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPermissionRecordCustomerRoleMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPermissionRecordCustomerRoleMapping).Return(nil, nil).Once()

	repo := repository.NewPermissionRecordCustomerRoleMappingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPermissionRecordCustomerRoleMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPermissionRecordCustomerRoleMappingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPermissionRecordCustomerRoleMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPermissionRecordCustomerRoleMapping.ID}
	update := bson.M{"$set": mockItemPermissionRecordCustomerRoleMapping}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPermissionRecordCustomerRoleMappingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPermissionRecordCustomerRoleMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
