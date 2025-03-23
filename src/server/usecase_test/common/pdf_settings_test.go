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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPdfSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PdfSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewPdfSettingsUsecase(mockRepo, timeout)

	pdfsettingsID := primitive.NewObjectID().Hex()

	updatedPdfSettings := domain.PdfSettings{
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

	mockRepo.On("FetchByID", mock.Anything, pdfsettingsID).Return(updatedPdfSettings, nil)

	result, err := usecase.FetchByID(context.Background(), pdfsettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPdfSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestPdfSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PdfSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewPdfSettingsUsecase(mockRepo, timeout)

	newPdfSettings := &domain.PdfSettings{
		LogoPictureID:                      primitive.NewObjectID(),
		LetterPageSizeEnabled:              true,
		RenderOrderNotes:                   true,
		DisablePdfInvoicesForPendingOrders: false,
		LtrFontName:                        "Arial",
		RtlFontName:                        "Tahoma",
		InvoiceFooterTextColumn1:           "Thank you for your business!",
		InvoiceFooterTextColumn2:           "Contact us at support@example.com",
		BaseFontSize:                       12.0,
		ImageTargetSize:                    300,
	}

	mockRepo.On("Create", mock.Anything, newPdfSettings).Return(nil)

	err := usecase.Create(context.Background(), newPdfSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPdfSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PdfSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewPdfSettingsUsecase(mockRepo, timeout)

	updatedPdfSettings := &domain.PdfSettings{
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

	mockRepo.On("Update", mock.Anything, updatedPdfSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedPdfSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPdfSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PdfSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewPdfSettingsUsecase(mockRepo, timeout)

	pdfsettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, pdfsettingsID).Return(nil)

	err := usecase.Delete(context.Background(), pdfsettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPdfSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PdfSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewPdfSettingsUsecase(mockRepo, timeout)

	fetchedPdfSettings := []domain.PdfSettings{
		{
			ID:                                 primitive.NewObjectID(),
			LogoPictureID:                      primitive.NewObjectID(),
			LetterPageSizeEnabled:              true,
			RenderOrderNotes:                   true,
			DisablePdfInvoicesForPendingOrders: false,
			LtrFontName:                        "Arial",
			RtlFontName:                        "Tahoma",
			InvoiceFooterTextColumn1:           "Thank you for your business!",
			InvoiceFooterTextColumn2:           "Contact us at support@example.com",
			BaseFontSize:                       12.0,
			ImageTargetSize:                    300,
		},
		{
			ID:                                 primitive.NewObjectID(),
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
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPdfSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPdfSettings, result)
	mockRepo.AssertExpectations(t)
}
