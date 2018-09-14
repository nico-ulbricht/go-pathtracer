package geometry

type Ray struct {
	Bounces           uint8
	Color             *Vector
	Origin, Direction *Vector
	Intensity         float64
}

func NewRay(origin, direction *Vector) *Ray {
	normalizedDirection := direction.Normalize()
	return &Ray{
		0.,
		NewVector(1., 1., 1.),
		origin,
		normalizedDirection,
		1.,
	}
}

func NewZeroRay() *Ray {
	origin := NewVector(0., 0., 0.)
	direction := NewVector(0., 0., 1.)
	return NewRay(origin, direction)
}

func (ray *Ray) Reset() *Ray {
	ray.Bounces = 0.
	ray.Color = NewVector(1., 1., 1.)
	ray.Intensity = 1.
	return ray
}
