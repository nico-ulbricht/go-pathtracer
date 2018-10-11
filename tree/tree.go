package tree

import (
	"github.com/nico-ulbricht/go-pathtracer/geometry"
	"github.com/nico-ulbricht/go-pathtracer/render"
)

type Tree struct {
	BoundingBox *geometry.Box
	Root        *Node
}

func NewTree(boundingBox *geometry.Box, node *Node) *Tree {
	return &Tree{boundingBox, node}
}

func NewTreeFromObjects(objects []*render.Object) *Tree {
	boundingBox := objects[0].Surface.BoundingBox()
	for _, object := range objects {
		boundingBox = boundingBox.Extend(object.Surface.BoundingBox())
	}

	node := NewNode(objects)
	node.Split(0)
	return NewTree(boundingBox, node)
}
