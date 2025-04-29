// shared_helpers.go
// these are used where a bool value is set to a pointer in scenarios where the value needs to send a valid false value
// and not an empty value.
package jamfpro

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"

	"golang.org/x/crypto/sha3"
)

// BoolPtr returns a pointer to a bool value
func BoolPtr(b bool) *bool {
	return &b
}

// TruePtr returns a pointer to a true bool value
func TruePtr() *bool {
	return BoolPtr(true)
}

// FalsePtr returns a pointer to a false bool value
func FalsePtr() *bool {
	return BoolPtr(false)
}

// IntPtr returns a pointer to an int value
func IntPtr(i int) *int {
	return &i
}

// StringPtr returns a pointer to a string value
func StringPtr(s string) *string {
	return &s
}

// IncrementStringID increments the given ID string.
// It returns the incremented ID as a string or panics if the input is not convertible to an integer.
func IncrementStringID(currentID string) string {
	id, err := strconv.Atoi(currentID)
	if err != nil {
		panic(fmt.Sprintf("Error converting ID to int: %v", err))
	}
	return strconv.Itoa(id + 1)
}

// CalculateSHA3_512 calculates the SHA3-512 hash of the supplied file in the path.
func CalculateSHA3_512(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file for SHA3-512 calculation: %v", err)
	}
	defer file.Close()

	hash := sha3.New512()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("failed to calculate SHA3-512: %v", err)
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
