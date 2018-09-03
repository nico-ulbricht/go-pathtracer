package material

import "github.com/wahtye/gotracer/geometry"

type DiffuseMaterial struct {
	reflectance float64
}

func NewDiffuseMaterial(reflectance float64) *DiffuseMaterial {
	return &DiffuseMaterial{reflectance}
}

func (mat *DiffuseMaterial) Reflect(ray *geometry.Ray, intersection *geometry.Intersection) *geometry.Ray {
	if ray.Direction.Dot(intersection.Normal) > 0. {
		intersection.Normal = intersection.Normal.MultiplyScalar(-1.)
	}

	hemisphereVector := geometry.NewHemisphereVector()
	ray.Direction = hemisphereVector.RotateTowards(intersection.Normal)
	ray.Origin = intersection.Point
	ray.Intensity *= mat.reflectance
	return ray
}
