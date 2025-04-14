package security

import (
	customers "earnforglance/server/domain/customers"
)

func PasswordsMatch(PasswordFormatID int, customerPassword string, enteredPassword string) bool {

	if customerPassword == "" || enteredPassword == "" {
		return false
	}

	savedPassword := ""
	switch PasswordFormatID {
	case int(customers.Clear): // Keep only one case for customers.Clear
		savedPassword = enteredPassword
	case int(customers.Encrypted):
		savedPassword = "" //EncryptText(enteredPassword, customerPassword.PasswordSalt,nil)
	case int(customers.Hashed):
		savedPassword = "" //crs.encryptionService.CreatePasswordHash(enteredPassword, customerPassword.PasswordSalt, crs.customerSettings.HashedPasswordFormat)
	}

	if customerPassword == "" {
		return false
	}

	return customerPassword == savedPassword
}
