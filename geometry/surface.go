package geometry

type Surface interface {
	Intersect(*Ray) (bool, *Intersection)
}
