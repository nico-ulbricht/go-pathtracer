package material

import "github.com/nico-ulbricht/go-pathtracer/geometry"

type DiffuseMaterial struct {
	color       *geometry.Vector
	reflectance float64
}

func NewDiffuseMaterial(color *geometry.Vector, reflectance float64) *DiffuseMaterial {
	return &DiffuseMaterial{color, reflectance}
}

func (mat *DiffuseMaterial) Reflect(ray *geometry.Ray, intersection *geometry.Intersection) *geometry.Ray {
	hemisphereVector := geometry.NewHemisphereVector()
	ray.AccumulatedColor = ray.AccumulatedColor.Add(mat.color.Sq())
	ray.ColorBounces++
	ray.Direction = hemisphereVector.RotateTowards(intersection.Normal)
	ray.Origin = intersection.Point
	ray.Intensity *= mat.reflectance * .96
	return ray
}
