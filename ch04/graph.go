package main

import (
  "fmt"
)

var graph = make(map[string]map[string]bool)

func main() {
  addEdge("manoel", "alynne")
  addEdge("manoel", "pedro")
  addEdge("manoel", "maryna")

  fmt.Println(hasEdge("manoel", "alynne"))
  fmt.Println(hasEdge("eu", "ela"))

  adj := graph["eu"]
  fmt.Println(adj)

  m := make(map[string]string)
  fmt.Println(m)
}

func addEdge(from, to string) {
  edges := graph[from]
  if edges == nil {
    edges = make(map[string]bool)
    graph[from] = edges
  }
  edges[to] = true
}

func hasEdge(from, to string) bool {
  return graph[from][to]
}

