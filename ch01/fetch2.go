// Fetch prints the content found at a URL

package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "strings"
)

func main() {
  for _, url := range os.Args[1:] {
    prefix := "http://"
    urlWithPrefix := url
    if !strings.HasPrefix(url, prefix) {
      urlWithPrefix = prefix + url
    }
    resp, err := http.Get(urlWithPrefix)
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
      os.Exit(1)
    }
    b, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: reanding %s: %v\n", urlWithPrefix, err)
      os.Exit(1)
    }
    fmt.Printf("%s", b)
  }
}

