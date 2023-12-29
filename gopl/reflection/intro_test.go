package ch12

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

// reflect.Type and reflect.Value
// The `reflect` package defines two important types:
// * Type: represents a Go type.
// * Value
//
// `reflect.TypeOf` function accepts any interface{} and
// returns its dynamic type as a reflect.Type
func TestReflectType(t *testing.T) {
	// Assigns the value 3 to the interface{} parameter.
	// Assignment from a concrete value to an interface type performs an implicit interface conversion,
	// which creates an interface value consisting of two components:
	// * its dynamic type is the operand's type (int)
	// * its dynamic value is the operand's value (3)

	ty := reflect.TypeOf(3) // A reflect.Type
	t.Logf("reflect.TypeOf(3).String(): %s\n", ty.String())
	t.Logf("reflect.TypeOf(3): s%\n", ty)

	// reflect.TypeOf always returns a concrete type.
	// The code below prints *os.File, not io.Writer.
	var w io.Writer = os.Stdout
	t.Logf("reflect.TypeOf(os.Stdout): %s\n", reflect.TypeOf(w))

	// reflect.Type satisfies `fmt.Stringer`.
	// fmt.Printf provides a shorthand `%T` that uses `reflect.TypeOf` internally:
	fmt.Printf("%T\n", 3)
}

// A `reflect.Value` can hold a value of any type.
// `reflect.ValueOf` function accepts any interface{} and returns a
// `reflect.Value` containing the interface's dynamic value.
// The result of `reflect.ValueOf` are always concrete, but `reflect.Value`
// can hold interface values too.
func TestReflectValue(t *testing.T) {
	v := reflect.ValueOf(3)
	t.Logf("reflect.ValueOf(3): %v\n", v)
	t.Logf("reflect.ValueOf(3).String(): %v\n", v.String())
	// `reflect.Value` also satisfies `fmt.Stringer`, but unless the Value
	// holds a string, the result of the `String` method reveals only the type.
	// Instead, use the `fmt` package's `%v` verb.

	// Calling the `Type` method on a `Value` returns its type as a `reflect.Type`:
	t.Logf("reflect.ValueOf(3).Type(): %s\n", v.Type().String())

	// The inverse operation to `reflect.ValueOf` is the reflect.Value.Interface method.
	// It returns an interface{} holdong the same concrete value as the `reflect.Value`:
	x := v.Interface() // an interface{}
	i := x.(int)       // an int
	t.Logf("v.Interface().(int)%d\n", i)

	// A `reflect.Value()` and an interface{} can both hold arbitrary values.
	// The difference is that an empty interface hides the representation and
	// intrinsic operations of he value it holds and exposes none of its methods,
	// so unless we know its dynamic type and use a type assertion to peer inside it,
	// there is little we can do to the value within.
	// In constrast, a `Value` has many methods for inspecting its contents, regardless
	// of its type.
	// See implementaion of fmt.Any.
}
