package render

import (
	"math/rand"

	"github.com/nico-ulbricht/go-pathtracer/geometry"
	"github.com/nico-ulbricht/go-pathtracer/material"
)

const MAX_BOUNCES = 4

type Tracer struct {
	width, height int
	camera        *Camera
	tree          *Tree
	photonBuffer  []*geometry.Photon
	photonChannel chan []*geometry.Photon
}

func NewTracer(width, height int, tree *Tree, photonChannel chan []*geometry.Photon) *Tracer {
	photonBuffer := make([]*geometry.Photon, 100000)
	for idx := range photonBuffer {
		photonBuffer[idx] = geometry.NewPhoton(0, 0, geometry.NewVector(0, 0, 0), 0)
	}

	camera := NewCamera(width, height, 500)
	return &Tracer{
		width, height,
		camera,
		tree,
		photonBuffer,
		photonChannel,
	}
}

func (tracer *Tracer) Trace() {
	ray := geometry.NewZeroRay()

	for {
		for _, photon := range tracer.photonBuffer {
			photon.X = rand.Intn(tracer.width)
			photon.Y = rand.Intn(tracer.height)
			ray = tracer.camera.GetRayAt(photon.X, photon.Y, ray)
			photon = tracer.processPhoton(photon, ray)
		}

		tracer.photonChannel <- tracer.photonBuffer
	}
}

func (tracer *Tracer) processPhoton(photon *geometry.Photon, ray *geometry.Ray) *geometry.Photon {
	photon.Intensity = 0.
	isIntersection, closestIntersection, closestObject := tracer.tree.Intersect(ray)

	if isIntersection == false || ray.Bounces > MAX_BOUNCES {
		return photon
	}

	switch objectMaterial := closestObject.Material.(type) {
	case material.WhiteBodyMaterial:
		reflectionRay := objectMaterial.Reflect(ray, closestIntersection)
		reflectionRay.Bounces++
		return tracer.processPhoton(photon, reflectionRay)
	case material.BlackBodyMaterial:
		photon.Color = objectMaterial.GetColor(ray)
		photon.Intensity = objectMaterial.GetIntensity(ray)
		return photon
	}

	return photon
}
