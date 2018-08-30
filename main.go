package main

import (
	"github.com/wahtye/gotracer/geometry"
	"github.com/wahtye/gotracer/render"
)

func main() {
	scene := render.NewScene()
	scene.AddSurface(geometry.NewSphere(geometry.NewVector(0, 0, 0), 5))

	render.NewRenderer(500, 500, scene).Render()
}
