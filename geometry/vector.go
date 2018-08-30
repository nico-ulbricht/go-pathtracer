package geometry

type Vector struct {
	x, y, z float64
}

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func (vec *Vector) Add(vec2 *Vector) *Vector {
	return &Vector{
		vec.x + vec2.x,
		vec.y + vec2.y,
		vec.z + vec2.z,
	}
}

func (vec *Vector) Subtract(vec2 *Vector) *Vector {
	return &Vector{
		vec.x - vec2.x,
		vec.y - vec2.y,
		vec.z - vec2.z,
	}
}

func (vec *Vector) MultiplyScalar(multiplier float64) *Vector {
	return &Vector{
		vec.x * multiplier,
		vec.y * multiplier,
		vec.z * multiplier,
	}
}

func (vec *Vector) DivideScalar(divider float64) *Vector {
	return &Vector{
		vec.x * divider,
		vec.y * divider,
		vec.z * divider,
	}
}
