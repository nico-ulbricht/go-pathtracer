package render

import (
	"fmt"
	"time"

	"github.com/nico-ulbricht/go-pathtracer/geometry"
)

type Pixel struct {
	accumulation float64
	colors       []*geometry.Vector
	samples      uint64
}

func NewPixel() *Pixel {
	return &Pixel{0, []*geometry.Vector{}, 0}
}

type Gatherer struct {
	width, height int
	canvas        []*Pixel
	canvasChannel chan []*Pixel
	photonChannel chan []*geometry.Photon
}

func NewGatherer(width, height int, canvasChannel chan []*Pixel, photonChannel chan []*geometry.Photon) *Gatherer {
	canvas := make([]*Pixel, width*height)
	for idx := range canvas {
		canvas[idx] = NewPixel()
	}

	return &Gatherer{
		width, height,
		canvas,
		canvasChannel,
		photonChannel,
	}
}

func (gatherer *Gatherer) Gather() {
	iterations := 0
	start := time.Now()
	total := 0
	for {
		photons := <-gatherer.photonChannel
		for _, photon := range photons {
			position := photon.Y*gatherer.width + photon.X
			gatherer.canvas[position].accumulation += photon.Intensity
			gatherer.canvas[position].colors = append(gatherer.canvas[position].colors, photon.Color)
			gatherer.canvas[position].samples++
			total++
		}

		iterations++
		if iterations%5 == 0 {
			secondsPassed := time.Since(start).Seconds()
			raysPerSecond := int(float64(total) / secondsPassed)
			fmt.Printf("%d rays per second, %d total (%.2fs)\n", raysPerSecond, total, secondsPassed)
			gatherer.canvasChannel <- gatherer.canvas
		}
	}
}
