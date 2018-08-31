package material

import "github.com/wahtye/gotracer/geometry"

type Material interface{}

type WhiteBodyMaterial interface {
	reflect(ray *geometry.Ray, intersection *geometry.Intersection) *geometry.Ray
}

type BlackBodyMaterial interface {
	getIntensity(ray *geometry.Ray) float64
}
