// Echo3 print its command line arguments
package main

import (
  "fmt"
  "strings"
  "os"
)

func main() {
  fmt.Println(os.Args[0] + " " + strings.Join(os.Args[1:], " "))
}

