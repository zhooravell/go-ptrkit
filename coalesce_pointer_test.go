package ptrkit

import (
	"fmt"
	"testing"
)

func TestCoalescePointer_int(t *testing.T) {
	var (
		first  *int
		second = Pointer(42)
		third  *int
	)

	result := CoalescePointer(first, second, third)

	if result == nil || *result != 42 {
		t.Errorf("Expected pointer to 42, got %v", result)
	}
}

func TestCoalescePointer_int_all_nil(t *testing.T) {
	var (
		first  *int
		second *int
		third  *int
	)

	result := CoalescePointer(first, second, third)

	if result != nil {
		t.Errorf("Expected pointer to nil, got %v", *result)
	}
}

func TestCoalescePointer_string(t *testing.T) {
	var (
		first  *string
		second = Pointer("test")
		third  *string
	)

	result := CoalescePointer(first, second, third)

	if result == nil || *result != "test" {
		t.Errorf("Expected pointer to 42, got %v", result)
	}
}

func TestCoalescePointer_string_all_nil(t *testing.T) {
	var (
		first  *string
		second *string
		third  *string
	)

	result := CoalescePointer(first, second, third)

	if result != nil {
		t.Errorf("Expected pointer to nil, got %v", *result)
	}
}

func ExampleCoalescePointer() {
	var (
		first  *string
		second = Pointer("test")
		third  *string
	)

	result := CoalescePointer(first, second, third)

	if result != nil {
		fmt.Println(*result)
	}

	// Output: test
}
