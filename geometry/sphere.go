package geometry

type Sphere struct {
	center *Vector
	radius float64
}

func NewSphere(center *Vector, radius float64) *Sphere {
	return &Sphere{center, radius}
}

func (sphere *Sphere) Intersect(ray *Ray) (bool, *Intersection) {
	return false, &Intersection{}
}
