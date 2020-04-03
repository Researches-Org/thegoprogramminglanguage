// reverse reverses a slice of ints in place.

package main

import (
  "fmt"
)

func main() {
  a := [...]int{0, 1, 2, 3, 4, 5}
  c := a[:]
  reverse1(&c)
  fmt.Println(a)

  b := [...]int{0, 1, 2, 3, 4, 5}
  rotateLeft(b[:], 2)
  fmt.Println(b)

}

func reverse(s []int) {
  for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
}

func rotateLeft(s []int, p int) {
  reverse(s[:p])
  reverse(s[p:])
  reverse(s)
}

// reverse using pointer
func reverse1(p *[]int) {
  for i, j := 0, len(*p) - 1; i < j; i, j = i+1, j-1 {
    (*p)[i], (*p)[j] = (*p)[j],(*p)[i] 
  }
}

// rotate left that operates in a single pass
// example:
// slice: 1 2 3 4 5
// index: 0 1 2 3 4
// position: 2
// result should be: 3 4 5 1 2
func rotateLeft1(s []int, p int) {
  if len(s) < p || p <= 0 {
    return
  }

  tail := make([]int, p)
  copy(tail, s[:p])
  copy(s, s[p:])
  for i, n := range tail {
    s[len(s) - p + i] = n
  }
}

