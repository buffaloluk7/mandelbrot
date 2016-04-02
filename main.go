package main

import (
	"github.com/buffaloluk7/mandelbrot/mandelbrot"
	"github.com/op/go-logging"
	"os"
	"image/jpeg"
	"time"
)

var log = logging.MustGetLogger("main")

func main() {
	// Setup logging
	var backend = logging.NewLogBackend(os.Stdout, "", 0)
	var backendLeveled = logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backendLeveled)

	// Parse console arguments
	/*if len(os.Args) != 2 {
		log.Panic("Invalid number of arguments. Expected: 1")
		return
	}

	specs := mandelbrot.ReadFromFile(os.Args[1])*/
	specs := mandelbrot.ReadFromFile("data/mb0.spec")
	generator := mandelbrot.NewMandelbrotGenerator(specs)

	start := time.Now()
	imageData := generator.CreateMandelbrot()
	log.Infof("Took %s to create mandelbrot set.", time.Since(start))

	file, err := os.Create("output.jpg")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	if err := jpeg.Encode(file, imageData, &jpeg.Options{jpeg.DefaultQuality}); err != nil {
		log.Panic("Unable to encode image.")
	}
}