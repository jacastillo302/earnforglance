package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultSearchTermReportLine struct {
	mock.Mock
}

func (m *MockSingleResultSearchTermReportLine) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.SearchTermReportLine); ok {
		*v.(*domain.SearchTermReportLine) = *result
	}
	return args.Error(1)
}

var mockItemSearchTermReportLine = &domain.SearchTermReportLine{
	ID:      primitive.NewObjectID(), // Existing ID of the record to update
	Keyword: "smartphone",
	Count:   200,
}

func TestSearchTermReportLineRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionSearchTermReportLine

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSearchTermReportLine{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemSearchTermReportLine, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSearchTermReportLineRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSearchTermReportLine.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSearchTermReportLine{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSearchTermReportLineRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSearchTermReportLine.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestSearchTermReportLineRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSearchTermReportLine

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemSearchTermReportLine).Return(nil, nil).Once()

	repo := repository.NewSearchTermReportLineRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemSearchTermReportLine)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestSearchTermReportLineRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSearchTermReportLine

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemSearchTermReportLine.ID}
	update := bson.M{"$set": mockItemSearchTermReportLine}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewSearchTermReportLineRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemSearchTermReportLine)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
