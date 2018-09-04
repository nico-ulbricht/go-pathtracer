package render

import (
	"github.com/wahtye/go-pathtracer/geometry"
)

type Camera struct {
	width, height int
	position      *geometry.Vector
}

func NewCamera(width, height int, distance float64) *Camera {
	position := geometry.NewVector(float64(width-width/2), float64(height-height/2), -distance)
	return &Camera{width, height, position}
}

func (camera *Camera) GetRayAt(x, y int, ray *geometry.Ray) *geometry.Ray {
	ray.Reset()

	ray.Origin.X = camera.position.X
	ray.Origin.Y = camera.position.Y
	ray.Origin.Z = camera.position.Z

	point := geometry.NewVector(float64(x), float64(y), 0.)
	ray.Direction = point.Subtract(camera.position).Normalize()

	return ray
}
