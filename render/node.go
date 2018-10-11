package render

import (
	"github.com/nico-ulbricht/go-pathtracer/geometry"
)

const MAX_NODE_OBJECTS = 4
const MAX_DEPTH = 10

type Node struct {
	Axis        geometry.Axis
	BoundingBox *geometry.Box
	IsLeaf      bool
	Left        *Node
	Right       *Node
	Point       float64
	Objects     []*Object
}

func NewNode(objects []*Object) *Node {
	boundingBox := objects[0].Surface.BoundingBox()
	for _, object := range objects {
		boundingBox = boundingBox.Extend(object.Surface.BoundingBox())
	}

	return &Node{
		BoundingBox: boundingBox,
		IsLeaf:      false,
		Left:        nil,
		Right:       nil,
		Objects:     objects,
	}
}

func (node *Node) Split(depth int) *Node {
	if len(node.Objects) <= MAX_NODE_OBJECTS || depth > MAX_DEPTH {
		node.IsLeaf = true
		return node
	}

	axis := geometry.AxisIndexed[depth%3]
	points := []float64{}
	for _, object := range node.Objects {
		boundingBox := object.Surface.BoundingBox()
		points = append(points, boundingBox.MinPosition.GetAxis(axis))
		points = append(points, boundingBox.MaxPosition.GetAxis(axis))
	}

	medianPoint := points[int(len(points)/2)]
	leftObjects := []*Object{}
	rightObjects := []*Object{}

	for _, object := range node.Objects {
		boundingBox := object.Surface.BoundingBox()
		left, right := boundingBox.Partition(axis, medianPoint)
		if left == true {
			leftObjects = append(leftObjects, object)
		}

		if right == true {
			rightObjects = append(rightObjects, object)
		}
	}

	node.Left = NewNode(leftObjects).Split(depth + 1)
	node.Right = NewNode(rightObjects).Split(depth + 1)
	return node
}

func (node *Node) Intersect(ray *geometry.Ray) (bool, *geometry.Intersection, *Object) {
	// if node.IsLeaf {
	return node.IntersectObjects(ray)
	// }

	// var first *Node
	// var second *Node
	// rayPoint := ray.Origin.GetAxis(node.Axis)
	// rayDirection := ray.Direction.GetAxis(node.Axis)
	// leftFirst := (rayPoint < node.Point) || (rayPoint == node.Point && rayDirection <= 0)
	// if leftFirst {
	// 	first = node.Left
	// 	second = node.Right
	// } else {
	// 	first = node.Right
	// 	second = node.Left
	// }
	// return false, nil, nil
}

func (node *Node) IntersectObjects(ray *geometry.Ray) (bool, *geometry.Intersection, *Object) {
	var closestIntersection *geometry.Intersection
	var closestObject *Object
	for _, object := range node.Objects {
		isIntersection, intersection := object.Surface.Intersect(ray)
		if isIntersection == true && (closestIntersection == nil || closestIntersection.Distance > intersection.Distance) {
			closestIntersection = intersection
			closestObject = object
		}
	}

	isIntersection := closestIntersection != nil
	return isIntersection, closestIntersection, closestObject
}
