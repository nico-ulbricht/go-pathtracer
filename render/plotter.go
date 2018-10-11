package render

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/nico-ulbricht/go-pathtracer/material"
)

const GAMMA = .75

type Plotter struct {
	width, height int
	canvasChannel chan []*Pixel
}

func NewPlotter(width, height int, canvasChannel chan []*Pixel) *Plotter {
	return &Plotter{
		width, height,
		canvasChannel,
	}
}

func (plotter *Plotter) Plot() {
	img := image.NewRGBA(image.Rect(0, 0, plotter.width, plotter.height))
	for {
		canvas := <-plotter.canvasChannel
		for x := 0; x < plotter.width; x++ {
			for y := 0; y < plotter.height; y++ {
				position := y*plotter.width + x
				pixel := canvas[position]

				accumulation := math.Min(1., math.Pow(pixel.accumulation/float64(pixel.samples), GAMMA))
				accumulatedColor := material.BlendColors(pixel.colors...)

				rColor := uint8(math.Floor(accumulatedColor.X * accumulation * 255.))
				gColor := uint8(math.Floor(accumulatedColor.Y * accumulation * 255.))
				bColor := uint8(math.Floor(accumulatedColor.Z * accumulation * 255.))

				img.SetRGBA(x, y, color.RGBA{rColor, gColor, bColor, 255})
			}
		}

		file, _ := os.Create("image.png")
		png.Encode(file, img)
	}
}
