package mandelbrot

import (
	"image"
	"image/color"
	"math"
	"sync"
	"github.com/buffaloluk7/mandelbrot/specs"
)

type MandelbrotGenerator struct {
	specs         *specs.Specs
	imageData     *image.RGBA
}

type Palette []color.RGBA

var palette = func() Palette {
	palette := make([]color.RGBA, 16)

	palette[0] = color.RGBA{66, 30, 15, 0}
	palette[1] = color.RGBA{25, 7, 26, 0}
	palette[2] = color.RGBA{9, 1, 47, 0}
	palette[3] = color.RGBA{4, 4, 73, 0}
	palette[4] = color.RGBA{0, 7, 100, 0}
	palette[5] = color.RGBA{12, 44, 138, 0}
	palette[6] = color.RGBA{24, 82, 177, 0}
	palette[7] = color.RGBA{57, 125, 209, 0}
	palette[8] = color.RGBA{134, 181, 229, 0}
	palette[9] = color.RGBA{211, 236, 248, 0}
	palette[10] = color.RGBA{241, 233, 191, 0}
	palette[11] = color.RGBA{248, 201, 95, 0}
	palette[12] = color.RGBA{255, 170, 0, 0}
	palette[13] = color.RGBA{204, 128, 0, 0}
	palette[14] = color.RGBA{153, 87, 0, 0}
	palette[15] = color.RGBA{106, 52, 3, 0}

	return palette
}()

func NewMandelbrotGenerator(specs *specs.Specs) *MandelbrotGenerator {
	imageData := image.NewRGBA(image.Rect(0, 0, specs.Width - 1, specs.Height - 1))

	return &MandelbrotGenerator{
		specs:specs,
		imageData:imageData}
}

func (g MandelbrotGenerator) CreateMandelbrot(sharpnessFactor int) *image.RGBA {
	// Calculate number of tasks
	numberOfTasks := int(math.Ceil(float64(g.specs.Height) / float64(g.specs.NumberOfLinesPerTask)))
	// Adjust number of tasks depending on sharpnessFactor (all lines inside the sharpnessFactor becomes a single line)
	// e.g. sharpnessFactor = 4, height = 15, numberOfLinesPerTask = 2 --> 2 Tasks
	// e.g. sharpnessFactor = 2, height = 15, numberOfLinesPerTask = 2 --> 4 Tasks
	// e.g. sharpnessFactor = 1, height = 15, numberOfLinesPerTask = 2 --> 8 Tasks
	if sharpnessFactor > 1 {
		numberOfTasks = int(math.Ceil(float64(numberOfTasks) / float64(sharpnessFactor)))
	}

	// Create tasks
	taskChannel := make(chan *Task)
	go g.createTasks(taskChannel, numberOfTasks, sharpnessFactor)

	// Setup barrier (for calculation and processing go routines)
	barrier := &sync.WaitGroup{}
	barrier.Add(numberOfTasks)

	// Process tasks
	go g.calculateMandelbrot(taskChannel, sharpnessFactor, barrier)

	// Wait for all go routines to finish
	barrier.Wait()

	return g.imageData
}

func (g MandelbrotGenerator) createTasks(taskChannel chan <- *Task, numberOfTasks, sharpnessFactor int) {
	numberOfLines := g.specs.NumberOfLinesPerTask * sharpnessFactor

	for i := 0; i < numberOfTasks; i++ {
		startLineIndex := i * numberOfLines

		// Adjust number of lines for last task
		if i == numberOfTasks - 1 && g.specs.Height % numberOfLines != 0 {
			numberOfLines = g.specs.Height % numberOfLines
		}

		taskChannel <- NewTask(startLineIndex, numberOfLines)
	}
}

func (g MandelbrotGenerator) calculateMandelbrot(taskChannel <- chan *Task, sharpnessFactor int, barrier *sync.WaitGroup) {
	scaler := NewCoordinateScaler(g.specs)
	calculator := NewMandelbrotCalculator(g.specs.MaximumNumberOfIterations)

	width := g.specs.Width
	imageSize := g.specs.Height * width

	for {
		task := <-taskChannel

		go func() {
			defer barrier.Done()

			isInitialSharpnessFactor := sharpnessFactor == g.specs.InitialSharpnessFactor
			previousSharpnessFactor := sharpnessFactor * 2

			for y := task.startLineIndex; y < task.startLineIndex + task.numberOfLines; y += sharpnessFactor {
				for x := 0; x < width; x += sharpnessFactor {
					if !isInitialSharpnessFactor &&
					x % previousSharpnessFactor < sharpnessFactor &&
					y % previousSharpnessFactor < sharpnessFactor {
						continue
					}

					real, imaginary := scaler.Scale(x, y)
					value := calculator.FindValue(real, imaginary)

					valueColor := color.RGBA{0, 0, 0, 0}
					if value > 0 && value < g.specs.MaximumNumberOfIterations {
						valueColor = palette[value % 16]
					}

					for innerY := 0; innerY < sharpnessFactor; innerY++ {
						for innerX := 0; innerX < sharpnessFactor; innerX++ {
							index := (y + innerY) * width + x + innerX
							if index < imageSize {
								g.imageData.SetRGBA(x + innerX, y + innerY, valueColor)
							}
						}
					}
				}
			}
		}()
	}
}