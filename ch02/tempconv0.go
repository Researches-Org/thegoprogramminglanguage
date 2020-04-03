// Package tempconv performs Celcius and Fahrenheit temperature computations.

package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string {
  return fmt.Sprintf("%gºC", c)
}

const (
  AbsoluteZeroC Celsius = -273.15
  FreezingC     Celsius = 0
  BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32)*5/9) }

func main() {
  fmt.Printf("%g\n", BoilingC - FreezingC)
  boilingF := CToF(BoilingC)
  fmt.Printf("%g\n", boilingF - CToF(FreezingC))
  // compilation error: type mismatch
  // fmt.Printf("%g\n", boilingF - FreezingC)

  var c Celsius
  var f Fahrenheit
  fmt.Println(c == 0)          // true
  fmt.Println(f >= 0)          // true
  // fmt.Println(c == f)       compilation error: type mismatch
  fmt.Println(c == Celsius(f)) // true, because both c and f are zero
  fmt.Println(BoilingC)
}

