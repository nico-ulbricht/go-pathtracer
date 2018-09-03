package render

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

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

				accumulation := math.Min(1., pixel.accumulation/float64(pixel.samples))
				intensity := uint8(math.Floor(accumulation * 255.))
				img.SetRGBA(x, y, color.RGBA{intensity, intensity, intensity, 255})
			}
		}

		file, _ := os.Create("image.png")
		png.Encode(file, img)
	}
}
