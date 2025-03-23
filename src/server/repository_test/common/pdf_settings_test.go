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

type MockSingleResultPdfSettings struct {
	mock.Mock
}

func (m *MockSingleResultPdfSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PdfSettings); ok {
		*v.(*domain.PdfSettings) = *result
	}
	return args.Error(1)
}

var mockItemPdfSettings = &domain.PdfSettings{
	ID:                                 primitive.NewObjectID(), // Existing ID of the record to update
	LogoPictureID:                      primitive.NewObjectID(),
	LetterPageSizeEnabled:              false,
	RenderOrderNotes:                   false,
	DisablePdfInvoicesForPendingOrders: true,
	LtrFontName:                        "Verdana",
	RtlFontName:                        "Courier New",
	InvoiceFooterTextColumn1:           "Updated footer text column 1",
	InvoiceFooterTextColumn2:           "Updated footer text column 2",
	BaseFontSize:                       10.0,
	ImageTargetSize:                    500,
}

func TestPdfSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPdfSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPdfSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPdfSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPdfSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPdfSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPdfSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPdfSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPdfSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPdfSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPdfSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPdfSettings).Return(nil, nil).Once()

	repo := repository.NewPdfSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPdfSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPdfSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPdfSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPdfSettings.ID}
	update := bson.M{"$set": mockItemPdfSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPdfSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPdfSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
