package ptrkit

import (
	"fmt"
	"testing"
)

func TestDefault_nil_int(t *testing.T) {
	var (
		value        *int
		defaultValue = 42

		result = Default(value, defaultValue)
	)

	if result != defaultValue {
		t.Errorf("Expected %v, got %v", defaultValue, result)
	}
}

func TestDefault_non_nil_int(t *testing.T) {
	value := 1
	result := Default(&value, 42)

	if result != value {
		t.Errorf("Expected %v, got %v", value, result)
	}
}

func TestDefault_nil_string(t *testing.T) {
	var (
		value        *string
		defaultValue = "42"

		result = Default(value, defaultValue)
	)

	if result != defaultValue {
		t.Errorf("Expected %v, got %v", defaultValue, result)
	}
}

func TestDefault_non_nil_string(t *testing.T) {
	value := "1"
	result := Default(&value, "42")

	if result != value {
		t.Errorf("Expected %v, got %v", value, result)
	}
}

func TestDefault_nil_bool(t *testing.T) {
	var (
		value        *bool
		defaultValue = true

		result = Default(value, defaultValue)
	)

	if result != defaultValue {
		t.Errorf("Expected %v, got %v", defaultValue, result)
	}
}

func TestDefault_non_nil_bool(t *testing.T) {
	value := true
	result := Default(&value, false)

	if result != value {
		t.Errorf("Expected %v, got %v", value, result)
	}
}

func TestDefault_nil_float64(t *testing.T) {
	var (
		value        *float64
		defaultValue = 42.2

		result = Default(value, defaultValue)
	)

	if result != defaultValue {
		t.Errorf("Expected %v, got %v", defaultValue, result)
	}
}

func TestDefault_non_nil_float64(t *testing.T) {
	value := 2.2
	result := Default(&value, 42.2)

	if result != value {
		t.Errorf("Expected %v, got %v", value, result)
	}
}

func ExampleDefault() {
	var value *string
	defaultValue := "default"

	result := Default(value, defaultValue)
	fmt.Println(result)

	// Output: default
}
