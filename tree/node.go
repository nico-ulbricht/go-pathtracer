package tree

import (
	"time"

	"github.com/nico-ulbricht/go-pathtracer/geometry"
	"github.com/nico-ulbricht/go-pathtracer/render"
)

const MAX_NODE_OBJECTS = 3

type Node struct {
	Left    *Node
	Right   *Node
	Objects []*render.Object
}

func NewNode(objects []*render.Object) *Node {
	return &Node{
		Left:    nil,
		Right:   nil,
		Objects: objects,
	}
}

func (node *Node) Split(depth int) *Node {
	if len(node.Objects) <= MAX_NODE_OBJECTS {
		return nil
	}

	axis := geometry.AxisIndexed[depth%3]
	points := []float64{}
	for _, object := range node.Objects {
		boundingBox := object.Surface.BoundingBox()
		points = append(points, boundingBox.MinPosition.GetAxis(axis))
		points = append(points, boundingBox.MaxPosition.GetAxis(axis))
	}

	medianPoint := points[int(len(points)/2)]
	leftObjects := []*render.Object{}
	rightObjects := []*render.Object{}

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

	time.Sleep(time.Second)
	node.Left = NewNode(leftObjects)
	node.Right = NewNode(rightObjects)
	node.Left.Split(depth + 1)
	node.Right.Split(depth + 1)

	return node
}
