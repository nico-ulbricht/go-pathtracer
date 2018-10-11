package geometry

import (
	"log"
	"math"
)

type Box struct {
	minPosition, maxPosition *Vector
}

func NewBox(minPosition, maxPosition *Vector) *Box {
	return &Box{minPosition, maxPosition}
}

func (box *Box) BoundingBox() *Box {
	return box
}

func (box *Box) Intersect(ray *Ray) (bool, *Intersection) {
	var normal *Vector
	directionInverse := ray.Direction.Invert()
	tMin := 1e-10
	tMax := math.Inf(1)

	for i := 0; i < 3; i++ {
		var axis string
		if i == 0 {
			axis = "X"
		} else if i == 1 {
			axis = "Y"
		} else {
			axis = "Z"
		}

		t0 := (box.minPosition.GetAxis(axis) - ray.Origin.GetAxis(axis)) * directionInverse.GetAxis(axis)
		t1 := (box.maxPosition.GetAxis(axis) - ray.Origin.GetAxis(axis)) * directionInverse.GetAxis(axis)

		isNegative := directionInverse.GetAxis(axis) < 0
		if isNegative == true {
			t0, t1 = t1, t0
		}

		if t0 > tMin {
			tMin = t0
			if isNegative == true {
				normal = getAxisAlignedNormal("+" + axis)
			} else {
				normal = getAxisAlignedNormal("-" + axis)
			}
		}

		if t1 < tMax {
			tMax = t1
		}

		if tMax < tMin {
			return false, NoIntersection
		}
	}

	if normal == nil {
		return false, NoIntersection
	}

	pointOfIntersection := ray.Origin.Add(ray.Direction.MultiplyScalar(tMin))
	return true, NewIntersection(tMin, normal, pointOfIntersection)
}

func getAxisAlignedNormal(axis string) *Vector {
	switch axis {
	case "-X":
		return NewVector(-1., 0, 0)
	case "+X":
		return NewVector(1., 0, 0)
	case "-Y":
		return NewVector(0, -1., 0)
	case "+Y":
		return NewVector(0, 1., 0)
	case "-Z":
		return NewVector(0, 0, -1.)
	case "+Z":
		return NewVector(0, 0, 1.)
	}

	log.Fatalf("Incorrect axis: %s\n", axis)
	return NewVector(0, 0, 0)
}

func (box *Box) Extend(box2 *Box) *Box {
	return NewBox(
		box.minPosition.Min(box2.minPosition),
		box.maxPosition.Max(box2.maxPosition),
	)
}
