package geometry

type Ray struct {
	Bounces           uint8
	AccumulatedColor  *Vector
	ColorBounces      uint8
	Origin, Direction *Vector
	Intensity         float64
}

func NewRay(origin, direction *Vector) *Ray {
	normalizedDirection := direction.Normalize()
	return &Ray{
		0.,
		NewVector(0., 0., 0.),
		0.,
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
	ray.ColorBounces = 0.
	ray.AccumulatedColor = NewVector(0., 0., 0.)
	ray.Intensity = 1.
	return ray
}
