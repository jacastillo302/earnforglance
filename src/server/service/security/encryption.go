package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
)

// CreatePasswordHash creates a password hash using the specified algorithm
func CreatePasswordHash(password, saltkey, passwordFormat string) (string, error) {
	combined := password + saltkey

	switch strings.ToUpper(passwordFormat) {
	case "SHA256":
		hash := sha256.Sum256([]byte(combined))
		return hex.EncodeToString(hash[:]), nil
	case "SHA512":
		hash := sha512.Sum512([]byte(combined))
		return hex.EncodeToString(hash[:]), nil
	default:
		return "", errors.New("unsupported hash algorithm")
	}
}

// EncryptText encrypts plain text using the specified or default encryption key
func EncryptText(plainText string, encryptionPrivateKey string, encryptionKey string, useAesEncryptionAlgorithm bool) (string, error) {
	if plainText == "" {
		return plainText, nil
	}

	key := encryptionPrivateKey
	if key == "" {
		key = encryptionKey
	}
	if key == "" {
		return "", errors.New("encryption key is required")
	}

	block, iv, err := getEncryptionAlgorithm(key, useAesEncryptionAlgorithm)
	if err != nil {
		return "", err
	}

	encrypted, err := encryptTextToMemory([]byte(plainText), block, iv)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// getEncryptionAlgorithm returns cipher.Block and IV based on settings
func getEncryptionAlgorithm(encryptionKey string, useAes bool) (cipher.Block, []byte, error) {
	if encryptionKey == "" {
		return nil, nil, errors.New("encryptionKey cannot be empty")
	}

	if useAes {
		key := []byte(encryptionKey)
		if len(key) < 16 {
			padding := make([]byte, 16-len(key))
			key = append(key, padding...)
		}
		block, err := aes.NewCipher(key[:16])
		if err != nil {
			return nil, nil, err
		}
		iv := key[:block.BlockSize()]
		return block, iv, nil
	} else {
		key := []byte(encryptionKey)
		if len(key) < 24 {
			padding := make([]byte, 24-len(key))
			key = append(key, padding...)
		}
		block, err := des.NewTripleDESCipher(key[:24])
		if err != nil {
			return nil, nil, err
		}
		iv := key[len(key)-block.BlockSize():]
		return block, iv, nil
	}
}

// encryptTextToMemory encrypts data using the given cipher.Block and IV (CBC mode)
func encryptTextToMemory(data []byte, block cipher.Block, iv []byte) ([]byte, error) {
	blockSize := block.BlockSize()
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	data = append(data, padtext...)

	mode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(data))
	mode.CryptBlocks(encrypted, data)
	return encrypted, nil
}

// DecryptText decrypts the given base64-encoded cipher text using the specified or default encryption key
func DecryptText(cipherText string, encryptionPrivateKey string, encryptionKey string, useAesEncryptionAlgorithm bool) (string, error) {
	if cipherText == "" {
		return cipherText, nil
	}

	key := encryptionPrivateKey
	if key == "" {
		key = encryptionKey
	}
	if key == "" {
		return "", errors.New("encryption key is required")
	}

	block, iv, err := getEncryptionAlgorithm(key, useAesEncryptionAlgorithm)
	if err != nil {
		return "", err
	}

	buffer, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	plain, err := decryptTextFromMemory(buffer, block, iv)
	if err != nil {
		return "", err
	}

	return plain, nil
}

// decryptTextFromMemory decrypts data using the given cipher.Block and IV (CBC mode)
func decryptTextFromMemory(data []byte, block cipher.Block, iv []byte) (string, error) {
	blockSize := block.BlockSize()
	if len(data)%blockSize != 0 {
		return "", errors.New("invalid encrypted data length")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(data))
	mode.CryptBlocks(decrypted, data)

	// Remove PKCS#7 padding
	paddingLen := int(decrypted[len(decrypted)-1])
	if paddingLen > blockSize || paddingLen == 0 {
		return "", errors.New("invalid padding")
	}
	return string(decrypted[:len(decrypted)-paddingLen]), nil
}
