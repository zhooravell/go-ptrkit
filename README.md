# GoLang pointer toolkit
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)
[![GO CI](https://github.com/zhooravell/go-ptrkit/actions/workflows/go.yml/badge.svg)](https://github.com/zhooravell/go-ptrkit/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/zhooravell/go-ptrkit)](https://goreportcard.com/report/github.com/zhooravell/go-ptrkit)

> A simple and lightweight Go package with helper functions for working with pointers.  
> Useful when you need to quickly create or extract values from pointers in a clean way.

## ✨ Features

- `Pointer` — wraps a value into a pointer.
- `PointerValue` — safely extracts a value from a pointer (returns zero value if `nil`).
- `IsPointer` — checks if the given value is a pointer.
- `Default` — returns the value of the pointer if it's not nil.
- `CoalescePointer` — returns the first non-nil pointer from the provided list of pointers.
- `PointerDepth` — returns the depth of pointer indirection for the given value.