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

func TestDiscountProductMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.DiscountProductMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountProductMappingUsecase(mockRepo, timeout)

	discountID := primitive.NewObjectID().Hex()

	updatedDiscountProductMapping := domain.DiscountProductMapping{
		DiscountMapping: domain.DiscountMapping{
			DiscountID: primitive.NewObjectID(),
		},
		EntityID: primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, discountID).Return(updatedDiscountProductMapping, nil)

	result, err := usecase.FetchByID(context.Background(), discountID)

	assert.NoError(t, err)
	assert.Equal(t, updatedDiscountProductMapping, result)
	mockRepo.AssertExpectations(t)
}

func TestDiscountProductMappingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.DiscountProductMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountProductMappingUsecase(mockRepo, timeout)

	newDiscountProductMapping := &domain.DiscountProductMapping{
		DiscountMapping: domain.DiscountMapping{
			DiscountID: primitive.NewObjectID(),
		},
		EntityID: primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newDiscountProductMapping).Return(nil)

	err := usecase.Create(context.Background(), newDiscountProductMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountProductMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.DiscountProductMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountProductMappingUsecase(mockRepo, timeout)

	updatedDiscountProductMapping := &domain.DiscountProductMapping{
		DiscountMapping: domain.DiscountMapping{
			DiscountID: primitive.NewObjectID(),
		},
		EntityID: primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedDiscountProductMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedDiscountProductMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountProductMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.DiscountProductMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountProductMappingUsecase(mockRepo, timeout)

	discountID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, discountID).Return(nil)

	err := usecase.Delete(context.Background(), discountID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountProductMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.DiscountProductMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountProductMappingUsecase(mockRepo, timeout)

	fetchedDiscountProductMappings := []domain.DiscountProductMapping{
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

	mockRepo.On("Fetch", mock.Anything).Return(fetchedDiscountProductMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedDiscountProductMappings, result)
	mockRepo.AssertExpectations(t)
}
