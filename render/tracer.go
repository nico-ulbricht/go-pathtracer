package render

import (
	"math/rand"

	"github.com/wahtye/gotracer/geometry"
	"github.com/wahtye/gotracer/material"
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
	intersection := geometry.NewZeroIntersection()
	ray := geometry.NewZeroRay()

	for {
		for _, photon := range tracer.photonBuffer {
			photon.X = rand.Intn(tracer.width)
			photon.Y = rand.Intn(tracer.height)
			ray = tracer.camera.GetRayAt(photon.X, photon.Y, ray)
			photon = tracer.processPhoton(photon, ray, intersection)
		}

		tracer.photonChannel <- tracer.photonBuffer
	}
}

func (tracer *Tracer) processPhoton(photon *geometry.Photon, ray *geometry.Ray, intersection *geometry.Intersection) *geometry.Photon {
	photon.Intensity = 0.

	var closestIntersection *geometry.Intersection
	var closestObject *Object
	for _, object := range tracer.scene.Objects {
		isIntersection, intersection := object.Surface.Intersect(ray, intersection)
		if isIntersection == true && (closestIntersection == nil || closestIntersection.Distance > intersection.Distance) {
			closestIntersection = intersection
			closestObject = object
		}
	}

	if closestObject == nil || ray.Probability < .6 {
		return photon
	}

	switch objectMaterial := closestObject.Material.(type) {
	case material.WhiteBodyMaterial:
		reflectionRay := objectMaterial.Reflect(ray, closestIntersection)
		return tracer.processPhoton(photon, reflectionRay, intersection)
	case material.BlackBodyMaterial:
		photon.Intensity = objectMaterial.GetIntensity(ray)
		return photon
	}

	return photon
}
