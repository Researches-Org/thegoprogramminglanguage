package main

import (
    "fmt"
    "os"
    "golang.org/x/net/html"
)

var depth int

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "outline: %v\n", err)
        os.Exit(1)
    }
    forEachNode(doc, startElement, endElement)
}

func startElement(n * html.Node) {
    if n.Type == html.ElementNode {
        fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data)
        depth++
    }
}

func endElement(n * html.Node) {
    if n.Type == html.ElementNode {
        depth--
        fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
    }
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
    apply(n, pre)
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }
    apply(n, post)
}

func apply(n *html.Node, f func(n *html.Node)) {
    if f != nil {
        f(n)
    }
}

