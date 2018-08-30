package render

import "github.com/wahtye/gotracer/geometry"

type Scene struct {
	Surfaces []geometry.Surface
}

func NewScene() *Scene {
	surfaces := []geometry.Surface{}
	return &Scene{surfaces}
}

func (scene *Scene) AddSurface(surface geometry.Surface) {
	scene.Surfaces = append(scene.Surfaces, surface)
}
