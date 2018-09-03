package material

import "github.com/wahtye/gotracer/geometry"

type Material interface{}

type WhiteBodyMaterial interface {
	Reflect(ray *geometry.Ray, intersection *geometry.Intersection) *geometry.Ray
}

type BlackBodyMaterial interface {
	GetIntensity(ray *geometry.Ray) float64
}
