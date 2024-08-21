// shared_helpers_test.go
// these are used where a bool value is set to a pointer in scenarios where the value needs to send a valid false value
// and not an empty value.
package jamfpro

import (
	"strconv"
	"strings"
	"testing"
)

func TestIncrementStringID(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Positive integer", "5", "6"},
		{"Zero", "0", "1"},
		{"Negative integer", "-1", "0"},
		{"Large number", "999999", "1000000"},
		{"Maximum int32", "2147483647", "2147483648"},
		{"Minimum int32", "-2147483648", "-2147483647"},
		{"Empty string", "", "Error converting ID to int:"},
		{"Non-numeric string", "abc", "Error converting ID to int:"},
		{"Mixed string", "123abc", "Error converting ID to int:"},
		{"Decimal number", "5.5", "Error converting ID to int:"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IncrementStringID(tt.input)
			if tt.expected == "Error converting ID to int:" {
				if !strings.HasPrefix(result, tt.expected) {
					t.Errorf("IncrementStringID(%q) = %q, want error message starting with %q", tt.input, result, tt.expected)
				}
			} else if result != tt.expected {
				t.Errorf("IncrementStringID(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIncrementStringIDEdgeCases(t *testing.T) {
	// Test for very large numbers (beyond int64)
	veryLargeNumber := "9223372036854775808" // 2^63 (one more than max int64)
	result := IncrementStringID(veryLargeNumber)
	if !strings.HasPrefix(result, "Error converting ID to int:") {
		t.Errorf("IncrementStringID(%q) should return an error for numbers larger than int64", veryLargeNumber)
	}

	// Test for very small numbers (beyond negative int64)
	verySmallNumber := "-9223372036854775809" // -2^63 - 1 (one less than min int64)
	result = IncrementStringID(verySmallNumber)
	if !strings.HasPrefix(result, "Error converting ID to int:") {
		t.Errorf("IncrementStringID(%q) should return an error for numbers smaller than negative int64", verySmallNumber)
	}
}

func TestIncrementStringIDConsistency(t *testing.T) {
	// Test that consecutive calls produce expected results
	id := "1"
	for i := 2; i <= 5; i++ {
		id = IncrementStringID(id)
		expected := strconv.Itoa(i)
		if id != expected {
			t.Errorf("Consecutive call %d: IncrementStringID returned %q, want %q", i, id, expected)
		}
	}
}
