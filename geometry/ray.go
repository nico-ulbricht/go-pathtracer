package geometry

type Ray struct {
	Origin, Direction *Vector
}

func NewRay(origin, direction *Vector) *Ray {
	return &Ray{origin, direction}
}

func NewZeroRay() *Ray {
	origin := NewVector(0., 0., 0.)
	direction := NewVector(0., 0., 1.)
	return NewRay(origin, direction)
}
