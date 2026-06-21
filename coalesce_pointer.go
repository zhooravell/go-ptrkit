package ptrkit

// CoalescePointer returns the first non-nil pointer from the provided list of pointers.
// If all pointers are nil, it returns nil.
// Example usage:
//
//	var a *int = nil
//	var b *int = Pointer(42)
//	var c *int = Pointer(100)
//	ptr := CoalescePointer(a, b, c) // ptr points to 42
func CoalescePointer[T any](ptrs ...*T) *T {
	for _, ptr := range ptrs {
		if ptr != nil {
			return ptr
		}
	}

	return nil
}
