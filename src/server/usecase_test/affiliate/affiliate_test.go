package usecase_test

import (
	"context"
	domian "earnforglance/server/domain/affiliate"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	test "earnforglance/server/usecase/affiliate"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAffiliateUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.AffiliateRepository)
	time := time.Duration(10)
	usecase := test.NewAffiliateUsecase(mockRepo, time) // Assuming a constructor exists

	affiliateID := primitive.NewObjectID().Hex()
	expectedAffiliate := domian.Affiliate{
		ID:              primitive.NewObjectID(),
		AdminComment:    "Test Comment",
		FriendlyUrlName: "test-url",
		Deleted:         false,
		Active:          true,
	}

	mockRepo.On("FetchByID", mock.Anything, affiliateID).Return(expectedAffiliate, nil)

	result, err := usecase.FetchByID(context.Background(), affiliateID)

	assert.NoError(t, err)
	assert.Equal(t, expectedAffiliate, result)
	mockRepo.AssertExpectations(t)
}

func TestAffiliateUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.AffiliateRepository)
	time := time.Duration(10)
	usecase := test.NewAffiliateUsecase(mockRepo, time) // Assuming a constructor exists

	newAffiliate := &domian.Affiliate{
		AdminComment:    "New Comment",
		FriendlyUrlName: "new-url",
		Deleted:         false,
		Active:          true,
	}

	mockRepo.On("Create", mock.Anything, newAffiliate).Return(nil)

	err := usecase.Create(context.Background(), newAffiliate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAffiliateUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.AffiliateRepository)
	time := time.Duration(10)
	usecase := test.NewAffiliateUsecase(mockRepo, time) // Assuming a constructor exists

	updatedAffiliate := &domian.Affiliate{
		ID:              primitive.NewObjectID(),
		AdminComment:    "Updated Comment",
		FriendlyUrlName: "updated-url",
		Deleted:         false,
		Active:          true,
	}

	mockRepo.On("Update", mock.Anything, updatedAffiliate).Return(nil)

	err := usecase.Update(context.Background(), updatedAffiliate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAffiliateUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.AffiliateRepository)
	time := time.Duration(10)
	usecase := test.NewAffiliateUsecase(mockRepo, time) // Assuming a constructor exists

	affiliateID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, affiliateID).Return(nil)

	err := usecase.Delete(context.Background(), affiliateID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAffiliateUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.AffiliateRepository)
	time := time.Duration(10)
	usecase := test.NewAffiliateUsecase(mockRepo, time) // Assuming a constructor exists

	expectedAffiliates := []domian.Affiliate{
		{
			ID:              primitive.NewObjectID(),
			AdminComment:    "Comment 1",
			FriendlyUrlName: "url-1",
			Deleted:         false,
			Active:          true,
		},
		{
			ID:              primitive.NewObjectID(),
			AdminComment:    "Comment 2",
			FriendlyUrlName: "url-2",
			Deleted:         false,
			Active:          true,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedAffiliates, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedAffiliates, result)
	mockRepo.AssertExpectations(t)
}
