// Echo1 prints its command line arguments.
package main

import (
  "fmt"
  "os"
  "time"
)

func main() {
  var s, sep string
  start := time.Now()
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  end := time.Now()
  elapsed := end.Sub(start)
  fmt.Println(s)
  fmt.Println("elapsed:", elapsed)
}

