package apikey

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"strings"
)

const (
	secretLength   = 32 // Length of the secret
	checksumLength = 4  // Length of checksum
)

// CalculateChecksum calculates a CRC32 checksum
func CalculateChecksum(data []byte) []byte {
	checksum := crc32.ChecksumIEEE(data)
	return []byte{
		byte(checksum >> 24),
		byte(checksum >> 16),
		byte(checksum >> 8),
		byte(checksum),
	}
}

// GenerateAPIKey generates an API key with a checksum
func GenerateAPIKey() string {
	identifier := "sqra"
	secret := make([]byte, secretLength)
	_, err := rand.Read(secret)
	if err != nil {
		panic(err)
	}

	checksum := CalculateChecksum(secret)
	encodedSecret := hex.EncodeToString(secret)
	encodedChecksum := hex.EncodeToString(checksum)

	return fmt.Sprintf("%s_%s_%s", identifier, encodedSecret, encodedChecksum)
}

// ValidateAPIKey validates the API key by checking the checksum
func ValidateAPIKey(apiKey string) bool {
	parts := strings.Split(apiKey, "_")
	if len(parts) != 3 || parts[0] != "sqra" {
		return false
	}

	secret, err := hex.DecodeString(parts[1])
	if err != nil {
		return false
	}

	if len(secret) != secretLength {
		return false
	}

	checksum, err := hex.DecodeString(parts[2])
	if err != nil || len(checksum) != checksumLength {
		return false
	}

	expectedChecksum := CalculateChecksum(secret)

	for i := range checksum {
		if checksum[i] != expectedChecksum[i] {
			return false
		}
	}

	return true
}
