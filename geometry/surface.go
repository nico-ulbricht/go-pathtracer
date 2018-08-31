package geometry

type Surface interface {
	Intersect(*Ray, *Intersection) (bool, *Intersection)
}
