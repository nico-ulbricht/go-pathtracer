package main

import (
	"math"

	"github.com/nico-ulbricht/go-pathtracer/geometry"
	"github.com/nico-ulbricht/go-pathtracer/material"
	"github.com/nico-ulbricht/go-pathtracer/render"
	"github.com/nico-ulbricht/go-pathtracer/tree"
)

func main() {
	scene := buildScene(500.)
	tree.NewTreeFromObjects(scene.Objects)
	render.NewRenderer(500, 500, scene).Render()
}

func buildScene(size float64) *render.Scene {
	scene := render.NewScene()
	floorY := .5*size/2 + size/2
	diffuseMaterial := material.NewDiffuseMaterial(geometry.NewVector(1., 1., 1.), 1.)
	reflectiveMaterial := material.NewReflectiveMaterial(.8, 1.)
	redColorMaterial := material.NewDiffuseMaterial(geometry.NewVector(1., 0., 0.), 1.)
	greenColorMaterial := material.NewDiffuseMaterial(geometry.NewVector(0., 1., 0.), 1.)
	emissiveMaterial := material.NewEmissiveMaterial(geometry.NewVector(1., 1., 1.), 1.)
	ceilingY := -.75*size/2 + size/2

	scene.AddObject(diffuseMaterial, geometry.NewPlane(geometry.NewVector(0, 0, .5*size), geometry.NewVector(0, 0, -1.)))
	scene.AddObject(diffuseMaterial, geometry.NewPlane(geometry.NewVector(0, 0, -5*size), geometry.NewVector(0, 0, 1.)))
	scene.AddObject(redColorMaterial, geometry.NewPlane(geometry.NewVector(.75*size/2+size/2, 0, 0), geometry.NewVector(-1., 0, 0)))
	scene.AddObject(greenColorMaterial, geometry.NewPlane(geometry.NewVector(-.75*size/2+size/2, 0, 0), geometry.NewVector(1., 0, 0)))
	scene.AddObject(emissiveMaterial, geometry.NewPlane(geometry.NewVector(0, ceilingY, 0), geometry.NewVector(0, 1., 0)))
	scene.AddObject(diffuseMaterial, geometry.NewPlane(geometry.NewVector(0, floorY, 0), geometry.NewVector(0, -1., 0)))

	sphereCount := 12
	radius := 20.
	for i := 0; i < 5; i++ {
		for j := 0; j < sphereCount; j++ {
			angle := float64(2.*math.Pi/float64(sphereCount)) * float64(j)
			circleRadius := 100.
			circleHeight := floorY - radius*2.5*float64(i)
			circlePosition := geometry.NewVector(math.Sin(angle)*circleRadius, circleHeight, math.Cos(angle)*circleRadius)
			offsetPosition := geometry.NewVector(size/2+radius, -radius, -size/4)
			sphere := geometry.NewSphere(offsetPosition.Add(circlePosition), radius)

			var mat material.Material
			if (i+j)%2 != 0 {
				mat = diffuseMaterial
			} else {
				mat = reflectiveMaterial
			}

			scene.AddObject(mat, sphere)
		}
	}

	return scene
}
