// Echo3 print its command line arguments
package main

import (
  "fmt"
  "strings"
  "os"
  "time"
)

func main() {
  start := time.Now()
  fmt.Println(os.Args[0] + " " + strings.Join(os.Args[1:], " "))
  end := time.Now()
  elapsed := end.Sub(start)
  fmt.Println("elapsed:", elapsed)
}

