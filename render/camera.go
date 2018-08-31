package render

import "github.com/wahtye/gotracer/geometry"

type Camera struct{}

func NewCamera() *Camera {
	return &Camera{}
}

func (camera *Camera) GetRayAt(x, y int, ray *geometry.Ray) *geometry.Ray {
	ray.Origin.X = float64(x)
	ray.Origin.Y = float64(y)
	ray.Origin.Z = 0.
	ray.Direction.X = 0.
	ray.Direction.Y = 0.
	ray.Direction.Z = 1.
	ray.Probability = 1.
	return ray
}
