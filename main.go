package main

import (
	"fmt"
	"sync"

	"github.com/wahtye/gotracer/geometry"
	"github.com/wahtye/gotracer/render"
)

func main() {
	width := 500
	height := 500
	photonChannel := make(chan []*geometry.Photon)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		tracer := render.NewTracer(width, height, photonChannel)
		tracer.Trace()
	}()

	go func() {
		defer wg.Done()
		<-photonChannel
		fmt.Println("got something 1")
	}()

	go func() {
		defer wg.Done()
		<-photonChannel
		fmt.Println("got something 2")
	}()

	wg.Wait()
}
