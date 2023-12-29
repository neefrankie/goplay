package format

import (
	"fmt"
	"reflect"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		// The Len method returns the number of elements
		// of a slice or array value.
		// Index(i) retrieves the element at index i, also
		// as a reflect.Value.
		for i := 0; i < v.Len(); i++ {
			// Recursively invokes itself on each
			// element of the sequence, appending the subscript
			// notation "[i]" to the path.
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		// The NumField method reports the number of fields
		// in the struct.
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			// Field(i) returns the value of the i-th field
			// as a reflect.Value. The list of fields includes
			// ones promoted from anonymous fields.
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		// The MapKeys method returns a slice of reflect.Value,
		// one per map key. As usual when iterating a map,
		// the order is undefined.
		for _, key := range v.MapKeys() {
			// MapIndex(key) returns the value corresponding to
			// key.
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			// The Elem method returns the variable pointed by
			// a pointer, again as a reflect.Value.
			// This operation would be safe even if the pointer
			// value is nil, in which case the result would
			// have kind Invalid.
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
