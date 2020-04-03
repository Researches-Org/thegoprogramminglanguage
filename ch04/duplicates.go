package main

import "fmt"

func main() {
  s := make([]string, 0)
  s = append(s, "a", "b", "b", "c", "c", "d")
  fmt.Println(s)
  s = removeDuplicate(s)
  fmt.Println(s)
}

func removeDuplicate(s []string) []string {
  if len(s) <= 1 {
    return s
  }

  i := 0
  for j := 1; j < len(s); j++ {
    if s[i] != s[j] {
      i++
      s[i] = s[j]
    }
  }

  return s[:i + 1]
}

