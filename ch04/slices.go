package main

import (
  "fmt"
)

func main() {
  s := []int{5, 6, 7, 8, 9}
  fmt.Println(remove(s, 2)) // "[5 6 8 9]"

  r := []int{5, 6, 7, 8, 9}
  fmt.Println(remove1(r, 2)) // "[5 6 9 8]"
}

func equal(x, y []string) bool {
  if (len(x) != len(y)) {
    return false
  }
  for i := range x {
    if x[i] != y[i] {
      return false
    }
  }
  return true
}

func remove(slice []int, i int) []int {
  copy(slice[i:], slice[i + 1:])
  return slice[:len(slice) - 1]
}

func remove1(slice []int, i int) []int {
  slice[i] = slice[len(slice) - 1]
  return slice[:len(slice) - 1]
}

