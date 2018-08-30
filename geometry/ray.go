package geometry

type Ray struct {
	origin, direction *Vector
}

func NewRay(origin, direction *Vector) *Ray {
	return &Ray{origin, direction}
}
