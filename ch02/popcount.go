package main

import (
  "fmt"
  "time"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
  for i := range pc {
    pc[i] = pc[i/2] + byte(i&1)
  }
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
  return int(pc[byte(x>>(0*8))] +
             pc[byte(x>>(1*8))] +
             pc[byte(x>>(2*8))] +
             pc[byte(x>>(3*8))] +
             pc[byte(x>>(4*8))] +
             pc[byte(x>>(5*8))] +
             pc[byte(x>>(6*8))] +
             pc[byte(x>>(7*8))])
}

func PopCount1(x uint64) int {
  y := x
  count := 0
  for y > 0 {
    if (y & 1) == 1 {
      count++
    }
    y = y >> 1
  }
  return count
}

func logTime(y uint64, pcount func(x uint64) int) {
  start := time.Now()
  c := pcount(y)
  end := time.Now()
  elapsed := end.Sub(start)
  fmt.Printf("Pop count %d in elapsed %v\n", c, elapsed)
}

func main() {
  logTime(127, PopCount)
  logTime(127, PopCount1)
}
