package mandelbrot

import (
	"image"
	"image/color"
	"math"
	"sync"
	"github.com/buffaloluk7/mandelbrot/specs"
)

type MandelbrotGenerator struct {
	specs *specs.Specs
	imageData *image.RGBA
	numberOfTasks, numberOfLinesPerTask int
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
	numberOfLinesPerTask := 30
	numberOfTasks := int(math.Ceil(float64(specs.Height) / float64(numberOfLinesPerTask)))

	return &MandelbrotGenerator{
		specs:specs,
		imageData:imageData,
		numberOfTasks:numberOfTasks,
		numberOfLinesPerTask:numberOfLinesPerTask}
}

func (g MandelbrotGenerator) CreateMandelbrot(sharpnessFactor int) *image.RGBA {
	// Create tasks
	taskChannel := make(chan *Task)
	go g.createTasks(taskChannel, g.numberOfTasks)

	// Setup barrier (for calculation and processing go routines)
	barrier := &sync.WaitGroup{}
	barrier.Add(g.numberOfTasks)

	// Process tasks
	go g.calculateMandelbrot(taskChannel, barrier, sharpnessFactor)

	// Wait for all go routines to finish
	barrier.Wait()

	return g.imageData
}

func (g MandelbrotGenerator) createTasks(taskChannel chan <- *Task, numberOfTasks int) {
	for i := 0; i < numberOfTasks; i++ {
		startLineIndex := i * g.numberOfLinesPerTask
		numberOfLines := g.numberOfLinesPerTask
		if i == numberOfTasks - 1 && g.specs.Height % g.numberOfLinesPerTask != 0 {
			numberOfLines = g.specs.Height % g.numberOfLinesPerTask
		}

		taskChannel <- NewTask(startLineIndex, numberOfLines)
	}
}

func (g MandelbrotGenerator) calculateMandelbrot(taskChannel <- chan *Task, barrier *sync.WaitGroup, sharpnessFactor int) {
	scaler := NewCoordinateScaler(g.specs)
	calculator := NewMandelbrotCalculator(g.specs.MaximumNumberOfIterations)

	for {
		task := <-taskChannel

		go func() {
			defer barrier.Done()

			width := g.specs.Width
			numberOfPoints := task.numberOfLines * width

			for y := task.startLineIndex; y < task.startLineIndex + task.numberOfLines; y = y + sharpnessFactor {
				for x := 0; x < width; x = x + sharpnessFactor {
					real, imaginary := scaler.Scale(x, y)
					value := calculator.FindValue(real, imaginary)

					valueColor := color.RGBA{0, 0, 0, 0}
					if value > 0 && value < g.specs.MaximumNumberOfIterations {
						valueColor = palette[value % 16]
					}

					for innerY := 0; innerY < sharpnessFactor; innerY++ {
						for innerX := 0; innerX < sharpnessFactor; innerX++ {
							index := ((y - task.startLineIndex + innerY) * width) + x + innerX
							if index < numberOfPoints {
								g.imageData.SetRGBA(x + innerX, y + innerY, valueColor)
							}
						}
					}
				}
			}
		}()
	}
}