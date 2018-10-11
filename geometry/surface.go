package geometry

type Surface interface {
	BoundingBox() *Box
	Intersect(*Ray) (bool, *Intersection)
}
