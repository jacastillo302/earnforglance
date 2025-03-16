package domain

// GdprRequestType represents a GDPR request type
type GdprRequestType int

const (
	// ConsentAgree represents consent (agree)
	ConsentAgree GdprRequestType = 1

	// ConsentDisagree represents consent (disagree)
	ConsentDisagree GdprRequestType = 5

	// ExportData represents export data
	ExportData GdprRequestType = 10

	// DeleteCustomer represents delete customer
	DeleteCustomer GdprRequestType = 15

	// ProfileChanged represents user changed profile
	ProfileChanged GdprRequestType = 20
)
