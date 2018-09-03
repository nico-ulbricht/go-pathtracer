package main

import (
	"math"
	"math/rand"

	"github.com/wahtye/gotracer/geometry"
	"github.com/wahtye/gotracer/material"
	"github.com/wahtye/gotracer/render"
)

func main() {
	scene := render.NewScene()
	sphereCount := 6
	for i := 0; i < sphereCount; i++ {
		angle := float64(2.*math.Pi/float64(sphereCount)) * float64(i)
		spherePosition := geometry.NewVector(math.Sin(angle)*125+250, math.Cos(angle)*125+250, 80)

		sphere := geometry.NewSphere(spherePosition, 30)

		var mat material.Material
		if i%2 == 0 {
			mat = material.NewDiffuseMaterial(rand.Float64())
		} else {
			mat = material.NewEmissiveMaterial(rand.Float64(), 6000)
		}

		scene.AddObject(mat, sphere)
	}

	render.NewRenderer(500, 500, scene).Render()
}
