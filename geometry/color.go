package geometry

import "math"

type Color struct {
	R, G, B float64
}

func NewColor(r, g, b float64) *Color {
	return &Color{r, g, b}
}

func BlendColors(colors ...*Color) *Color {
	totalColors := float64(len(colors))
	rAccumulated := 0.
	gAccumulated := 0.
	bAccumulated := 0.

	for _, color := range colors {
		rAccumulated += color.R * color.R / totalColors
		gAccumulated += color.G * color.G / totalColors
		bAccumulated += color.B * color.B / totalColors
	}

	return NewColor(
		math.Sqrt(rAccumulated),
		math.Sqrt(gAccumulated),
		math.Sqrt(bAccumulated),
	)
}

func (color *Color) MultiplyScalar(times float64) *Color {
	return NewColor(
		color.R*times,
		color.G*times,
		color.B*times,
	)
}
