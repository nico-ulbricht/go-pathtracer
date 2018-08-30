package main

import (
	"github.com/wahtye/gotracer/render"
)

func main() {
	render.NewRenderer(500, 500).Render()
}
