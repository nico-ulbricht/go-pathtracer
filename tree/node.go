package tree

import "github.com/nico-ulbricht/go-pathtracer/geometry"

type Node struct {
	Left     *Node
	Right    *Node
	Surfaces []geometry.Surface
}

func NewNode(surfaces []geometry.Surface) *Node {
	return &Node{
		Left:     nil,
		Right:    nil,
		Surfaces: surfaces,
	}
}

func (node *Node) Split(depth int) *Node {
	return node
}
