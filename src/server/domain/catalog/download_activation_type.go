package domain

// DownloadActivationType represents a download activation type
type DownloadActivationType int

const (
	// WhenOrderIsPaid represents activation when the order is paid
	WhenOrderIsPaid DownloadActivationType = 0

	// Manually represents manual activation
	Manually DownloadActivationType = 10
)
