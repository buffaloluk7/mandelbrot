package main

import (
	"github.com/buffaloluk7/mandelbrot/mandelbrot"
	"github.com/op/go-logging"
	"os"
	"image/jpeg"
	"time"
	"flag"
	"runtime/pprof"
)

var log = logging.MustGetLogger("main")

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

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

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

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