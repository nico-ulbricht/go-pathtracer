package render

import (
	"math/rand"

	"github.com/wahtye/gotracer/geometry"
)

type Tracer struct {
	width, height int
	photonBuffer  []*geometry.Photon
	photonChannel chan []*geometry.Photon
}

func NewTracer(width, height int, photonChannel chan []*geometry.Photon) *Tracer {
	photonBuffer := make([]*geometry.Photon, 250000)
	for idx := range photonBuffer {
		photonBuffer[idx] = geometry.NewPhoton(0, 0, 0, 0)
	}

	return &Tracer{
		width, height,
		photonBuffer,
		photonChannel,
	}
}

func (tracer *Tracer) Trace() {
	for {
		for _, photon := range tracer.photonBuffer {
			xPos := rand.Intn(tracer.width)
			yPos := rand.Intn(tracer.height)

			tracedPhoton := geometry.NewPhoton(xPos, yPos, 1., 400.)
			photon.X = tracedPhoton.X
			photon.Y = tracedPhoton.Y
			photon.Intensity = tracedPhoton.Intensity
			photon.Wavelength = tracedPhoton.Wavelength
		}

		tracer.photonChannel <- tracer.photonBuffer
	}
}
