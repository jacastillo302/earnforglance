package security

import (
	customers "earnforglance/server/domain/customers"
	service "earnforglance/server/service/customers"
	"time"
)

func ValidateCustomer(customer customers.Customer) string {

	sMessage := "Successful"
	if customer.ID.Hex() == "" {
		sMessage = "Account.Login.WrongCredentials.CustomerNotExist"
	}

	if customer.Deleted {
		sMessage = "Account.Login.WrongCredentials.Deleted"
	}

	if !customer.Active {
		sMessage = "Account.Login.WrongCredentials.NotActive"
	}

	if customer.CannotLoginUntilDateUtc != nil {
		if customer.CannotLoginUntilDateUtc.After(time.Now()) {
			sMessage = "Account.Login.WrongCredentials.LockedOut"
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
	return customerPassword == savedPassword
}
