package mandelbrot

import (
	"image"
	"image/color"
	"math"
	"sync"
)

type MandelbrotGenerator struct {
	specs *Specs
	numberOfTasks, numberOfLinesPerTask int
}

func NewMandelbrotGenerator(specs *Specs) *MandelbrotGenerator {
	numberOfLinesPerTask := 30
	numberOfTasks := int(math.Ceil(float64(specs.Height) / float64(numberOfLinesPerTask)))

	return &MandelbrotGenerator{
		specs:specs,
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
	imageData := image.NewRGBA(image.Rect(0, 0, g.specs.Width - 1, g.specs.Height - 1))
	go g.calculateMandelbrot(taskChannel, imageData, barrier, sharpnessFactor)

	// Wait for all go routines to finish
	barrier.Wait()

	return imageData
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

func (g MandelbrotGenerator) calculateMandelbrot(taskChannel <- chan *Task, imageData *image.RGBA, barrier *sync.WaitGroup, sharpnessFactor int) {
	scaler := NewCoordinateScaler(g.specs)
	calculator := NewMandelbrotCalculator(g.specs.MaximumNumberOfIterations)

	for {
		task := <-taskChannel

		go func(task *Task, imageData *image.RGBA, width, height int, scaler *CoordinateScaler, calculator *MandelbrotCalculator, barrier *sync.WaitGroup, sharpnessFactor int) {
			defer barrier.Done()

			length := task.numberOfLines * width

			for y := task.startLineIndex; y < task.startLineIndex + task.numberOfLines; y = y + sharpnessFactor {
				for x := 0; x < width; x = x + sharpnessFactor {
					r, i := scaler.Scale(x, y)
					mandelbrotValue := (uint8)(calculator.FindValue(r, i))

					for innerY := 0; innerY < sharpnessFactor; innerY++ {
						for innerX := 0; innerX < sharpnessFactor; innerX++ {
							index := ((y - task.startLineIndex + innerY) * width) + x + innerX
							if index < length {
								r, g, b := mandelbrotValue, mandelbrotValue, mandelbrotValue
								imageData.SetRGBA(x + innerX, y + innerY, color.RGBA{R:r, G:g, B:b})
							}
						}
					}
				}
			}
		}(task, imageData, g.specs.Width, g.specs.Height, scaler, calculator, barrier, sharpnessFactor)
	}
}