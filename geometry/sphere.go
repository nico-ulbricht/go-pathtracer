package geometry

import (
	"math"
)

type Sphere struct {
	center   *Vector
	radius   float64
	radiusSq float64
}

func NewSphere(center *Vector, radius float64) *Sphere {
	radiusSq := math.Pow(radius, 2.)
	return &Sphere{center, radius, radiusSq}
}

func (sphere *Sphere) BoundingBox() *Box {
	center := sphere.center
	return NewBox(
		NewVector(center.X-sphere.radius, center.Y-sphere.radius, center.Z-sphere.radius),
		NewVector(center.X+sphere.radius, center.Y+sphere.radius, center.Z+sphere.radius),
	)
}

func (sphere *Sphere) Center() *Vector {
	return sphere.center
}

func (sphere *Sphere) Intersect(ray *Ray) (bool, *Intersection) {
	centerToOrigin := sphere.center.Subtract(ray.Origin)
	a := 1.
	b := ray.Direction.Dot(centerToOrigin) * 2.
	c := centerToOrigin.Dot(centerToOrigin) - sphere.radiusSq
	discriminant := b*b - 4.*a*c

	if discriminant < 0 {
		return false, NoIntersection
	}

	discriminantSqrt := math.Sqrt(discriminant)
	distanceOne := -0.5 * (-b + discriminantSqrt) / a
	distanceTwo := -0.5 * (-b - discriminantSqrt) / a

	var distance float64
	if distanceOne > 0. && distanceOne < distanceTwo {
		distance = distanceOne
	} else if distanceTwo > 0. && distanceTwo < distanceOne {
		distance = distanceTwo
	} else {
		return false, NoIntersection
	}

	pointOfIntersection := ray.Origin.Add(ray.Direction.MultiplyScalar(distance))
	centerToIntersection := pointOfIntersection.Subtract(sphere.center).Normalize()
	return true, NewIntersection(distance, centerToIntersection, pointOfIntersection)
}
