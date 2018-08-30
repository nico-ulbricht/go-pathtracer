package geometry

import "math"

type Vector struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func (vec *Vector) Add(vec2 *Vector) *Vector {
	return NewVector(
		vec.X+vec2.X,
		vec.Y+vec2.Y,
		vec.Z+vec2.Z,
	)
}

func (vec *Vector) Subtract(vec2 *Vector) *Vector {
	return NewVector(
		vec.X-vec2.X,
		vec.Y-vec2.Y,
		vec.Z-vec2.Z,
	)
}

func (vec *Vector) MultiplyScalar(multiplier float64) *Vector {
	return NewVector(
		vec.X*multiplier,
		vec.Y*multiplier,
		vec.Z*multiplier,
	)
}

func (vec *Vector) DivideScalar(divider float64) *Vector {
	return NewVector(
		vec.X/divider,
		vec.Y/divider,
		vec.Z/divider,
	)
}

func (vec *Vector) Dot(vec2 *Vector) float64 {
	return vec.X*vec2.X + vec.Y*vec2.Y + vec.Z*vec2.Z

}

func (vec *Vector) Normalize() *Vector {
	magnitude := vec.Magnitude()
	return vec.DivideScalar(magnitude)
}

func (vec *Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(vec.X, 2.) + math.Pow(vec.Y, 2.) + math.Pow(vec.Z, 2.))
}
