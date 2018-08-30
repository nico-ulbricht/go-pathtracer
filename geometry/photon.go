package geometry

type Photon struct {
	X, Y       int
	Intensity  float64
	Wavelength float64
}

func NewPhoton(x, y int, intensity, wavelength float64) *Photon {
	return &Photon{
		x, y,
		intensity,
		wavelength,
	}
}
