package main

import (
	"math"
	"math/rand"

	"github.com/wahtye/gotracer/geometry"
	"github.com/wahtye/gotracer/material"
	"github.com/wahtye/gotracer/render"
)

func main() {
	scene := buildScene(500.)
	render.NewRenderer(500, 500, scene).Render()
}

func buildScene(size float64) *render.Scene {
	scene := render.NewScene()
	scene.AddObject(material.NewDiffuseMaterial(1.), geometry.NewPlane(geometry.NewVector(0, 0, 110), geometry.NewVector(0, 0, -1.)))
	scene.AddObject(material.NewDiffuseMaterial(1.), geometry.NewPlane(geometry.NewVector(0, 0, -110), geometry.NewVector(0, 0, 1.)))

	sphereCount := 6
	for i := 0; i < sphereCount; i++ {
		angle := float64(2.*math.Pi/float64(sphereCount)) * float64(i)
		spherePosition := geometry.NewVector(math.Sin(angle)*size/4.+size/2., math.Cos(angle)*size/4.+size/2., 80)

		sphere := geometry.NewSphere(spherePosition, size/16.)

		var mat material.Material
		if i%2 == 0 {
			mat = material.NewDiffuseMaterial(.8)
		} else {
			mat = material.NewEmissiveMaterial(rand.Float64()*2.+1., 6000)
		}

		scene.AddObject(mat, sphere)
	}

	return scene
}
