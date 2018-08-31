package geometry

type Ray struct {
	Origin, Direction *Vector
	Probability       float64
}

func NewRay(origin, direction *Vector) *Ray {
	return &Ray{origin, direction, 1.}
}

func NewZeroRay() *Ray {
	origin := NewVector(0., 0., 0.)
	direction := NewVector(0., 0., 1.)
	return NewRay(origin, direction)
}
