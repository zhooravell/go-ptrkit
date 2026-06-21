package ptrkit

// PointerValue returns a value of the input pointer.
// Example usage:
//
//	var x int = 42
//	val := Pointer(&x) // 42
func PointerValue[K any](val *K) (zero K) {
	if val == nil {
		return zero
	}

	return *val
}
