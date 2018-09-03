package geometry

type Intersection struct {
	Distance float64
	Normal   *Vector
	Point    *Vector
}

var NoIntersection = &Intersection{0, NewZeroVector(), NewZeroVector()}

func NewIntersection(distance float64, normal, point *Vector) *Intersection {
	return &Intersection{distance, normal, point}
}
