package main

import (
    "image/color"
    "fmt"
    cp "./coloredpoint"
    point "./point"
)

func main() {
    var coloredp cp.ColoredPoint
    coloredp.X = 1
    fmt.Println(coloredp.Point.X)
    coloredp.Point.Y = 2
    fmt.Println(coloredp.Y)

    red := color.RGBA{255, 0, 0, 255}
    blue := color.RGBA{0, 0, 255, 255}
    var p = cp.ColoredPoint{point.Point{1, 1}, red}
    var q = cp.ColoredPoint{point.Point{5, 4}, blue}
    fmt.Println(p.Distance(q.Point))
    p.ScaleBy(2)
    q.ScaleBy(2)
    fmt.Println(p.Distance(q.Point))
}

