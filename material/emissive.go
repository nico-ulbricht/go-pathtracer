package material

import (
	"github.com/wahtye/go-pathtracer/geometry"
)

type EmissiveMaterial struct {
	color     *geometry.Vector
	intensity float64
}

func NewEmissiveMaterial(color *geometry.Vector, intensity float64) *EmissiveMaterial {
	return &EmissiveMaterial{color, intensity}
}

func (mat *EmissiveMaterial) GetColor(ray *geometry.Ray) *geometry.Vector {
	return BlendColors(ray.Color, mat.color)
}

func (mat *EmissiveMaterial) GetIntensity(ray *geometry.Ray) float64 {
	return mat.intensity * ray.Intensity
}
