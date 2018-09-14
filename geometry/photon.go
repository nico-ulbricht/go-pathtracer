package geometry

type Photon struct {
	X, Y      int
	Color     *Vector
	Intensity float64
}

func NewPhoton(x, y int, color *Vector, intensity float64) *Photon {
	return &Photon{
		x, y,
		color,
		intensity,
	}
}
