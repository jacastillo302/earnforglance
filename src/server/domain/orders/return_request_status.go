package domain

// ReturnRequestStatus represents a return status
type ReturnRequestStatus int

const (
	// Pending represents a pending return status
	ReturnPending ReturnRequestStatus = 0

	// Received represents a received return status
	Received ReturnRequestStatus = 10

	// ReturnAuthorized represents a return authorized status
	ReturnAuthorized ReturnRequestStatus = 20

	// ItemsRepaired represents an items repaired status
	ItemsRepaired ReturnRequestStatus = 30

	// ItemsRefunded represents an items refunded status
	ItemsRefunded ReturnRequestStatus = 40

	// RequestRejected represents a request rejected status
	RequestRejected ReturnRequestStatus = 50

	// Cancelled represents a cancelled return status
	ReturnCancelled ReturnRequestStatus = 60
)
