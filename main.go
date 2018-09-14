package main

import (
	"math"

	"github.com/wahtye/go-pathtracer/geometry"
	"github.com/wahtye/go-pathtracer/material"
	"github.com/wahtye/go-pathtracer/render"
)

func main() {
	scene := buildScene(500.)
	render.NewRenderer(500, 500, scene).Render()
}

func buildScene(size float64) *render.Scene {
	scene := render.NewScene()
	diffuseMaterial := material.NewDiffuseMaterial(1.)
	redColorMaterial := material.NewColorMaterial(geometry.NewVector(1., .2, .2), 1.)
	greenColorMaterial := material.NewColorMaterial(geometry.NewVector(.2, 1., .2), 1.)
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

	scene.AddObject(diffuseMaterial, geometry.NewBox(geometry.NewVector(125., floorY-60., -150.), geometry.NewVector(175., floorY, -100.)))

	sphereCount := 2
	radius := size / 12.
	for i := 0; i < sphereCount; i++ {
		angle := float64(1.8*math.Pi/float64(sphereCount)) * float64(i)
		spherePosition := geometry.NewVector(math.Sin(angle)*size/6.+size/2., floorY-radius, math.Cos(angle)*size/6.-size/4.)
		sphere := geometry.NewSphere(spherePosition, radius)

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
