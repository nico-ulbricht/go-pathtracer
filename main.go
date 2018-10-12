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

var colors = []*geometry.Vector{
	geometry.NewVector(1, .5, .5),
	geometry.NewVector(.5, 1, .5),
	geometry.NewVector(.5, .5, 1),
	geometry.NewVector(.5, 1, 1),
	geometry.NewVector(1, 1, 1),
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

	scene.AddObject(diffuseMaterial,
		geometry.NewBox(geometry.NewVector(0, 0, 249), geometry.NewVector(500, 500, 250)))
	scene.AddObject(redColorMaterial,
		geometry.NewBox(geometry.NewVector(500, 0, -400), geometry.NewVector(501, 500, 500)))
	scene.AddObject(greenColorMaterial,
		geometry.NewBox(geometry.NewVector(0, 0, -400), geometry.NewVector(1, 500, 500)))
	scene.AddObject(emissiveMaterial,
		geometry.NewBox(geometry.NewVector(0, ceilingY, -400), geometry.NewVector(500, ceilingY+1., 500)))
	scene.AddObject(diffuseMaterial,
		geometry.NewBox(geometry.NewVector(0, floorY, -400), geometry.NewVector(500, floorY+1., 500)))

	sphereCount := 24
	radius := 10.
	for i := 0; i < 12; i++ {
		for j := 0; j < sphereCount; j++ {
			angle := float64(2.*math.Pi/float64(sphereCount)) * float64(j)
			circleRadius := 100.
			circleHeight := floorY - radius*2.5*float64(i)
			circlePosition := geometry.NewVector(math.Sin(angle)*circleRadius, circleHeight, math.Cos(angle)*circleRadius)
			offsetPosition := geometry.NewVector(size/2+radius, -radius, -size/4)

			var mat material.Material
			var sphere *geometry.Sphere
			if (i+j)%6 == 0 {
				sphere = geometry.NewSphere(offsetPosition.Add(circlePosition), radius*1.5)
				mat = reflectiveMaterial
			} else {
				// diffuse := material.NewDiffuseMaterial(
				// 	colors[int(math.Floor(rand.Float64()*5))],
				// 	rand.Float64()*.5+.5)
				sphere = geometry.NewSphere(offsetPosition.Add(circlePosition), radius)
				mat = diffuseMaterial
			}

			scene.AddObject(mat, sphere)
		}
	}

	return scene
}
