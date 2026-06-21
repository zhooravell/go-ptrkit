package ptrkit

import (
	"fmt"
	"testing"
)

func TestIsPointer(t *testing.T) {
	var (
		intVal               = 42
		int8Val      int8    = 8
		int16Val     int16   = 16
		int32Val     int32   = 32
		int64Val     int64   = 64
		uintVal      uint    = 42
		uint8Val     uint8   = 8
		uint16Val    uint16  = 16
		uint32Val    uint32  = 32
		uint64Val    uint64  = 64
		float32Val   float32 = 3.14
		float64Val           = 3.1415
		boolVal              = true
		stringVal            = "test"
		arrayVal             = [3]int{1, 2, 3}
		sliceVal             = []int{1, 2, 3}
		mapVal               = map[string]int{"one": 1}
		chanVal              = make(chan int)
		structVal            = struct{ X int }{42}
		interfaceVal any     = "any value"
		funcVal              = func() {}
	)

	intPtr := &intVal
	int8Ptr := &int8Val
	int16Ptr := &int16Val
	int32Ptr := &int32Val
	int64Ptr := &int64Val
	uintPtr := &uintVal
	uint8Ptr := &uint8Val
	uint16Ptr := &uint16Val
	uint32Ptr := &uint32Val
	uint64Ptr := &uint64Val
	float32Ptr := &float32Val
	float64Ptr := &float64Val
	boolPtr := &boolVal
	stringPtr := &stringVal
	arrayPtr := &arrayVal
	slicePtr := &sliceVal
	mapPtr := &mapVal
	chanPtr := &chanVal
	structPtr := &structVal
	interfacePtr := &interfaceVal
	funcPtr := &funcVal

	ptrToPtr := &intPtr

	var nilIntPtr *int = nil
	var nilStructPtr *struct{} = nil

	tests := map[string]struct {
		val  any
		want bool
	}{
		"int":       {intVal, false},
		"int8":      {int8Val, false},
		"int16":     {int16Val, false},
		"int32":     {int32Val, false},
		"int64":     {int64Val, false},
		"uint":      {uintVal, false},
		"uint8":     {uint8Val, false},
		"uint16":    {uint16Val, false},
		"uint32":    {uint32Val, false},
		"uint64":    {uint64Val, false},
		"float32":   {float32Val, false},
		"float64":   {float64Val, false},
		"bool":      {boolVal, false},
		"string":    {stringVal, false},
		"array":     {arrayVal, false},
		"slice":     {sliceVal, false},
		"map":       {mapVal, false},
		"chan":      {chanVal, false},
		"struct":    {structVal, false},
		"interface": {interfaceVal, false},
		"func":      {funcVal, false},

		"*int":       {intPtr, true},
		"*int8":      {int8Ptr, true},
		"*int16":     {int16Ptr, true},
		"*int32":     {int32Ptr, true},
		"*int64":     {int64Ptr, true},
		"*uint":      {uintPtr, true},
		"*uint8":     {uint8Ptr, true},
		"*uint16":    {uint16Ptr, true},
		"*uint32":    {uint32Ptr, true},
		"*uint64":    {uint64Ptr, true},
		"*float32":   {float32Ptr, true},
		"*float64":   {float64Ptr, true},
		"*bool":      {boolPtr, true},
		"*string":    {stringPtr, true},
		"*array":     {arrayPtr, true},
		"*slice":     {slicePtr, true},
		"*map":       {mapPtr, true},
		"*chan":      {chanPtr, true},
		"*struct":    {structPtr, true},
		"*interface": {interfacePtr, true},
		"*func":      {funcPtr, true},

		"**int":         {ptrToPtr, true},
		"nil":           {nil, false},
		"nil *int":      {nilIntPtr, true},
		"nil *struct{}": {nilStructPtr, true},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := IsPointer(tt.val); got != tt.want {
				t.Errorf("IsPointer(%v) = %v, want %v", tt.val, got, tt.want)
			}
		})
	}
}

func ExampleIsPointer() {
	fmt.Println(IsPointer(123))
	fmt.Println(IsPointer(&struct{ X int }{42}))

	// Output:
	// false
	// true
}
