package material

import "github.com/wahtye/gotracer/geometry"

type EmissiveMaterial struct {
	intensity, temperature float64
}

func NewEmissiveMaterial(intensity, temperature float64) *EmissiveMaterial {
	return &EmissiveMaterial{intensity, temperature}
}

func (mat *EmissiveMaterial) getIntensity(ray *geometry.Ray) float64 {
	return 1.
}
