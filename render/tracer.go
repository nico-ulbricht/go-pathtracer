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
	ray := geometry.NewZeroRay()

	for {
		for _, photon := range tracer.photonBuffer {
			photon = tracer.processPhoton(photon, ray)
		}

		tracer.photonChannel <- tracer.photonBuffer
	}
}

func (tracer *Tracer) processPhoton(photon *geometry.Photon, ray *geometry.Ray) *geometry.Photon {
	xPos := rand.Intn(tracer.width)
	yPos := rand.Intn(tracer.height)
	photon.X = xPos
	photon.Y = yPos

	ray = tracer.camera.GetRayAt(xPos, yPos, ray)
	for _, object := range tracer.scene.Objects {
		isIntersection, _ := object.Surface.Intersect(ray)
		if isIntersection == true {
			photon.Intensity = 1.
			return photon
		}
	}

	photon.Intensity = 0.
	return photon
}
