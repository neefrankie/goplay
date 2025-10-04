package ch7

import (
	"flag"
	"fmt"

	"gopl/ch2"
)

type celsiusFlag struct {
	ch2.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Scanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Celsius = ch2.Celsius(value)
		return nil

	case "F", "°F":
		f.Celsius = ch2.FToC(ch2.Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value ch2.Celsius, usage string) *ch2.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
