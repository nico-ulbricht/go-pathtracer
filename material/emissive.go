package material

import (
	"github.com/wahtye/go-pathtracer/geometry"
)

type EmissiveMaterial struct {
	color     *geometry.Color
	intensity float64
}

func NewEmissiveMaterial(color *geometry.Color, intensity float64) *EmissiveMaterial {
	return &EmissiveMaterial{color, intensity}
}

func (mat *EmissiveMaterial) GetColor(ray *geometry.Ray) *geometry.Color {
	return mat.color
}

func (mat *EmissiveMaterial) GetIntensity(ray *geometry.Ray) float64 {
	return mat.intensity * ray.Intensity
}
