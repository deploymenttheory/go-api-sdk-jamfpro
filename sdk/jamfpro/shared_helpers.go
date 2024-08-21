// shared_helpers.go
// these are used where a bool value is set to a pointer in scenarios where the value needs to send a valid false value
// and not an empty value.
package jamfpro

import "strconv"

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

// Helper function
func IncrementID(currentID string) string {
	id, err := strconv.Atoi(currentID)
	if err != nil || id == -1 {
		return "1"
	}
	return strconv.Itoa(id + 1)
}
