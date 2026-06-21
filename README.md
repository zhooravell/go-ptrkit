# GoLang pointer toolkit
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)
[![GO CI](https://github.com/zhooravell/go-ptrkit/actions/workflows/go.yml/badge.svg)](https://github.com/zhooravell/go-ptrkit/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/zhooravell/go-ptrkit)](https://goreportcard.com/report/github.com/zhooravell/go-ptrkit)

> A simple and lightweight Go package with helper functions for working with pointers.  
> Useful when you need to quickly create or extract values from pointers in a clean way.

![cover](cover.webp)

## ✨ Features

- `Pointer` — wraps a value into a pointer.
- `PointerValue` — safely extracts a value from a pointer (returns zero value if `nil`).
- `IsPointer` — checks if the given value is a pointer.
- `Default` — returns the value of the pointer if it's not nil.
- `CoalescePointer` — returns the first non-nil pointer from the provided list of pointers.
- `PointerDepth` — returns the depth of pointer indirection for the given value.

## Examples

### Pointer

Wraps any value into a pointer. Useful when an API requires a pointer but you only have a literal or local variable.

```go
ptr := ptrkit.Pointer(42)
fmt.Println(*ptr) // 42

ptr2 := ptrkit.Pointer("hello")
fmt.Println(*ptr2) // hello
```

### PointerValue

Safely dereferences a pointer. Returns the zero value of the type when the pointer is `nil`, avoiding a nil-dereference panic.

```go
x := 42
fmt.Println(ptrkit.PointerValue(&x)) // 42
fmt.Println(ptrkit.PointerValue[int](nil)) // 0
```

### IsPointer

Reports whether the given value is a pointer. Works with any type at runtime via reflection.

```go
x := 42
fmt.Println(ptrkit.IsPointer(&x)) // true
fmt.Println(ptrkit.IsPointer(x))  // false
fmt.Println(ptrkit.IsPointer(nil)) // false
```

### Default

Returns the dereferenced pointer value when non-nil, or falls back to the provided default. A concise alternative to an inline nil-check.

```go
var x *int = nil
fmt.Println(ptrkit.Default(x, 42)) // 42

y := 7
fmt.Println(ptrkit.Default(&y, 42)) // 7
```

### CoalescePointer

Returns the first non-nil pointer from a variadic list, similar to SQL `COALESCE`. Handy for merging optional config values with fallbacks.

```go
var a *int = nil
b := ptrkit.Pointer(42)
c := ptrkit.Pointer(100)

ptr := ptrkit.CoalescePointer(a, b, c)
fmt.Println(*ptr) // 42
```

### PointerDepth

Returns the number of pointer indirections for a value. Useful for reflection-based utilities that need to unwrap nested pointers.

```go
x := 42
pp := &x
fmt.Println(ptrkit.PointerDepth(x))   // 0
fmt.Println(ptrkit.PointerDepth(&x))  // 1
fmt.Println(ptrkit.PointerDepth(&pp)) // 2
fmt.Println(ptrkit.PointerDepth(nil)) // 0
```

## Resources

- [Go Specification: Pointer types](https://go.dev/ref/spec#Pointer_types)
- [Effective Go: Pointers vs. Values](https://go.dev/doc/effective_go#pointers_vs_values)
- [Go Tour: Pointers](https://go.dev/tour/moretypes/1)
- [Go Specification: Type parameters](https://go.dev/ref/spec#Type_parameter_declarations)
- [An Introduction to Generics](https://go.dev/blog/intro-generics)
- [Go Tour: Generics](https://go.dev/tour/generics/1)