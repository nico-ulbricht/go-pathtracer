package material

import (
	"github.com/wahtye/go-pathtracer/geometry"
)

type DiffuseMaterial struct {
	reflectance float64
}

func NewDiffuseMaterial(reflectance float64) *DiffuseMaterial {
	return &DiffuseMaterial{reflectance}
}

func (mat *DiffuseMaterial) Reflect(ray *geometry.Ray, intersection *geometry.Intersection) *geometry.Ray {
	hemisphereVector := geometry.NewHemisphereVector()
	ray.Direction = hemisphereVector.RotateTowards(intersection.Normal)
	ray.Origin = intersection.Point
	ray.Intensity *= mat.reflectance * .96
	return ray
}
