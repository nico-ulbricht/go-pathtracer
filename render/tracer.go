package render

import (
	"math/rand"

	"github.com/wahtye/go-pathtracer/geometry"
	"github.com/wahtye/go-pathtracer/material"
)

const MAX_BOUNCES = 4

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

	camera := NewCamera(width, height, 500)
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

	var closestIntersection *geometry.Intersection
	var closestObject *Object
	for _, object := range tracer.scene.Objects {
		isIntersection, intersection := object.Surface.Intersect(ray)
		if isIntersection == true && (closestIntersection == nil || closestIntersection.Distance > intersection.Distance) {
			closestIntersection = intersection
			closestObject = object
		}
	}

	if closestObject == nil || ray.Bounces > MAX_BOUNCES {
		return photon
	}

	switch objectMaterial := closestObject.Material.(type) {
	case material.WhiteBodyMaterial:
		reflectionRay := objectMaterial.Reflect(ray, closestIntersection)
		reflectionRay.Bounces++
		return tracer.processPhoton(photon, reflectionRay)
	case material.BlackBodyMaterial:
		photon.Intensity = objectMaterial.GetIntensity(ray)
		return photon
	}

	return photon
}
