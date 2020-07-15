package main

import (
    "fmt"
    "math"
)

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

// same thing, but as method of Point type
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

type Path []Point

// returns the distance travelled along the path
func (path Path) Distance() float64 {
    sum := 0.0
    for i := range path {
        if i > 0 {
            sum += path[i - 1].Distance(path[i])
        }
    }
    return sum
}

func main() {
    p := Point{0, 0}
    q := Point{4, 4}

    fmt.Println(p.Distance(q))

    perim := Path{
        {1, 1},
        {5, 1},
        {5, 4},
        {1, 1},
    }

    fmt.Println(perim.Distance())
}

