package ptrkit

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPointer(t *testing.T) {
	tests := map[string]any{
		"int":     123,
		"int8":    int8(42),
		"int16":   int16(1),
		"int32":   int32(33),
		"int64":   int64(123),
		"uint":    uint(123),
		"uint8":   uint8(42),
		"uint16":  uint16(1),
		"uint32":  uint32(33),
		"uint64":  uint64(123),
		"uintptr": uintptr(0xdeadbeef),

		"float32": float32(1.23),
		"float64": 1.23,

		"complex64":  complex64(1 + 2i),
		"complex128": complex(1, 2),

		"bool":       true,
		"bool_false": false,

		"string": "hello",
		"byte":   byte('a'),
		"rune":   rune('Я'),

		"array":     [3]int{1, 2, 3},
		"slice":     []string{"a", "b", "c"},
		"map":       map[string]int{"one": 1, "two": 2},
		"chan":      make(chan int),
		"struct":    struct{ X, Y int }{1, 2},
		"interface": any("test"),
		"func":      func() {},
		"empty_str": "",
		"nil_slice": []int(nil),
		"nil_map":   map[string]int(nil),
		"nil_chan":  (chan int)(nil),
		"nil_func":  (func())(nil),
	}

	for name, val := range tests {
		t.Run(name, func(t *testing.T) {
			got := Pointer(val)

			if reflect.ValueOf(got).Kind() != reflect.Ptr {
				t.Error("Pointer() should return a pointer")
			}

			if val == nil {
				if got != nil && !reflect.ValueOf(got).IsNil() {
					t.Errorf("Pointer() = %v, want pointer to nil", got)
				}
				return
			}

			if name == "func" || name == "nil_func" {
				if got == nil {
					t.Error("Pointer() returned nil for function")
				}
				return
			}

			gotValue := reflect.ValueOf(got).Elem().Interface()
			if !reflect.DeepEqual(val, gotValue) {
				t.Errorf("Pointer() = %v, want %v", gotValue, val)
			}
		})
	}
}

func ExamplePointer() {
	fmt.Println(*Pointer(123))

	// Output: 123
}
