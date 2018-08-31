package geometry

type Intersection struct {
	normal *Vector
	point  *Vector
}

func NewIntersection(point, normal *Vector) *Intersection {
	return &Intersection{point, normal}
}
