package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/security"
	test "earnforglance/server/usecase/security"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestAclRecordUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.AclRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewAclRecordUsecase(mockRepo, timeout)

	aclRecordID := bson.NewObjectID().Hex()
	updatedAclRecord := domain.AclRecord{
		EntityID:       bson.NewObjectID(),
		EntityName:     "Category",
		CustomerRoleID: bson.NewObjectID(),
		IsRead:         true,
		IsDelete:       true,
		IsUpdate:       true,
		IsCreate:       true,
	}

	mockRepo.On("FetchByID", mock.Anything, aclRecordID).Return(updatedAclRecord, nil)

	result, err := usecase.FetchByID(context.Background(), aclRecordID)

	assert.NoError(t, err)
	assert.Equal(t, updatedAclRecord, result)
	mockRepo.AssertExpectations(t)
}

func TestAclRecordUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.AclRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewAclRecordUsecase(mockRepo, timeout)

	newAclRecord := &domain.AclRecord{
		EntityID:       bson.NewObjectID(),
		EntityName:     "Product",
		CustomerRoleID: bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newAclRecord).Return(nil)

	err := usecase.Create(context.Background(), newAclRecord)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAclRecordUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.AclRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewAclRecordUsecase(mockRepo, timeout)

	updatedAclRecord := &domain.AclRecord{
		EntityID:       bson.NewObjectID(),
		EntityName:     "Category",
		CustomerRoleID: bson.NewObjectID(),
		IsRead:         true,
		IsDelete:       true,
		IsUpdate:       true,
		IsCreate:       true,
	}

	mockRepo.On("Update", mock.Anything, updatedAclRecord).Return(nil)

	err := usecase.Update(context.Background(), updatedAclRecord)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAclRecordUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.AclRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewAclRecordUsecase(mockRepo, timeout)

	aclRecordID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, aclRecordID).Return(nil)

	err := usecase.Delete(context.Background(), aclRecordID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAclRecordUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.AclRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewAclRecordUsecase(mockRepo, timeout)

	fetchedAclRecords := []domain.AclRecord{
		{
			EntityID:       bson.NewObjectID(),
			EntityName:     "Product",
			CustomerRoleID: bson.NewObjectID(),
			IsRead:         true,
			IsDelete:       true,
			IsUpdate:       true,
			IsCreate:       true,
		},
		{
			EntityID:       bson.NewObjectID(),
			EntityName:     "Category",
			CustomerRoleID: bson.NewObjectID(),
			IsRead:         true,
			IsDelete:       false,
			IsUpdate:       true,
			IsCreate:       false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedAclRecords, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedAclRecords, result)
	mockRepo.AssertExpectations(t)
}
