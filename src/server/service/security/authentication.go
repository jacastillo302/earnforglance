package security

import (
	domain "earnforglance/server/domain/customers"
	"net/http"
)

// CookieAuthenticationService handles cookie-based authentication
type CookieAuthenticationService struct {
	cachedCustomer *domain.Customer
}

// SignIn signs in a customer and sets authentication cookie
func (s *CookieAuthenticationService) SignIn(w http.ResponseWriter, customer *domain.Customer, isPersistent bool) error {

	if customer == nil {
		return nil
	}

	// Here you would serialize principal and authProps into a secure cookie
	// For demonstration, set a simple cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "EarnAuth",
		Value:    customer.Username, // In production, store a secure token or encrypted data
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})

	s.cachedCustomer = customer
	return nil
}

// SignOut signs out the customer
func (s *CookieAuthenticationService) SignOut(w http.ResponseWriter) {
	s.cachedCustomer = nil
	http.SetCookie(w, &http.Cookie{
		Name:     "EarnAuth",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	})

}
