package ptrkit

import "reflect"

// IsPointer checks if the given value is a pointer.
// It returns true if the value is a pointer, and false otherwise.
// Example usage:
//
//	var x int = 42
//	IsPointer(&x) // true
//	IsPointer(x)  // false
func IsPointer(v any) bool {
	if v == nil {
		return false
	}

	return reflect.TypeOf(v).Kind() == reflect.Ptr
}
