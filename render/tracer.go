package render

import (
	"math/rand"

	"github.com/wahtye/gotracer/geometry"
)

type Tracer struct {
	width, height int
	camera        *Camera
	scene         *Scene
	photonBuffer  []*geometry.Photon
	photonChannel chan []*geometry.Photon
}

func NewTracer(width, height int, scene *Scene, photonChannel chan []*geometry.Photon) *Tracer {
	photonBuffer := make([]*geometry.Photon, 250000)
	for idx := range photonBuffer {
		photonBuffer[idx] = geometry.NewPhoton(0, 0, 0)
	}

	camera := NewCamera()
	return &Tracer{
		width, height,
		camera,
		scene,
		photonBuffer,
		photonChannel,
	}
}

func (tracer *Tracer) Trace() {
	for {
		for _, photon := range tracer.photonBuffer {
			xPos := rand.Intn(tracer.width)
			yPos := rand.Intn(tracer.height)

			tracedPhoton := tracer.traceAtPosition(xPos, yPos)
			photon.X = tracedPhoton.X
			photon.Y = tracedPhoton.Y
			photon.Intensity = tracedPhoton.Intensity
		}

		tracer.photonChannel <- tracer.photonBuffer
	}
}

func (tracer *Tracer) traceAtPosition(x, y int) *geometry.Photon {
	ray := tracer.camera.GetRayAt(x, y)
	for _, surface := range tracer.scene.Surfaces {
		isIntersection, _ := surface.Intersect(ray)
		if isIntersection == true {
			return geometry.NewPhoton(x, y, 1.)
		}
	}

	return geometry.NewPhoton(x, y, 0.)
}
