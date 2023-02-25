package main

import (
	"flag"
	"fmt"
)

// eg7.7 解释为什么帮助信息在它的默认值是20.0没有包含°C的情况下输出了°C。
// 原因在于默认值20.0其实是Celsius类型，而Celsius类型中定义了String()方法输出°C
var temp = CelsiusFlag("temp", 20.0, "the temperature")

// 对tempFlag加入支持开尔文温度
// [K] = [°C] + 273.15
// [K] = ([°F] + 459.67)*5/9
func main() {
	flag.Parse()
	fmt.Println(*temp)
}

type Celsius float64
type Fahrenheit float64
type Kelvins float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func CToK(c Celsius) Kelvins { return Kelvins(c + 273.15) }

func KToC(k Kelvins) Celsius { return Celsius(k - 273.15) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FToK(f Fahrenheit) Kelvins { return Kelvins((f + 459.67) * 5 / 9) }

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvins(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
