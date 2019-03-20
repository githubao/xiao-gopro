// 实现flag.Value接口
// author: baoqiang
// time: 2019/3/18 下午2:28
package ch07

import (
	"flag"
	"fmt"
)

func RunCelsiusFlag() {
	temp := CelsiusFlag("temp", 20.0, "the temperature")
	flag.Parse()
	fmt.Println(temp.String())
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32.0)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (c *Celsius) String() string {
	return fmt.Sprintf("%.6f°C", float64(*c))
}
