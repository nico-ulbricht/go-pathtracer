package material

import "github.com/wahtye/go-pathtracer/geometry"

type Material interface{}

type WhiteBodyMaterial interface {
	Reflect(ray *geometry.Ray, intersection *geometry.Intersection) *geometry.Ray
}

type BlackBodyMaterial interface {
	GetColor(ray *geometry.Ray) *geometry.Vector
	GetIntensity(ray *geometry.Ray) float64
}
