package security

import (
	customers "earnforglance/server/domain/customers"
	service "earnforglance/server/service/customers"
	"fmt"
	"time"
)

func ValidateCustomer(customer customers.Customer) string {

	sMessage := "Successful"
	if customer.ID.Hex() == "" {
		sMessage = "The customer not exist"
	}

	if customer.Deleted {
		sMessage = "The customer account has been deleted"
	}

	if !customer.Active {
		sMessage = "The account has not been activated"
	}

	if customer.CannotLoginUntilDateUtc != nil {
		if customer.CannotLoginUntilDateUtc.After(time.Now()) {
			sMessage = "the customer account is locked out"
		}
	}

	return sMessage
}

func PasswordsMatch(PasswordFormatID int, customerPassword string, passwordSalt string, enteredPassword string, encryptionKey string, useAesEncryptionAlgorithm bool, hashedPasswordFormat string) bool {

	if customerPassword == "" || enteredPassword == "" {
		return false
	}

	savedPassword := ""
	switch PasswordFormatID {
	case int(customers.Clear): // Keep only one case for customers.Clear
		savedPassword = enteredPassword
	case int(customers.Encrypted):

		sPassword, err := EncryptText(enteredPassword, passwordSalt, encryptionKey, useAesEncryptionAlgorithm)
		if err != nil {
			return false
		}
		savedPassword = sPassword
	case int(customers.Hashed):
		hashAlgorithm := service.CustomerServicesDefaults{}
		hash, err := CreatePasswordHash(enteredPassword, passwordSalt, hashAlgorithm.DefaultHashedPasswordFormat())

		if err != nil {
			return false
		}

		savedPassword = hash
	}

	if customerPassword == "" {
		return false
	}

	fmt.Println("enteredPassword", enteredPassword)
	fmt.Println("savedPassword", savedPassword)
	fmt.Println("customerPassword", customerPassword)

	return customerPassword == savedPassword
}
