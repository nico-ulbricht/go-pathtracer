package geometry

type Surface interface {
	BoundingBox() *Box
	Center() *Vector
	Intersect(*Ray) (bool, *Intersection)
}
