package material

import "github.com/wahtye/gotracer/geometry"

type DiffuseMaterial struct {
	reflectance float64
}

func NewDiffuseMaterial(reflectance float64) *DiffuseMaterial {
	return &DiffuseMaterial{reflectance}
}

func (mat *DiffuseMaterial) reflect(ray *geometry.Ray) *geometry.Ray {
	ray.Probability *= mat.reflectance
	return ray
}
