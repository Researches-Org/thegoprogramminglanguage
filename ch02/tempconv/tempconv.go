// Package tempconv performs Celsius and Fahrenheit conversions.
// Add t yp es, const ants, and f unc t ions to tempconv for pro cessing temp eratures in t he Kelv in s c ale, w here zero Kelv in is −273.15°C and a dif ference of 1K has t he s ame mag ni- tude as 1°C
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
  AbsoluteZeroC Celsius = -273.15
  FreezingC     Celsius = 0
  BoilingC      Celsius = 100
  ZeroKelvin    Kelvin  = 0
)

func (c Celsius) String() string {
  return fmt.Sprintf("%gºC", c)
}

func (f Fahrenheit) String() string {
  return fmt.Sprintf("%gºF", f)
}

func (k Kelvin) String() string {
  return fmt.Sprintf("%gºK", k)
}

