// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.

package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  counts := make(map[string]int)
  files := os.Args[1:]
  if len(files) == 0 {
    countLines(os.Stdin, counts)
  } else {
    for _, arg := range files {
      f, err := os.Open(arg)
      if err != nil {
        fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
        continue
      }
      if countLines(f, counts) {
        fmt.Println("file has duplicate:", arg)
      }
      f.Close()
    }
  }
  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}

func countLines(f *os.File, counts map[string]int) bool {
  input := bufio.NewScanner(f)
  hasDuplicate := false
  for input.Scan() {
    s := input.Text()
    if counts[s] > 0 {
      hasDuplicate = true
    }
    counts[s]++
  }
  return hasDuplicate
  // NOTE: ignoring potencial errors from input.Err()
}

