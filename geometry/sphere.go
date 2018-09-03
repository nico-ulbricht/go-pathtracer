package geometry

import (
	"math"
)

type Sphere struct {
	center *Vector
	radius float64
}

func NewSphere(center *Vector, radius float64) *Sphere {
	return &Sphere{center, radius}
}

func (sphere *Sphere) Intersect(ray *Ray) (bool, *Intersection) {
	originToCenter := sphere.center.Subtract(ray.Origin)
	dotDirectionCenter := ray.Direction.Dot(originToCenter.Normalize())

	// Sphere on the other side of the Ray
	if dotDirectionCenter < 0 {
		return false, NoIntersection
	}

	distanceOriginToCenter := originToCenter.Magnitude()
	degreesDirectionOrigin := math.Acos(dotDirectionCenter)
	distanceDirectionToCenter := math.Sin(degreesDirectionOrigin) * distanceOriginToCenter

	// Ray misses Sphere
	if distanceDirectionToCenter-sphere.radius > 0.000001 {
		return false, NoIntersection
	}

	distancePow := math.Pow(distanceDirectionToCenter, 2)
	distanceOriginToMidpoint := math.Sqrt(math.Pow(distanceOriginToCenter, 2) - distancePow)
	distanceIntersectionToMidpoint := math.Sqrt(math.Pow(sphere.radius, 2) - distancePow)
	distanceOriginToIntersection1 := distanceOriginToMidpoint - distanceIntersectionToMidpoint
	distanceOriginToIntersection2 := distanceOriginToMidpoint + distanceIntersectionToMidpoint

	distanceToIntersection := math.Min(distanceOriginToIntersection1, distanceOriginToIntersection2)
	pointOfIntersection := ray.Origin.Add(ray.Direction.MultiplyScalar(distanceToIntersection))
	centerToIntersection := pointOfIntersection.Subtract(sphere.center).Normalize()

	return true, NewIntersection(distanceToIntersection, centerToIntersection, pointOfIntersection)
}
