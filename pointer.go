package ptrkit

// Pointer returns a pointer to the input value.
// Example usage:
//
//	var x int = 42
//	ptr := Pointer(x) // ptr points to x
func Pointer[K any](val K) *K {
	return &val
}
