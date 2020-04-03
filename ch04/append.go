package main

import (
  "fmt"
)

func main() {
  var x, y []int
  for i := 0; i < 10; i++ {
    y = appendInt(x, i)
    fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
    x = y
  }

  y = append0(y, y...)
  fmt.Println(y)
}

func appendInt(x []int, y int) []int {
  var z []int
  zlen := len(x) + 1
  if zlen <= cap(x) {
    // there is room to grow. Extend the slice
    z = x[:zlen]
  } else {
    zcap := zlen
    if zcap < 2 * len(x) {
      zcap = 2 * len(x)
    }
    z = make([]int, zlen, zcap)
    copy(z, x) // built-in function; see text
  }
  z[len(x)] = y
  return z
}

func append0(x []int, y ...int) []int {
  var z []int
  zlen := len(x) + len(y)
  if zlen <= cap(x) {
    z = x[:zlen]
  } else {
    zcap := zlen
    if zcap < 2 * len(x) {
      zcap = 2 * len(x)
    }
    z = make([]int, zlen, zcap)
    copy(z, x)
  }
  copy(z[len(x):], y)
  return z
}

