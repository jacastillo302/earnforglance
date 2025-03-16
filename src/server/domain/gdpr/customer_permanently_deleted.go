package domain

const (
	CollectionCustomerPermanentlyDeleted = "customer_permanently_deletes"
)

// CustomerPermanentlyDeleted represents a customer permanently deleted (GDPR)
type CustomerPermanentlyDeleted struct {
	CustomerID int    `bson:"customer_id"`
	Email      string `bson:"email"`
}

// NewCustomerPermanentlyDeleted creates a new CustomerPermanentlyDeleted instance
func NewCustomerPermanentlyDeleted(customerID int, email string) *CustomerPermanentlyDeleted {
	return &CustomerPermanentlyDeleted{
		CustomerID: customerID,
		Email:      email,
	}
}
