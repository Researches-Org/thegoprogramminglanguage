// comma inserts commas in a non-negative decimal integer string.

package main

import (
  "bytes"
  "fmt"
)

func main() {
  fmt.Println(commaIte("123"))
  fmt.Println(commaIte("1234"))
  fmt.Println(commaIte("12345"))
}

func comma(s string) string {
  n := len(s)
  if n <= 3 {
    return s
  }
  return comma(s[:n - 3]) + "," + s[n-3:]
}

func reverse(s string) string {
  runes := []rune(s)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

func commaIte(s string) string {
  var buf bytes.Buffer
  r := reverse(s)
  n := len(s)
  for i := 0; i < n; i += 3 {
    if i + 2 < n {
      buf.WriteString(r[i:i+3])
    } else {
      buf.WriteString(r[i:])
    }
    if (i + 3 < n) {
      buf.WriteString(",")    
    }    
  }
  return reverse(buf.String())
}

