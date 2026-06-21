package ptrkit

import (
	"reflect"
	"testing"
)

func TestPointerDepth(t *testing.T) {
	intVal := 42
	float64Val := 3.14
	stringVal := "hello"
	boolVal := true
	arrayVal := [3]int{1, 2, 3}
	sliceVal := []string{"a", "b", "c"}
	mapVal := map[string]int{"one": 1}
	structVal := struct{ Name string }{"Test"}
	chanVal := make(chan int)
	var interfaceVal any = "test"
	funcVal := func() {}

	intPtr := &intVal
	float64Ptr := &float64Val
	stringPtr := &stringVal
	boolPtr := &boolVal
	arrayPtr := &arrayVal
	slicePtr := &sliceVal
	mapPtr := &mapVal
	structPtr := &structVal
	chanPtr := &chanVal
	interfacePtr := &interfaceVal
	funcPtr := &funcVal

	intPtrPtr := &intPtr
	float64PtrPtr := &float64Ptr
	stringPtrPtr := &stringPtr
	boolPtrPtr := &boolPtr
	arrayPtrPtr := &arrayPtr
	slicePtrPtr := &slicePtr
	mapPtrPtr := &mapPtr
	structPtrPtr := &structPtr
	chanPtrPtr := &chanPtr
	interfacePtrPtr := &interfacePtr
	funcPtrPtr := &funcPtr

	intPtrPtrPtr := &intPtrPtr
	float64PtrPtrPtr := &float64PtrPtr
	stringPtrPtrPtr := &stringPtrPtr

	var nilIntPtr *int = nil
	var nilPtrPtr **int = nil

	tests := map[string]struct {
		value    any
		expected int
	}{
		"nil":       {nil, 0},
		"int":       {intVal, 0},
		"float64":   {float64Val, 0},
		"string":    {stringVal, 0},
		"bool":      {boolVal, 0},
		"array":     {arrayVal, 0},
		"slice":     {sliceVal, 0},
		"map":       {mapVal, 0},
		"struct":    {structVal, 0},
		"chan":      {chanVal, 0},
		"interface": {interfaceVal, 0},
		"func":      {funcVal, 0},

		"*int":       {intPtr, 1},
		"*float64":   {float64Ptr, 1},
		"*string":    {stringPtr, 1},
		"*bool":      {boolPtr, 1},
		"*array":     {arrayPtr, 1},
		"*slice":     {slicePtr, 1},
		"*map":       {mapPtr, 1},
		"*struct":    {structPtr, 1},
		"*chan":      {chanPtr, 1},
		"*interface": {interfacePtr, 1},
		"*func":      {funcPtr, 1},
		"nil *int":   {nilIntPtr, 1},

		"**int":       {intPtrPtr, 2},
		"**float64":   {float64PtrPtr, 2},
		"**string":    {stringPtrPtr, 2},
		"**bool":      {boolPtrPtr, 2},
		"**array":     {arrayPtrPtr, 2},
		"**slice":     {slicePtrPtr, 2},
		"**map":       {mapPtrPtr, 2},
		"**struct":    {structPtrPtr, 2},
		"**chan":      {chanPtrPtr, 2},
		"**interface": {interfacePtrPtr, 2},
		"**func":      {funcPtrPtr, 2},
		"nil **int":   {nilPtrPtr, 2},

		"***int":     {intPtrPtrPtr, 3},
		"***float64": {float64PtrPtrPtr, 3},
		"***string":  {stringPtrPtrPtr, 3},

		"empty interface{}":            {interface{}(nil), 0},
		"empty interface{} with value": {interface{}(10), 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := PointerDepth(tc.value)
			if got != tc.expected {
				t.Errorf("PointerDepth(%v) = %v, want %v", reflect.TypeOf(tc.value), got, tc.expected)
			}
		})
	}
}
