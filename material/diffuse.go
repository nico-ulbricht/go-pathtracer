package material

import "github.com/wahtye/gotracer/geometry"

type DiffuseMaterial struct {
	reflectance float64
}

func NewDiffuseMaterial(reflectance float64) *DiffuseMaterial {
	return &DiffuseMaterial{reflectance}
}

func (mat *DiffuseMaterial) Reflect(ray *geometry.Ray, intersection *geometry.Intersection) *geometry.Ray {
	nDotDir := ray.Direction.Dot(intersection.Normal)
	ray.Direction = ray.Direction.Subtract(intersection.Normal.MultiplyScalar(2 * nDotDir))
	ray.Origin = intersection.Point
	ray.Probability *= mat.reflectance * .96
	return ray
}
