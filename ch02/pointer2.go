package main

import "fmt"

var p = f()

func f() *int {
  v := 1
  return &v
}

func incr(p *int) int {
  *p++ // increments what p points to; does not change p
  return *p
}

func main() {
  fmt.Println(p)
  fmt.Println(*p)
  *p = 2
  fmt.Println(*p)

  v := 1
  incr(&v)              // side effect: v is now 2
  fmt.Println(incr(&v)) // "3" (and v is now 3)
}

