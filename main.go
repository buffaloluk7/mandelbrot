package main

import (
	"github.com/buffaloluk7/mandelbrot/mandelbrot"
	"github.com/op/go-logging"
	"os"
	"image/jpeg"
	"image"
	"image/color"
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
	specs := mandelbrot.ReadFromFile("data/mb1.spec")

	start := time.Now()

	calculator := mandelbrot.NewMandelbrotCalculator(specs.MaximumNumberOfIterations)
	scaler := mandelbrot.NewCoordinateScaler(specs.Minimum, specs.Maximum, specs.Width, specs.Height)

	imageData := image.NewRGBA(image.Rect(0, 0, specs.Width - 1, specs.Height - 1))

	for y := 0; y < specs.Height; y++ {
		for x := 0; x < specs.Width; x++ {
			mandelbrotValue := calculateMandelbrotValue(scaler, calculator, x, y)
			setPixel(imageData, mandelbrotValue, x, y)
		}
	}

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

func calculateMandelbrotValue(scaler *mandelbrot.CoordinateScaler, calculator *mandelbrot.MandelbrotCalculator, x, y int) uint8 {
	complexNumber := scaler.Scale(x, y)
	return (uint8)(calculator.FindValue(complexNumber))
}

func setPixel(imageData *image.RGBA, mandelbrotValue uint8, x, y int) {
	// Use smooth polynomials for r, g, b
	//t := mandelbrotValue / specs.MaximumNumberOfIterations
	//r := (uint8)(9 * (1 - t) * t * t * t * 255)
	//g := (uint8)(15 * (1 - t) * (1 - t) * t * t * 255)
	//b := (uint8)(9 * (1 - t) * (1 - t) * (1 - t) * t * 255)
	r, g, b := mandelbrotValue, mandelbrotValue, mandelbrotValue
	imageData.SetRGBA(x, y, color.RGBA{R:r, G:g, B:b})
}