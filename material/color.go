package material

import (
	"math"
)

type Color struct {
	r, g, b float64
}

func NewColor(r, g, b float64) *Color {
	return &Color{r, g, b}
}

func (color *Color) Blend(color2 *Color) *Color {
	return NewColor(
		math.Sqrt(color.r*color.r*.5+color2.r*color2.r*.5),
		math.Sqrt(color.g*color.g*.5+color2.g*color2.g*.5),
		math.Sqrt(color.b*color.b*.5+color2.b*color2.b*.5),
	)
}

func (color *Color) MultiplyScalar(times float64) *Color {
	return NewColor(
		color.r*times,
		color.g*times,
		color.b*times,
	)
}
