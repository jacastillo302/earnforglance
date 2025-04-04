package common

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
)

// CreateHash creates a hash for the given data using the specified hash algorithm.
// If trimByteCount > 0, only the first trimByteCount bytes of the data will be used.
func CreateHash(data []byte, hashAlgorithm string, trimByteCount int) (string, error) {
	if hashAlgorithm == "" {
		return "", errors.New("hashAlgorithm cannot be null or empty")
	}

	// Select the hash algorithm
	var algorithm hash.Hash
	switch hashAlgorithm {
	case "MD5":
		algorithm = md5.New()
	case "SHA1":
		algorithm = sha1.New()
	case "SHA256":
		algorithm = sha256.New()
	case "SHA512":
		algorithm = sha512.New()
	default:
		return "", errors.New("unrecognized hash name")
	}

	// Trim the data if necessary
	if trimByteCount > 0 && len(data) > trimByteCount {
		data = data[:trimByteCount]
	}

	// Compute the hash
	algorithm.Write(data)
	hashBytes := algorithm.Sum(nil)

	// Convert the hash to a hexadecimal string
	return hex.EncodeToString(hashBytes), nil
}
