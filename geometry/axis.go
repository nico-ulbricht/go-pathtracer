package geometry

type Axis string

const (
	AxisX Axis = "X"
	AxisY Axis = "Y"
	AxisZ Axis = "Z"
)

var AxisIndexed = [3]Axis{AxisX, AxisY, AxisZ}
var AxisAlignedNormals = map[string]*Vector{
	"-X": NewVector(-1., 0, 0),
	"+X": NewVector(1., 0, 0),
	"-Y": NewVector(0, -1., 0),
	"+Y": NewVector(0, 1., 0),
	"-Z": NewVector(0, 0, -1.),
	"+Z": NewVector(0, 0, 1.),
}
