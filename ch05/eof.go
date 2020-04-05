package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func main() {
    fmt.Println(readRunes())
}

func readRunes() (string, error) {
    slice := make([]string, 0)
    in := bufio.NewReader(os.Stdin)
    for {
        r, _, err := in.ReadRune()
        if err == io.EOF {
            break // finished reading
        }
        if err != nil {
            return "", fmt.Errorf("read failed: %v", err)
        }
        slice = append(slice, fmt.Sprintf("%v", r))
    }
    return fmt.Sprintf("%v", slice), nil
}

