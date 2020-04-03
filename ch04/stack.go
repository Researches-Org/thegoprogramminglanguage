package main

import (
  "errors"
  "fmt"
)

const (
  MAX_SIZE = 10 
)

func create() []int {
  return make([]int, 0, MAX_SIZE)
}

func top(stack []int) (int, error) {
  if len(stack) == 0 {
    return 0,errors.New("empty stack")
  }

  return stack[len(stack) - 1], nil
}

func push(stack []int, i int) ([]int, error) {
  if len(stack) == MAX_SIZE {
    return stack, errors.New("full stack")
  }
  stack = append(stack, i)
  return stack, nil
}

func pop(stack []int) (int, []int, error) {
  top,error := top(stack)
  if error != nil {
    return 0, nil, error
  }
  stack = stack[:len(stack) - 1]
  return top, stack, nil
}

func main() {
  stack := create()
  stack, _ = push(stack, 1)
  stack, _ = push(stack, 2)
  fmt.Println(top(stack))
  fmt.Println(top(stack))
  fmt.Println(stack)

  top, stack, _ := pop(stack)
  fmt.Println("top:", top)
  fmt.Println(stack)
}

