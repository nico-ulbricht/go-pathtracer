package render

import (
	"github.com/nico-ulbricht/go-pathtracer/geometry"
	"github.com/nico-ulbricht/go-pathtracer/material"
)

type Object struct {
	Material material.Material
	Surface  geometry.Surface
}

func NewObject(mat material.Material, surface geometry.Surface) *Object {
	return &Object{mat, surface}
}
