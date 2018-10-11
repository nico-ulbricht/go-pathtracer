package geometry

import (
	"fmt"
	"math"
)

type Box struct {
	center                   *Vector
	minPosition, maxPosition *Vector
}

func NewBox(minPosition, maxPosition *Vector) *Box {
	center := minPosition.Average(maxPosition)
	return &Box{center, minPosition, maxPosition}
}

func (box *Box) BoundingBox() *Box {
	return box
}

func (box *Box) Center() *Vector {
	return box.center
}

func (box *Box) Intersect(ray *Ray) (bool, *Intersection) {
	var normal *Vector
	directionInverse := ray.Direction.Invert()
	tMin := 1e-10
	tMax := math.Inf(1)

	for i := 0; i < 3; i++ {
		axis := AxisIndexed[i]
		t0 := (box.minPosition.GetAxis(axis) - ray.Origin.GetAxis(axis)) * directionInverse.GetAxis(axis)
		t1 := (box.maxPosition.GetAxis(axis) - ray.Origin.GetAxis(axis)) * directionInverse.GetAxis(axis)

		isNegative := directionInverse.GetAxis(axis) < 0
		if isNegative == true {
			t0, t1 = t1, t0
		}

		if t0 > tMin {
			tMin = t0
			if isNegative == true {
				alignedAxis := fmt.Sprintf("+%s", axis)
				normal = AxisAlignedNormals[alignedAxis]
			} else {
				alignedAxis := fmt.Sprintf("-%s", axis)
				normal = AxisAlignedNormals[alignedAxis]
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

func (box *Box) Extend(box2 *Box) *Box {
	return NewBox(
		box.minPosition.Min(box2.minPosition),
		box.maxPosition.Max(box2.maxPosition),
	)
}
