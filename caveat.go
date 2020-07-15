package main

import (
  "bytes"
  "fmt"
  "io"
)

const debug = false

func main() {
  // var buf *bytes.Buffer
  var buf io.Writer

  if debug {
    buf = new(bytes.Buffer) // enable collection of output
  }

  f(buf) // NOTE: subtly incorrect! if buf is of type *bytes.Buffer
  
  if debug {
    // ...use buf...
  }

}

func f(out io.Writer) {
  // ...do something
  if out != nil {
    fmt.Println("Trying to call write...")
    out.Write([]byte("done!\n"))
  }
}

