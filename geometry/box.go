package geometry

import (
	"fmt"
	"math"
)

type Box struct {
	center                   *Vector
	MinPosition, MaxPosition *Vector
}

func NewBox(MinPosition, MaxPosition *Vector) *Box {
	center := MinPosition.Average(MaxPosition)
	return &Box{center, MinPosition, MaxPosition}
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
		t0 := (box.MinPosition.GetAxis(axis) - ray.Origin.GetAxis(axis)) * directionInverse.GetAxis(axis)
		t1 := (box.MaxPosition.GetAxis(axis) - ray.Origin.GetAxis(axis)) * directionInverse.GetAxis(axis)

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
		box.MinPosition.Min(box2.MinPosition),
		box.MaxPosition.Max(box2.MaxPosition),
	)
}

func (box *Box) Partition(axis Axis, medianPoint float64) (left, right bool) {
	minPoint := box.MinPosition.GetAxis(axis)
	maxPoint := box.MaxPosition.GetAxis(axis)
	return minPoint <= medianPoint, maxPoint >= medianPoint
}
