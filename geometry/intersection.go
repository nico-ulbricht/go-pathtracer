package geometry

type Intersection struct {
	Distance float64
	Normal   *Vector
	Point    *Vector
}

func NewIntersection(distance float64, normal, point *Vector) *Intersection {
	return &Intersection{distance, normal, point}
}

func NewZeroIntersection() *Intersection {
	return &Intersection{0, NewZeroVector(), NewZeroVector()}
}
