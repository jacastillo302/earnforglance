package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/common"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/common"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestSearchTermUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.SearchTermRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSearchTermUsecase(mockRepo, timeout)

	searchTermID := bson.NewObjectID().Hex()

	updatedSearchTerm := domain.SearchTerm{
		ID:      bson.NewObjectID(), // Existing ID of the record to update
		Keyword: "smartphone",
		StoreID: bson.NewObjectID(),
		Count:   200,
	}

	mockRepo.On("FetchByID", mock.Anything, searchTermID).Return(updatedSearchTerm, nil)

	result, err := usecase.FetchByID(context.Background(), searchTermID)

	assert.NoError(t, err)
	assert.Equal(t, updatedSearchTerm, result)
	mockRepo.AssertExpectations(t)
}

func TestSearchTermUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.SearchTermRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSearchTermUsecase(mockRepo, timeout)

	newSearchTerm := &domain.SearchTerm{
		Keyword: "laptop",
		StoreID: bson.NewObjectID(),
		Count:   150,
	}

	mockRepo.On("Create", mock.Anything, newSearchTerm).Return(nil)

	err := usecase.Create(context.Background(), newSearchTerm)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSearchTermUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.SearchTermRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSearchTermUsecase(mockRepo, timeout)

	updatedSearchTerm := &domain.SearchTerm{
		ID:      bson.NewObjectID(), // Existing ID of the record to update
		Keyword: "smartphone",
		StoreID: bson.NewObjectID(),
		Count:   200,
	}

	mockRepo.On("Update", mock.Anything, updatedSearchTerm).Return(nil)

	err := usecase.Update(context.Background(), updatedSearchTerm)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSearchTermUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.SearchTermRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSearchTermUsecase(mockRepo, timeout)

	searchTermID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, searchTermID).Return(nil)

	err := usecase.Delete(context.Background(), searchTermID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSearchTermUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.SearchTermRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSearchTermUsecase(mockRepo, timeout)

	fetchedSearchTerms := []domain.SearchTerm{
		{
			ID:      bson.NewObjectID(),
			Keyword: "laptop",
			StoreID: bson.NewObjectID(),
			Count:   150,
		},
		{
			ID:      bson.NewObjectID(),
			Keyword: "smartphone",
			StoreID: bson.NewObjectID(),
			Count:   200,
		},
	}
	mockRepo.On("Fetch", mock.Anything).Return(fetchedSearchTerms, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedSearchTerms, result)
	mockRepo.AssertExpectations(t)
}
