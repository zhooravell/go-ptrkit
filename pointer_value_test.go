package ptrkit

import (
	"errors"
	"fmt"
	"testing"
)

func TestPointerValue_nil_array(t *testing.T) {
	var (
		val      *[3]int
		expected = [3]int{} // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_any(t *testing.T) {
	var (
		val      *any
		expected any // zero value

		result = PointerValue(val)
	)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_int(t *testing.T) {
	var (
		val      *int
		expected int // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_int8(t *testing.T) {
	var (
		val      *int8
		expected int8 // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_int16(t *testing.T) {
	var (
		val      *int16
		expected int16 // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_int32(t *testing.T) {
	var (
		val      *int32
		expected int32 // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_int64(t *testing.T) {
	var (
		val      *int64
		expected int64 // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_uint(t *testing.T) {
	var (
		val      *uint
		expected uint // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_uint8(t *testing.T) {
	var (
		val      *uint8
		expected uint8 // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_uint16(t *testing.T) {
	var (
		val      *uint16
		expected uint16 // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_uint32(t *testing.T) {
	var (
		val      *uint32
		expected uint32 // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_uint64(t *testing.T) {
	var (
		val      *uint64
		expected uint64 // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_float32(t *testing.T) {
	var (
		val      *float32
		expected float32 // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_float64(t *testing.T) {
	var (
		val      *float64
		expected float64 // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_complex64(t *testing.T) {
	var (
		val      *complex64
		expected complex64 // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_complex128(t *testing.T) {
	var (
		val      *complex128
		expected complex128 // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_uintptr(t *testing.T) {
	var (
		val      *uintptr
		expected uintptr // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_bool(t *testing.T) {
	var (
		val      *bool
		expected bool // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_string(t *testing.T) {
	var (
		val      *string
		expected string // zero value

		result = PointerValue(val)
	)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_error(t *testing.T) {
	var (
		val      *error
		expected error // zero value

		result = PointerValue(val)
	)

	if !errors.Is(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_nil_custom_struct(t *testing.T) {
	type TestStruct struct{}

	var (
		val      *TestStruct
		expected TestStruct // zero value

		result = PointerValue(val)
	)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPointerValue_non_nil_int(t *testing.T) {
	var (
		val = 42

		result = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_int8(t *testing.T) {
	var (
		val    int8 = 1
		result      = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_int16(t *testing.T) {
	var (
		val    int16 = 1
		result       = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_int32(t *testing.T) {
	var (
		val    int32 = 1
		result       = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_int64(t *testing.T) {
	var (
		val    int64 = 1
		result       = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_uint(t *testing.T) {
	var (
		val    uint = 1
		result      = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_uint8(t *testing.T) {
	var (
		val    uint8 = 1
		result       = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_uint16(t *testing.T) {
	var (
		val    uint16 = 1
		result        = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_uint32(t *testing.T) {
	var (
		val    uint32 = 1
		result        = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_uint64(t *testing.T) {
	var (
		val    uint64 = 1
		result        = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_float32(t *testing.T) {
	var (
		val    float32 = -1.5
		result         = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_float64(t *testing.T) {
	var (
		val    = 1.5
		result = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_complex64(t *testing.T) {
	var (
		val    complex64 = 12345
		result           = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_complex128(t *testing.T) {
	var (
		val    complex128 = 12345
		result            = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_uintptr(t *testing.T) {
	var (
		val    uintptr = 12345
		result         = PointerValue(&val)
	)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_bool(t *testing.T) {
	val := true
	result := PointerValue(&val)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_string(t *testing.T) {
	val := "test"
	result := PointerValue(&val)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_error(t *testing.T) {
	val := errors.New("test error")
	result := PointerValue(&val)

	if !errors.Is(result, val) {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func TestPointerValue_non_nil_custom_struct(t *testing.T) {
	type TestStruct struct{}

	val := TestStruct{}
	result := PointerValue(&val)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}
}

func ExamplePointerValue() {
	val := 42
	ptr := &val

	fmt.Println(PointerValue(ptr))

	// Output: 42
}
