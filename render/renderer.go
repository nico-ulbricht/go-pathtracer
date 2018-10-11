package render

import (
	"runtime"
	"sync"

	"github.com/nico-ulbricht/go-pathtracer/geometry"
)

type Renderer struct {
	width, height int
	scene         *Scene
	canvasChannel chan []*Pixel
	photonChannel chan []*geometry.Photon
}

func NewRenderer(width, height int, scene *Scene) *Renderer {
	canvasChannel := make(chan []*Pixel)
	photonChannel := make(chan []*geometry.Photon)
	return &Renderer{width, height, scene, canvasChannel, photonChannel}
}

func (renderer *Renderer) Render() {
	numCpus := runtime.NumCPU()

	var wg sync.WaitGroup
	wg.Add(numCpus)
	for i := 0; i < numCpus-2; i++ {
		go func() {
			defer wg.Done()
			NewTracer(
				renderer.width, renderer.height,
				renderer.scene,
				renderer.photonChannel,
			).Trace()
		}()
	}

	go func() {
		defer wg.Done()
		NewGatherer(
			renderer.width, renderer.height,
			renderer.canvasChannel,
			renderer.photonChannel,
		).Gather()
	}()

	go func() {
		defer wg.Done()
		NewPlotter(
			renderer.width,
			renderer.height,
			renderer.canvasChannel,
		).Plot()
	}()

	wg.Wait()
}
