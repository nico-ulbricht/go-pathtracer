package material

import (
	"github.com/nico-ulbricht/go-pathtracer/geometry"
)

type EmissiveMaterial struct {
	color     *geometry.Vector
	intensity float64
}

func NewEmissiveMaterial(color *geometry.Vector, intensity float64) *EmissiveMaterial {
	return &EmissiveMaterial{color, intensity}
}

func (mat *EmissiveMaterial) GetColor(ray *geometry.Ray) *geometry.Vector {
	accumulatedColor := ray.AccumulatedColor.Add(mat.color.Sq())
	return accumulatedColor.DivideScalar(float64(ray.ColorBounces + 1)).Sqrt()
}

func (mat *EmissiveMaterial) GetIntensity(ray *geometry.Ray) float64 {
	return mat.intensity * ray.Intensity
}
