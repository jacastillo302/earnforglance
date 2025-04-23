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

func TestSearchTermReportLineUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.SearchTermReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewSearchTermReportLineUsecase(mockRepo, timeout)

	commonID := bson.NewObjectID().Hex()

	updatedSearchTermReportLine := domain.SearchTermReportLine{
		ID:      bson.NewObjectID(), // Existing ID of the record to update
		Keyword: "smartphone",
		Count:   200,
	}

	mockRepo.On("FetchByID", mock.Anything, commonID).Return(updatedSearchTermReportLine, nil)

	result, err := usecase.FetchByID(context.Background(), commonID)

	assert.NoError(t, err)
	assert.Equal(t, updatedSearchTermReportLine, result)
	mockRepo.AssertExpectations(t)
}

func TestSearchTermReportLineUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.SearchTermReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewSearchTermReportLineUsecase(mockRepo, timeout)

	newSearchTermReportLine := &domain.SearchTermReportLine{
		Keyword: "laptop",
		Count:   150,
	}

	mockRepo.On("Create", mock.Anything, newSearchTermReportLine).Return(nil)

	err := usecase.Create(context.Background(), newSearchTermReportLine)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSearchTermReportLineUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.SearchTermReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewSearchTermReportLineUsecase(mockRepo, timeout)

	updatedSearchTermReportLine := &domain.SearchTermReportLine{
		ID:      bson.NewObjectID(), // Existing ID of the record to update
		Keyword: "smartphone",
		Count:   200,
	}

	mockRepo.On("Update", mock.Anything, updatedSearchTermReportLine).Return(nil)

	err := usecase.Update(context.Background(), updatedSearchTermReportLine)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSearchTermReportLineUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.SearchTermReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewSearchTermReportLineUsecase(mockRepo, timeout)

	commonID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, commonID).Return(nil)

	err := usecase.Delete(context.Background(), commonID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSearchTermReportLineUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.SearchTermReportLineRepository)
	timeout := time.Duration(10)
	usecase := test.NewSearchTermReportLineUsecase(mockRepo, timeout)

	fetchedSearchTermReportLines := []domain.SearchTermReportLine{
		{
			ID:      bson.NewObjectID(),
			Keyword: "laptop",
			Count:   150,
		},
		{
			ID:      bson.NewObjectID(),
			Keyword: "smartphone",
			Count:   200,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedSearchTermReportLines, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedSearchTermReportLines, result)
	mockRepo.AssertExpectations(t)
}
