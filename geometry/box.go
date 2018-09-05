package geometry

import (
	"math"
)

type Box struct {
	minPosition, maxPosition *Vector
}

func NewBox(minPosition, maxPosition *Vector) *Box {
	return &Box{minPosition, maxPosition}
}

func (box *Box) Intersect(ray *Ray) (bool, *Intersection) {
	directionInverse := ray.Direction.Invert()
	tmin := 1e-10
	tmax := math.Inf(1)

	for a := 0; a < 3; a++ {
		var axis string
		if a == 0 {
			axis = "X"
		} else if a == 1 {
			axis = "Y"
		} else {
			axis = "Z"
		}

		t0 := (box.minPosition.GetAxis(axis) - ray.Origin.GetAxis(axis)) * directionInverse.GetAxis(axis)
		t1 := (box.maxPosition.GetAxis(axis) - ray.Origin.GetAxis(axis)) * directionInverse.GetAxis(axis)

		if directionInverse.GetAxis(axis) < 0 {
			t0, t1 = t1, t0
		}
		if t0 > tmin {
			tmin = t0
		}
		if t1 < tmax {
			tmax = t1
		}
		if tmax < tmin {
			return false, NoIntersection
		}
	}

	pointOfIntersection := ray.Origin.Add(ray.Direction.MultiplyScalar(tmin))
	normal := NewVector(0, 0, -1)
	return true, NewIntersection(tmin, normal, pointOfIntersection)
}
