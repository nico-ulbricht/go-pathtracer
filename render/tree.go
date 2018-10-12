package render

import (
	"github.com/nico-ulbricht/go-pathtracer/geometry"
)

type Tree struct {
	Root *Node
}

func NewTree(node *Node) *Tree {
	return &Tree{node}
}

func NewTreeFromObjects(objects []*Object) *Tree {
	node := NewNode(objects)
	node.Split(0)
	return NewTree(node)
}

func (tree *Tree) Intersect(ray *geometry.Ray) (bool, *geometry.Intersection, *Object) {
	return tree.Root.Intersect(ray)
}
