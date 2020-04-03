// General-purpose unit-conversion program that reads numbers from
// its command-line arguments or from the standard input if there
// are no arguments, and converts each number into units like
// temperature in Celsius and Fahrenheit, length in feet and meters,
// weight in pounds and kilograms.

package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "./tempconv"
)

func main() {
  if len(os.Args) > 1 {
    for _, arg := range os.Args[1:] {
      printConvertedValues(arg)
    }
  } else {
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
      arg := input.Text()
      printConvertedValues(arg)
    }
  }
}

func printConvertedValues(arg string) {
  t, err := strconv.ParseFloat(arg, 64)
  if err != nil {
    fmt.Fprintf(os.Stderr, "cf: %v\n", err)
    os.Exit(1)
  }
  f := tempconv.Fahrenheit(t)
  c := tempconv.Celsius(t)
  fmt.Printf("%s = %s\n", f, c)
}

