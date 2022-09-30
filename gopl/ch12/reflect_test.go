package ch12

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	ty := reflect.TypeOf(3)
	t.Log(ty.String())

	var w io.Writer = os.Stdout
	t.Log(reflect.TypeOf(w))

	v := reflect.ValueOf(3)
	t.Log(v)
	t.Logf("%v\n", v)
	t.Log(v.String())

	x := v.Interface()
	i := x.(int)
	t.Logf("%d\n", i)
}
