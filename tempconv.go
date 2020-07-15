package main

import (
  "flag"
  "fmt"
)

type Celsius float64

func (c Celsius) String() string {
  return fmt.Sprintf("%v C", float64(c))
}

type Fahrenheit float64

func (f Fahrenheit) String() string {
  return fmt.Sprintf("%v F", float64(f))
}

func FToC(f Fahrenheit) Celsius {
  return Celsius((f - 32) * 5/9)
}

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
  var unit string
  var value float64
  fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
  switch unit {
  case "C":
    f.Celsius = Celsius(value)
    return nil
  case "F":
    f.Celsius = FToC(Fahrenheit(value))
    return nil
  }
  return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
  f := celsiusFlag{value}
  flag.CommandLine.Var(&f, name, usage)
  return &f.Celsius
}

var temperature = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
  flag.Parse()
  fmt.Println("Temperature ", *temperature)
}

