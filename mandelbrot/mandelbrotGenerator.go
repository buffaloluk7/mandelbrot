package mandelbrot

import (
	"image"
	"image/color"
	"math"
	"github.com/op/go-logging"
	"sync"
)

var log = logging.MustGetLogger("mandelbrotGenerator")

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

func (g MandelbrotGenerator) CreateMandelbrot() *image.RGBA {
	// Create tasks
	taskChannel := make(chan *Task)
	go g.createTasks(taskChannel, g.numberOfTasks)

	// Setup barrier (for calculation and processing go routines)
	barrier := &sync.WaitGroup{}
	barrier.Add(g.numberOfTasks * 2)

	// Process tasks
	valuesChannel := make(chan *[]MandelbrotValue)
	go g.calculateMandelbrot(taskChannel, valuesChannel, barrier)

	// Merge task results
	imageData := image.NewRGBA(image.Rect(0, 0, g.specs.Width - 1, g.specs.Height - 1))
	go g.processResults(imageData, valuesChannel, barrier)

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

		log.Debugf("Create new task with start line index %d and %d lines to process", startLineIndex, numberOfLines)
		taskChannel <- NewTask(startLineIndex, numberOfLines)
	}
}

func (g MandelbrotGenerator) calculateMandelbrot(taskChannel <- chan *Task, valuesChannel chan <- *[]MandelbrotValue, barrier *sync.WaitGroup) {
	scaler := NewCoordinateScaler(g.specs.Minimum, g.specs.Maximum, g.specs.Width, g.specs.Height)
	calculator := NewMandelbrotCalculator(g.specs.MaximumNumberOfIterations)

	for {
		task := <-taskChannel
		log.Debugf("Start processing task with line index %d", task.startLineIndex)

		go func(task *Task, width int, scaler *CoordinateScaler, calculator *MandelbrotCalculator, valuesChannel chan <- *[]MandelbrotValue, barrier *sync.WaitGroup) {
			defer barrier.Done()

			values := make([]MandelbrotValue, task.numberOfLines * width)

			for y := task.startLineIndex; y < task.startLineIndex + task.numberOfLines; y++ {
				for x := 0; x < width; x++ {
					complexNumber := scaler.Scale(x, y)
					mandelbrotValue := (uint8)(calculator.FindValue(complexNumber))

					index := (y - task.startLineIndex) * width + x
					values[index] = *NewMandelbrotValue(mandelbrotValue, x, y)
				}
			}

			valuesChannel <- &values
		}(task, g.specs.Width, scaler, calculator, valuesChannel, barrier)
	}
}

func (g MandelbrotGenerator) processResults(imageData *image.RGBA, valuesChannel <- chan *[]MandelbrotValue, barrier *sync.WaitGroup) {
	for {
		values := <-valuesChannel
		log.Debug("Add calculated mandelbrot values to image")

		go func(values *[]MandelbrotValue, imageData *image.RGBA, barrier *sync.WaitGroup) {
			defer barrier.Done()

			for _, value := range *values {
				r, g, b := value.value, value.value, value.value
				imageData.SetRGBA(value.x, value.y, color.RGBA{R:r, G:g, B:b})
			}
		}(values, imageData, barrier)
	}
}