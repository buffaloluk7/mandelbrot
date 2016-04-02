package main

import (
	"testing"
	"github.com/buffaloluk7/mandelbrot/mandelbrot"
	"time"
	"os"
	"github.com/op/go-logging"
)

func TestMandelbrotPerformance(t *testing.T) {
	var backend = logging.NewLogBackend(os.Stdout, "", 0)
	var backendLeveled = logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backendLeveled)

	specs := mandelbrot.ReadFromFile("data/mb0.spec")
	generator := mandelbrot.NewMandelbrotGenerator(specs)

	for i := 0; i < 10; i++ {
		start := time.Now()
		generator.CreateMandelbrot()
		log.Infof("#%d: Took %s to create mandelbrot set.", i + 1, time.Since(start))
	}
}