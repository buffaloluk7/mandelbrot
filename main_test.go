package main

import (
	"testing"
	"github.com/buffaloluk7/mandelbrot/mandelbrot"
	"time"
	"os"
	"github.com/op/go-logging"
	"sync"
)

func TestMandelbrotPerformance(t *testing.T) {
	var backend = logging.NewLogBackend(os.Stdout, "", 0)
	var backendLeveled = logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backendLeveled)

	specs := mandelbrot.ReadFromFile("data/mb0.spec")
	generator := mandelbrot.NewMandelbrotGenerator(specs)

	numberOfIterations := 10

	for i := 0; i < numberOfIterations; i++ {
		start := time.Now()
		generator.CreateMandelbrot(1)
		log.Infof("#%d: Took %s to create mandelbrot set.", i + 1, time.Since(start))
	}
}

func TestMandelbrotPerformanceParallel(t *testing.T) {
	var backend = logging.NewLogBackend(os.Stdout, "", 0)
	var backendLeveled = logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backendLeveled)

	specs := mandelbrot.ReadFromFile("data/mb0.spec")
	generator := mandelbrot.NewMandelbrotGenerator(specs)

	numberOfIterations := 10

	barrier := &sync.WaitGroup{}
	barrier.Add(numberOfIterations)

	for i := 0; i < numberOfIterations; i++ {
		go func(i int, barrier *sync.WaitGroup) {
			defer barrier.Done()

			start := time.Now()
			generator.CreateMandelbrot(1)
			log.Infof("#%d: Took %s to create mandelbrot set.", i + 1, time.Since(start))
		}(i, barrier)
	}

	barrier.Wait()
}