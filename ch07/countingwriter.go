package main

import (
  "bytes"
  "fmt"
  "io"
)

type ByteCounter struct {
  w     io.Writer
  count int64
}

func (bc *ByteCounter) Write(p []byte) (int, error) {
  n, err := bc.w.Write(p)
  bc.count += int64(n)
  return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
 bc := ByteCounter{w, 0}
 return &bc, &bc.count
}

func main() {  
  var buf bytes.Buffer
  var w, c = CountingWriter(&buf)
  w.Write([]byte("hello"))
  fmt.Println(*c)
}

