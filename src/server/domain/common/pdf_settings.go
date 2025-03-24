package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPdfSettings = "pdf_settings"
)

// PdfSettings represents PDF settings
type PdfSettings struct {
	ID                                 primitive.ObjectID `bson:"_id,omitempty"`
	LogoPictureID                      primitive.ObjectID `bson:"logo_picture_id"`
	LetterPageSizeEnabled              bool               `bson:"letter_page_size_enabled"`
	RenderOrderNotes                   bool               `bson:"render_order_notes"`
	DisablePdfInvoicesForPendingOrders bool               `bson:"disable_pdf_invoices_for_pending_orders"`
	LtrFontName                        string             `bson:"ltr_font_name"`
	RtlFontName                        string             `bson:"rtl_font_name"`
	InvoiceFooterTextColumn1           string             `bson:"invoice_footer_text_column1"`
	InvoiceFooterTextColumn2           string             `bson:"invoice_footer_text_column2"`
	BaseFontSize                       float64            `bson:"base_font_size"`
	ImageTargetSize                    int                `bson:"image_target_size"`
}

type PdfSettingsRepository interface {
	CreateMany(c context.Context, items []PdfSettings) error
	Create(c context.Context, PdfSettings *PdfSettings) error
	Update(c context.Context, PdfSettings *PdfSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PdfSettings, error)
	FetchByID(c context.Context, ID string) (PdfSettings, error)
}

type PdfSettingsUsecase interface {
	CreateMany(c context.Context, items []PdfSettings) error
	FetchByID(c context.Context, ID string) (PdfSettings, error)
	Create(c context.Context, PdfSettings *PdfSettings) error
	Update(c context.Context, PdfSettings *PdfSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PdfSettings, error)
}
