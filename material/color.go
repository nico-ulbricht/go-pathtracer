package material

import (
	"math"

	"github.com/wahtye/go-pathtracer/geometry"
)

func BlendColors(colors ...*geometry.Vector) *geometry.Vector {
	totalColors := float64(len(colors))
	rAccumulated := 0.
	gAccumulated := 0.
	bAccumulated := 0.

	for _, color := range colors {
		rAccumulated += color.X * color.X / totalColors
		gAccumulated += color.Y * color.Y / totalColors
		bAccumulated += color.Z * color.Z / totalColors
	}

	return geometry.NewVector(
		math.Sqrt(rAccumulated),
		math.Sqrt(gAccumulated),
		math.Sqrt(bAccumulated),
	)
}
