package render

import (
	"github.com/wahtye/gotracer/geometry"
	"github.com/wahtye/gotracer/material"
)

type Object struct {
	Material material.Material
	Surface  geometry.Surface
}

func NewObject(mat material.Material, surface geometry.Surface) *Object {
	return &Object{mat, surface}
}
