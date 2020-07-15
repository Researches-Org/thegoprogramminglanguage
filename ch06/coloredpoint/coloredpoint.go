package coloredpoint

import (
    "image/color"
    "../point"
)

type ColoredPoint struct {
    point.Point
    Color color.RGBA
}

