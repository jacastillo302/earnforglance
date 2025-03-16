package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionPdfSettings = "pdf_settings"
)

// PdfSettings represents PDF settings
type PdfSettings struct {
	ID                                 primitive.ObjectID `bson:"_id,omitempty"`
	LogoPictureID                      int                `bson:"logo_picture_id"`
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
