package material

import "github.com/wahtye/go-pathtracer/geometry"

type ColorMaterial struct {
	color       *geometry.Color
	reflectance float64
}

func NewColorMaterial(color *geometry.Color, reflectance float64) *ColorMaterial {
	return &ColorMaterial{color, reflectance}
}

func (mat *ColorMaterial) Reflect(ray *geometry.Ray, intersection *geometry.Intersection) *geometry.Ray {
	hemisphereVector := geometry.NewHemisphereVector()
	ray.Color = geometry.BlendColors(ray.Color, mat.color)
	ray.Direction = hemisphereVector.RotateTowards(intersection.Normal)
	ray.Origin = intersection.Point
	ray.Intensity *= mat.reflectance * .96
	return ray
}
