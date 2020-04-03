// map functions

package main

import (
  "fmt"
)

func main() {
  m := make(map[string]int)
  
  s1 := []string{"a"}
  s2 := []string{"a", "b"}
  s3 := []string{"a", "b", "c"}

  Add(m, s1)
  Add(m, s2)
  Add(m, s2)
  Add(m, s3)
  Add(m, s3)
  Add(m, s3)
  
  fmt.Println(Count(m, s1))
  fmt.Println(Count(m, s2))
  fmt.Println(Count(m, s3))
}

func k(list []string) string {
  return fmt.Sprintf("%q", list) 
}

func Add(m map[string]int, list []string) {
  m[k(list)]++
}

func Count(m map[string]int, list []string) int {
  return m[k(list)]
}

func equal(x, y map[string]int) bool {
  if len(x) != len(y) {
    return false
  }

  for k, xv := range x {
    if yv, ok := y[k]; !ok || yv != xv {
      return false
    }
  }

  return true
}

