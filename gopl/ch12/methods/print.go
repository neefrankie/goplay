package methods

import (
	"fmt"
	"reflect"
	"strings"
)

func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("typ0e %s\n", t)

	for i := 0; i < v.NumField(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name, strings.TrimPrefix(methodType.String(), "func"))
	}
}
