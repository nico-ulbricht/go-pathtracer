package tree

import "github.com/nico-ulbricht/go-pathtracer/geometry"

type Tree struct {
	BoundingBox *geometry.Box
	Root        *Node
}

func NewTree(boundingBox *geometry.Box, node *Node) *Tree {
	return &Tree{boundingBox, node}
}

func NewTreeFromSurfaces(surfaces []geometry.Surface) *Tree {
	boundingBox := surfaces[0].BoundingBox()
	for _, surface := range surfaces {
		boundingBox = boundingBox.Extend(surface.BoundingBox())
	}

	node := NewNode(surfaces)
	return NewTree(boundingBox, node)
}
