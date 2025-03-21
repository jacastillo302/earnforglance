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

func TestDiscountRequirementUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.DiscountRequirementRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewDiscountRequirementUsecase(mockRepo, timeout)

	discountRequirementID := primitive.NewObjectID().Hex()

	updatedDiscountRequirement := domain.DiscountRequirement{
		ID:                                primitive.NewObjectID(), // Existing ID of the record to update
		DiscountID:                        primitive.NewObjectID(),
		DiscountRequirementRuleSystemName: "CustomerRoleRequirement",
		ParentID:                          new(primitive.ObjectID),
		InteractionTypeID:                 new(int),
		IsGroup:                           true,
		InteractionType:                   new(domain.RequirementGroupInteractionType),
	}
	*updatedDiscountRequirement.ParentID = primitive.NewObjectID()
	*updatedDiscountRequirement.InteractionTypeID = 1
	*updatedDiscountRequirement.InteractionType = 3

	mockRepo.On("FetchByID", mock.Anything, discountRequirementID).Return(updatedDiscountRequirement, nil)

	result, err := usecase.FetchByID(context.Background(), discountRequirementID)

	assert.NoError(t, err)
	assert.Equal(t, updatedDiscountRequirement, result)
	mockRepo.AssertExpectations(t)
}

func TestDiscountRequirementUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.DiscountRequirementRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewDiscountRequirementUsecase(mockRepo, timeout)

	newDiscount := &domain.DiscountRequirement{
		DiscountID:                        primitive.NewObjectID(),
		DiscountRequirementRuleSystemName: "MinimumOrderTotalRequirement",
		ParentID:                          nil,
		InteractionTypeID:                 nil,
		IsGroup:                           false,
		InteractionType:                   nil,
	}

	mockRepo.On("Create", mock.Anything, newDiscount).Return(nil)

	err := usecase.Create(context.Background(), newDiscount)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountRequirementUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.DiscountRequirementRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewDiscountRequirementUsecase(mockRepo, timeout)

	updatedDiscountRequirement := &domain.DiscountRequirement{
		ID:                                primitive.NewObjectID(), // Existing ID of the record to update
		DiscountID:                        primitive.NewObjectID(),
		DiscountRequirementRuleSystemName: "CustomerRoleRequirement",
		ParentID:                          new(primitive.ObjectID),
		InteractionTypeID:                 new(int),
		IsGroup:                           true,
		InteractionType:                   new(domain.RequirementGroupInteractionType),
	}
	*updatedDiscountRequirement.ParentID = primitive.NewObjectID()
	*updatedDiscountRequirement.InteractionTypeID = 1
	*updatedDiscountRequirement.InteractionType = 3

	mockRepo.On("Update", mock.Anything, updatedDiscountRequirement).Return(nil)

	err := usecase.Update(context.Background(), updatedDiscountRequirement)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountRequirementUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.DiscountRequirementRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewDiscountRequirementUsecase(mockRepo, timeout)

	discountRequirementID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, discountRequirementID).Return(nil)

	err := usecase.Delete(context.Background(), discountRequirementID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountRequirementUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.DiscountRequirementRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewDiscountRequirementUsecase(mockRepo, timeout)

	fetchedDiscountRequirements := []domain.DiscountRequirement{
		{
			ID:                                primitive.NewObjectID(),
			DiscountID:                        primitive.NewObjectID(),
			DiscountRequirementRuleSystemName: "MinimumOrderTotalRequirement",
			ParentID:                          nil,
			InteractionTypeID:                 nil,
			IsGroup:                           false,
			InteractionType:                   nil,
		},
		{
			ID:                                primitive.NewObjectID(),
			DiscountID:                        primitive.NewObjectID(),
			DiscountRequirementRuleSystemName: "CustomerRoleRequirement",
			ParentID:                          new(primitive.ObjectID),
			InteractionTypeID:                 new(int),
			IsGroup:                           true,
			InteractionType:                   new(domain.RequirementGroupInteractionType),
		},
	}
	*fetchedDiscountRequirements[1].ParentID = primitive.NewObjectID()
	*fetchedDiscountRequirements[1].InteractionTypeID = 1
	*fetchedDiscountRequirements[1].InteractionType = 2

	mockRepo.On("Fetch", mock.Anything).Return(fetchedDiscountRequirements, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedDiscountRequirements, result)
	mockRepo.AssertExpectations(t)
}
