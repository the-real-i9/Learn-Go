package main

import (
	"flag"
	"fmt"
)

type Celcius float64

func (c Celcius) String() string {
	return fmt.Sprintf("%.1fdegC", c)
}

type Farenheit float64

func FToC(f Farenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}

type Kelvin float64

func KToC(k Kelvin) Celcius {
	return Celcius(k - 273.15)
}

type celciusFlag struct {
	Celcius
}

func (f *celciusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		f.Celcius = Celcius(value)
		return nil
	case "F":
		f.Celcius = FToC(Farenheit(value))
		return nil
	case "K":
		f.Celcius = KToC(Kelvin(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CelciusFlag(name string, value Celcius, usage string) *Celcius {
	f := celciusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celcius
}

var temp = CelciusFlag("temp", 20.0, "the temperature")

func parFlag() {
	flag.Parse()
	fmt.Println(*temp)
}
