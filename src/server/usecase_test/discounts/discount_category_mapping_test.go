package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/discounts"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/discounts"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDiscountCategoryMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.DiscountCategoryMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountCategoryMappingUsecase(mockRepo, timeout)

	discountCategoryMappingID := primitive.NewObjectID().Hex()

	updatedDiscountCategoryMapping := domain.DiscountCategoryMapping{
		DiscountMapping: domain.DiscountMapping{
			DiscountID: primitive.NewObjectID(),
		},
		EntityID: primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, discountCategoryMappingID).Return(updatedDiscountCategoryMapping, nil)

	result, err := usecase.FetchByID(context.Background(), discountCategoryMappingID)

	assert.NoError(t, err)
	assert.Equal(t, updatedDiscountCategoryMapping, result)
	mockRepo.AssertExpectations(t)
}

func TestDiscountCategoryMappingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.DiscountCategoryMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountCategoryMappingUsecase(mockRepo, timeout)

	newDiscountCategoryMapping := &domain.DiscountCategoryMapping{
		DiscountMapping: domain.DiscountMapping{
			DiscountID: primitive.NewObjectID(),
		},
		EntityID: primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newDiscountCategoryMapping).Return(nil)

	err := usecase.Create(context.Background(), newDiscountCategoryMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountCategoryMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.DiscountCategoryMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountCategoryMappingUsecase(mockRepo, timeout)

	updatedDiscountCategoryMapping := &domain.DiscountCategoryMapping{
		DiscountMapping: domain.DiscountMapping{
			DiscountID: primitive.NewObjectID(),
		},
		EntityID: primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedDiscountCategoryMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedDiscountCategoryMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountCategoryMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.DiscountCategoryMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountCategoryMappingUsecase(mockRepo, timeout)

	discountCategoryMappingID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, discountCategoryMappingID).Return(nil)

	err := usecase.Delete(context.Background(), discountCategoryMappingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountCategoryMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.DiscountCategoryMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountCategoryMappingUsecase(mockRepo, timeout)

	fetchedDiscountCategoryMappings := []domain.DiscountCategoryMapping{
		{
			DiscountMapping: domain.DiscountMapping{
				DiscountID: primitive.NewObjectID(),
			},
			EntityID: primitive.NewObjectID(),
		},
		{
			DiscountMapping: domain.DiscountMapping{
				DiscountID: primitive.NewObjectID(),
			},
			EntityID: primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedDiscountCategoryMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedDiscountCategoryMappings, result)
	mockRepo.AssertExpectations(t)
}
