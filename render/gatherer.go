package render

import (
	"github.com/wahtye/gotracer/geometry"
)

type Pixel struct {
	accumulation float64
	samples      uint64
}

func NewPixel() *Pixel {
	return &Pixel{0, 0}
}

type Gatherer struct {
	width, height int
	canvas        []*Pixel
	photonChannel chan []*geometry.Photon
}

func NewGatherer(width, height int, photonChannel chan []*geometry.Photon) *Gatherer {
	canvas := make([]*Pixel, width*height)
	for idx := range canvas {
		canvas[idx] = NewPixel()
	}

	return &Gatherer{
		width, height,
		canvas,
		photonChannel,
	}
}

func (gatherer *Gatherer) Gather() {
	for {
		photons := <-gatherer.photonChannel
		for _, photon := range photons {
			position := photon.Y*gatherer.width + photon.X
			gatherer.canvas[position].accumulation += photon.Intensity
			gatherer.canvas[position].samples++
			gatherer.photonCount++
		}
	}
}
