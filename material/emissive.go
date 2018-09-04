package material

import "github.com/wahtye/go-pathtracer/geometry"

type EmissiveMaterial struct {
	intensity, temperature float64
}

func NewEmissiveMaterial(intensity, temperature float64) *EmissiveMaterial {
	return &EmissiveMaterial{intensity, temperature}
}

func (mat *EmissiveMaterial) GetIntensity(ray *geometry.Ray) float64 {
	return mat.intensity * ray.Intensity
}
