package geometry

type Photon struct {
	X, Y      int
	Color     *Color
	Intensity float64
}

func NewPhoton(x, y int, color *Color, intensity float64) *Photon {
	return &Photon{
		x, y,
		color,
		intensity,
	}
}
