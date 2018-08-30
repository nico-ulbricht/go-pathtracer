package geometry

type Photon struct {
	X, Y      int
	Intensity float64
}

func NewPhoton(x, y int, intensity float64) *Photon {
	return &Photon{
		x, y,
		intensity,
	}
}
