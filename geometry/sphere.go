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
		return false, &Intersection{}
	}

	distanceOriginToCenter := originToCenter.Magnitude()
	degreesDirectionOrigin := math.Acos(dotDirectionCenter)
	distanceDirectionToCenter := math.Sin(degreesDirectionOrigin) * distanceOriginToCenter

	// Ray misses Sphere
	if distanceDirectionToCenter-sphere.radius > 0.000001 {
		return false, &Intersection{}
	}

	return true, &Intersection{}
}
