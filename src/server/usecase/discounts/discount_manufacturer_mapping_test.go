package usecase

import (
	"context"
	domain "earnforglance/server/domain/discounts"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDiscountManufacturerMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.DiscountManufacturerMappingRepository)
	timeout := time.Duration(10)
	usecase := NewDiscountManufacturerMappingUsecase(mockRepo, timeout)

	discountID := primitive.NewObjectID().Hex()

	updatedDiscountManufacturerMapping := domain.DiscountManufacturerMapping{
		DiscountMapping: domain.DiscountMapping{
			DiscountID: primitive.NewObjectID(),
		},
		EntityID: primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, discountID).Return(updatedDiscountManufacturerMapping, nil)

	result, err := usecase.FetchByID(context.Background(), discountID)

	assert.NoError(t, err)
	assert.Equal(t, updatedDiscountManufacturerMapping, result)
	mockRepo.AssertExpectations(t)
}

func TestDiscountManufacturerMappingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.DiscountManufacturerMappingRepository)
	timeout := time.Duration(10)
	usecase := NewDiscountManufacturerMappingUsecase(mockRepo, timeout)

	newDiscountManufacturerMapping := &domain.DiscountManufacturerMapping{
		DiscountMapping: domain.DiscountMapping{
			DiscountID: primitive.NewObjectID(),
		},
		EntityID: primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newDiscountManufacturerMapping).Return(nil)

	err := usecase.Create(context.Background(), newDiscountManufacturerMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountManufacturerMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.DiscountManufacturerMappingRepository)
	timeout := time.Duration(10)
	usecase := NewDiscountManufacturerMappingUsecase(mockRepo, timeout)

	updatedDiscountManufacturerMapping := &domain.DiscountManufacturerMapping{
		DiscountMapping: domain.DiscountMapping{
			DiscountID: primitive.NewObjectID(),
		},
		EntityID: primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedDiscountManufacturerMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedDiscountManufacturerMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountManufacturerMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.DiscountManufacturerMappingRepository)
	timeout := time.Duration(10)
	usecase := NewDiscountManufacturerMappingUsecase(mockRepo, timeout)

	discountID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, discountID).Return(nil)

	err := usecase.Delete(context.Background(), discountID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountManufacturerMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.DiscountManufacturerMappingRepository)
	timeout := time.Duration(10)
	usecase := NewDiscountManufacturerMappingUsecase(mockRepo, timeout)

	fetchedDiscountManufacturerMappings := []domain.DiscountManufacturerMapping{
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

	mockRepo.On("Fetch", mock.Anything).Return(fetchedDiscountManufacturerMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedDiscountManufacturerMappings, result)
	mockRepo.AssertExpectations(t)
}
