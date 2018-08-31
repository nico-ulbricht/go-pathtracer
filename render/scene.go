package render

import (
	"github.com/wahtye/gotracer/geometry"
	"github.com/wahtye/gotracer/material"
)

type Scene struct {
	Objects []*Object
}

func NewScene() *Scene {
	objects := []*Object{}
	return &Scene{objects}
}

func (scene *Scene) AddObject(material material.Material, surface geometry.Surface) {
	object := NewObject(material, surface)
	scene.Objects = append(scene.Objects, object)
}
