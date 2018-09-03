package geometry

type Intersection struct {
	Distance float64
	Normal   *Vector
	Point    *Vector
}

func NewIntersection(distance float64, point, normal *Vector) *Intersection {
	return &Intersection{distance, point, normal}
}

func NewZeroIntersection() *Intersection {
	return &Intersection{0, NewZeroVector(), NewZeroVector()}
}
