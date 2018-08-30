package main

import (
	"sync"

	"github.com/wahtye/gotracer/geometry"
	"github.com/wahtye/gotracer/render"
)

func main() {
	width := 500
	height := 500
	photonChannel := make(chan []*geometry.Photon)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		tracer := render.NewTracer(width, height, photonChannel)
		tracer.Trace()
	}()

	go func() {
		defer wg.Done()
		gatherer := render.NewGatherer(width, height, photonChannel)
		gatherer.Gather()
	}()

	wg.Wait()
}
