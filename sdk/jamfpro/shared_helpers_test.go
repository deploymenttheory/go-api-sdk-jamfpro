// shared_helpers_test.go
// these are used where a bool value is set to a pointer in scenarios where the value needs to send a valid false value
// and not an empty value.
package jamfpro

import (
	"strconv"
	"testing"
)

func TestBoolPtr(t *testing.T) {
	tests := []struct {
		name  string
		input bool
		want  bool
	}{
		{"True value", true, true},
		{"False value", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BoolPtr(tt.input)
			if result == nil {
				t.Fatalf("BoolPtr(%v) returned nil", tt.input)
			}
			if *result != tt.want {
				t.Errorf("BoolPtr(%v) = %v, want %v", tt.input, *result, tt.want)
			}
		})
	}
}

func TestTruePtr(t *testing.T) {
	result := TruePtr()
	if result == nil {
		t.Fatal("TruePtr() returned nil")
	}
	if *result != true {
		t.Errorf("TruePtr() = %v, want true", *result)
	}
}

func TestFalsePtr(t *testing.T) {
	result := FalsePtr()
	if result == nil {
		t.Fatal("FalsePtr() returned nil")
	}
	if *result != false {
		t.Errorf("FalsePtr() = %v, want false", *result)
	}
}

func TestBoolPtrUniqueness(t *testing.T) {
	ptr1 := BoolPtr(true)
	ptr2 := BoolPtr(true)
	if ptr1 == ptr2 {
		t.Error("BoolPtr returned the same pointer for different calls")
	}
}

func TestTruePtrFalsePtrDifference(t *testing.T) {
	truePtr := TruePtr()
	falsePtr := FalsePtr()
	if truePtr == falsePtr {
		t.Error("TruePtr and FalsePtr returned the same pointer")
	}
}

func TestConsistency(t *testing.T) {
	for i := 0; i < 100; i++ {
		trueResult := TruePtr()
		falseResult := FalsePtr()

		if *trueResult != true {
			t.Errorf("Iteration %d: TruePtr() = %v, want true", i, *trueResult)
		}
		if *falseResult != false {
			t.Errorf("Iteration %d: FalsePtr() = %v, want false", i, *falseResult)
		}
	}
}

func TestIncrementStringID(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		shouldPanic bool
	}{
		{"Positive integer", "5", "6", false},
		{"Zero", "0", "1", false},
		{"Negative integer", "-1", "0", false},
		{"Large number", "999999", "1000000", false},
		{"Maximum int32", "2147483647", "2147483648", false},
		{"Minimum int32", "-2147483648", "-2147483647", false},
		{"Empty string", "", "", true},
		{"Non-numeric string", "abc", "", true},
		{"Mixed string", "123abc", "", true},
		{"Decimal number", "5.5", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if tt.shouldPanic && r == nil {
					t.Errorf("IncrementStringID(%q) should have panicked", tt.input)
				} else if !tt.shouldPanic && r != nil {
					t.Errorf("IncrementStringID(%q) should not have panicked", tt.input)
				}
			}()

			result := IncrementStringID(tt.input)
			if !tt.shouldPanic {
				if result != tt.expected {
					t.Errorf("IncrementStringID(%q) = %q, want %q", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestIncrementStringIDEdgeCases(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Very large number", "9223372036854775808"},  // 2^63 (one more than max int64)
		{"Very small number", "-9223372036854775809"}, // -2^63 - 1 (one less than min int64)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("IncrementStringID(%q) should have panicked", tc.input)
				}
			}()
			IncrementStringID(tc.input)
		})
	}
}

func TestIncrementStringIDConsistency(t *testing.T) {
	id := "1"
	for i := 2; i <= 5; i++ {
		id = IncrementStringID(id)
		expected := strconv.Itoa(i)
		if id != expected {
			t.Errorf("Consecutive call %d: IncrementStringID returned %q, want %q", i, id, expected)
		}
	}
}
