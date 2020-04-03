package main

import (
  "fmt"
)

func zero_v0(ptr *[32]byte) {
  for i := range ptr {
    ptr[i] = 0
  }
}

func zero_v1(ptr *[32]byte) {
  *ptr = [32]byte{}
}

func zeroAndPrint(ptr *[32]byte, fn func(*[32]byte)) {
  fn(ptr)
  fmt.Println(ptr)
}

func main() {
  a := [32]byte{1, 2, 3}
  zeroAndPrint(&a, zero_v0)

  b := [32]byte{4, 5, 6}
  zeroAndPrint(&b, zero_v1)
}

