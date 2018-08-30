package geometry

type Intersection struct {
	point *Vector
}

func NewIntersection(point *Vector) *Intersection {
	return &Intersection{point}
}
