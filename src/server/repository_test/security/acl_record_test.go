package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/security"
	repository "earnforglance/server/repository/security"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultAclRecord struct {
	mock.Mock
}

func (m *MockSingleResultAclRecord) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.AclRecord); ok {
		*v.(*domain.AclRecord) = *result
	}
	return args.Error(1)
}

var mockItemAclRecord = &domain.AclRecord{
	EntityID:       bson.NewObjectID(),
	EntityName:     "Category",
	CustomerRoleID: bson.NewObjectID(),
	IsRead:         true,
	IsDelete:       true,
	IsUpdate:       true,
	IsCreate:       true,
}

func TestAclRecordRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionAclRecord

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultAclRecord{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemAclRecord, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewAclRecordRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemAclRecord.EntityID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultAclRecord{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewAclRecordRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemAclRecord.EntityID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestAclRecordRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionAclRecord

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemAclRecord).Return(nil, nil).Once()

	repo := repository.NewAclRecordRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemAclRecord)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestAclRecordRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionAclRecord

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemAclRecord.EntityID}
	update := bson.M{"$set": mockItemAclRecord}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewAclRecordRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemAclRecord)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
