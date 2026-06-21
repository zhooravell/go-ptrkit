package ptrkit

import "reflect"

// PointerDepth returns the depth of pointer indirection for the given value.
// It returns 0 if the value is not a pointer, 1 for a single pointer, 2 for a pointer to a pointer, and so on.
// Example usage:
//
//	var x int = 42
//	PointerDepth(x)        // 0
//	PointerDepth(&x)       // 1
//	PointerDepth(&(&x))    // 2
//	PointerDepth(nil)      // 0
func PointerDepth(v any) int {
	if v == nil {
		return 0
	}

	t := reflect.TypeOf(v)
	depth := 0

	for t.Kind() == reflect.Ptr {
		depth++
		t = t.Elem()
	}

	return depth
}
