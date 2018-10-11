package render

import (
	"math/rand"

	"github.com/nico-ulbricht/go-pathtracer/geometry"
	"github.com/nico-ulbricht/go-pathtracer/material"
)

type Object struct {
	Id       uint64
	Material material.Material
	Surface  geometry.Surface
}

func NewObject(mat material.Material, surface geometry.Surface) *Object {
	id := rand.Uint64()
	return &Object{id, mat, surface}
}
