package ptrkit

// Default returns the value of the pointer if it's not nil,
// otherwise it returns the provided default value.
// Example usage:
//
//	var x *int = nil
//	val := Default(x, 42) // val is 42
func Default[T any](val *T, defVal T) T {
	if val != nil {
		return PointerValue(val)
	}

	return defVal
}
