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
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestDiscountRequirementUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.DiscountRequirementRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountRequirementUsecase(mockRepo, timeout)

	discountRequirementID := bson.NewObjectID().Hex()

	updatedDiscountRequirement := domain.DiscountRequirement{
		ID:                                bson.NewObjectID(), // Existing ID of the record to update
		DiscountID:                        bson.NewObjectID(),
		DiscountRequirementRuleSystemName: "CustomerRoleRequirement",
		ParentID:                          new(bson.ObjectID),
		InteractionTypeID:                 new(int),
		IsGroup:                           true,
	}
	*updatedDiscountRequirement.ParentID = bson.NewObjectID()
	*updatedDiscountRequirement.InteractionTypeID = 1

	mockRepo.On("FetchByID", mock.Anything, discountRequirementID).Return(updatedDiscountRequirement, nil)

	result, err := usecase.FetchByID(context.Background(), discountRequirementID)

	assert.NoError(t, err)
	assert.Equal(t, updatedDiscountRequirement, result)
	mockRepo.AssertExpectations(t)
}

func TestDiscountRequirementUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.DiscountRequirementRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountRequirementUsecase(mockRepo, timeout)

	newDiscount := &domain.DiscountRequirement{
		DiscountID:                        bson.NewObjectID(),
		DiscountRequirementRuleSystemName: "MinimumOrderTotalRequirement",
		ParentID:                          nil,
		InteractionTypeID:                 nil,
		IsGroup:                           false,
	}

	mockRepo.On("Create", mock.Anything, newDiscount).Return(nil)

	err := usecase.Create(context.Background(), newDiscount)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountRequirementUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.DiscountRequirementRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountRequirementUsecase(mockRepo, timeout)

	updatedDiscountRequirement := &domain.DiscountRequirement{
		ID:                                bson.NewObjectID(), // Existing ID of the record to update
		DiscountID:                        bson.NewObjectID(),
		DiscountRequirementRuleSystemName: "CustomerRoleRequirement",
		ParentID:                          new(bson.ObjectID),
		InteractionTypeID:                 new(int),
		IsGroup:                           true,
	}
	*updatedDiscountRequirement.ParentID = bson.NewObjectID()
	*updatedDiscountRequirement.InteractionTypeID = 1

	mockRepo.On("Update", mock.Anything, updatedDiscountRequirement).Return(nil)

	err := usecase.Update(context.Background(), updatedDiscountRequirement)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountRequirementUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.DiscountRequirementRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountRequirementUsecase(mockRepo, timeout)

	discountRequirementID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, discountRequirementID).Return(nil)

	err := usecase.Delete(context.Background(), discountRequirementID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountRequirementUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.DiscountRequirementRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewDiscountRequirementUsecase(mockRepo, timeout)

	fetchedDiscountRequirements := []domain.DiscountRequirement{
		{
			ID:                                bson.NewObjectID(),
			DiscountID:                        bson.NewObjectID(),
			DiscountRequirementRuleSystemName: "MinimumOrderTotalRequirement",
			ParentID:                          nil,
			InteractionTypeID:                 nil,
			IsGroup:                           false,
		},
		{
			ID:                                bson.NewObjectID(),
			DiscountID:                        bson.NewObjectID(),
			DiscountRequirementRuleSystemName: "CustomerRoleRequirement",
			ParentID:                          new(bson.ObjectID),
			InteractionTypeID:                 new(int),
			IsGroup:                           true,
		},
	}
	*fetchedDiscountRequirements[1].ParentID = bson.NewObjectID()
	*fetchedDiscountRequirements[1].InteractionTypeID = 1

	mockRepo.On("Fetch", mock.Anything).Return(fetchedDiscountRequirements, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedDiscountRequirements, result)
	mockRepo.AssertExpectations(t)
}
