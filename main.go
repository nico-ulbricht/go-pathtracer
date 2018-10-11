package main

import (
	"math"

	"github.com/nico-ulbricht/go-pathtracer/geometry"
	"github.com/nico-ulbricht/go-pathtracer/material"
	"github.com/nico-ulbricht/go-pathtracer/render"
)

func main() {
	scene := buildScene(500.)
	render.NewRenderer(500, 500, scene).Render()
}

func buildScene(size float64) *render.Scene {
	scene := render.NewScene()
	diffuseMaterial := material.NewDiffuseMaterial(geometry.NewVector(1., 1., 1.), 1.)
	redColorMaterial := material.NewDiffuseMaterial(geometry.NewVector(1., 0., 0.), 1.)
	greenColorMaterial := material.NewDiffuseMaterial(geometry.NewVector(0., 1., 0.), 1.)
	emissiveMaterial := material.NewEmissiveMaterial(geometry.NewVector(1., 1., 1.), 1.)
	reflectiveMaterial := material.NewReflectiveMaterial(.8, 1.)

	floorY := .5*size/2 + size/2
	ceilingY := -.75*size/2 + size/2

	scene.AddObject(diffuseMaterial, geometry.NewPlane(geometry.NewVector(0, 0, .5*size), geometry.NewVector(0, 0, -1.)))
	scene.AddObject(diffuseMaterial, geometry.NewPlane(geometry.NewVector(0, 0, -5*size), geometry.NewVector(0, 0, 1.)))

	scene.AddObject(redColorMaterial, geometry.NewPlane(geometry.NewVector(.75*size/2+size/2, 0, 0), geometry.NewVector(-1., 0, 0)))
	scene.AddObject(greenColorMaterial, geometry.NewPlane(geometry.NewVector(-.75*size/2+size/2, 0, 0), geometry.NewVector(1., 0, 0)))

	scene.AddObject(emissiveMaterial, geometry.NewPlane(geometry.NewVector(0, ceilingY, 0), geometry.NewVector(0, 1., 0)))
	scene.AddObject(diffuseMaterial, geometry.NewPlane(geometry.NewVector(0, floorY, 0), geometry.NewVector(0, -1., 0)))

	sphereCount := 6
	radius := 20.
	for i := 0; i < sphereCount; i++ {
		angle := float64(2.*math.Pi/float64(sphereCount)) * float64(i)
		circleRadius := 100.
		circlePosition := geometry.NewVector(math.Sin(angle)*circleRadius, floorY, math.Cos(angle)*circleRadius)
		offsetPosition := geometry.NewVector(size/2, -radius, -size/4)
		sphere := geometry.NewSphere(offsetPosition.Add(circlePosition), radius)

		var mat material.Material
		if i%2 != 0 {
			mat = diffuseMaterial
		} else {
			mat = reflectiveMaterial
		}

		scene.AddObject(mat, sphere)
	}

	return scene
}
