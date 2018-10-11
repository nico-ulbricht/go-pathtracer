package render

import (
	"github.com/nico-ulbricht/go-pathtracer/geometry"
)

type Tree struct {
	BoundingBox *geometry.Box
	Root        *Node
}

func NewTree(boundingBox *geometry.Box, node *Node) *Tree {
	return &Tree{boundingBox, node}
}

func NewTreeFromObjects(objects []*Object) *Tree {
	boundingBox := objects[0].Surface.BoundingBox()
	for _, object := range objects {
		boundingBox = boundingBox.Extend(object.Surface.BoundingBox())
	}

	node := NewNode(objects)
	node.Split(0)
	return NewTree(boundingBox, node)
}

func (tree *Tree) Intersect(ray *geometry.Ray) (*geometry.Intersection, *Object) {
	var closestIntersection *geometry.Intersection
	var closestObject *Object
	// for _, object := range tracer.scene.Objects {
	// 	isIntersection, intersection := object.Surface.Intersect(ray)
	// 	if isIntersection == true && (closestIntersection == nil || closestIntersection.Distance > intersection.Distance) {
	// 		closestIntersection = intersection
	// 		closestObject = object
	// 	}
	// }

	return closestIntersection, closestObject
}
