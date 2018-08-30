package render

import "github.com/wahtye/gotracer/geometry"

type Camera struct {
}

func NewCamera() *Camera {
	return &Camera{}
}

func (camera *Camera) GetRayAt(x, y int) *geometry.Ray {
	origin := geometry.NewVector(float64(x), float64(y), 0.)
	direction := geometry.NewVector(0., 0., 1.)
	return geometry.NewRay(origin, direction)
}
