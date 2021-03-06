package geometry

type Plane struct {
	center *Vector
	normal *Vector
}

func NewPlane(center, normal *Vector) *Plane {
	return &Plane{center, normal}
}

func (plane *Plane) BoundingBox() *Box {
	return NewBox(
		NewVector(-INF, -INF, -INF),
		NewVector(INF, INF, INF),
	)
}

func (plane *Plane) Intersect(ray *Ray) (bool, *Intersection) {
	directionDotNormal := ray.Direction.Dot(plane.normal)

	// Behind or parallel
	if directionDotNormal >= 0. {
		return false, NoIntersection
	}

	rayToPlane := plane.center.Subtract(ray.Origin)
	rayToPlaneDotNormal := rayToPlane.Dot(plane.normal)

	distance := rayToPlaneDotNormal / directionDotNormal
	point := ray.Origin.Add(ray.Direction.MultiplyScalar(distance))
	return true, NewIntersection(distance, plane.normal, point)
}
