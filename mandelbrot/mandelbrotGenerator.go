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
}

func NewMandelbrotGenerator(specs *Specs) *MandelbrotGenerator {
	return &MandelbrotGenerator{specs:specs}
}

func (g MandelbrotGenerator) CreateMandelbrot() *image.RGBA {
	numberOfLinesPerTask := 30
	numberOfTasks := int(math.Ceil(float64(g.specs.Height) / float64(numberOfLinesPerTask)))

	taskChannel := make(chan *Task, numberOfTasks)

	for i := 0; i < numberOfTasks; i++ {
		startLineIndex := i * numberOfLinesPerTask
		numberOfLines := numberOfLinesPerTask
		if i == numberOfTasks - 1 && g.specs.Height % numberOfLinesPerTask != 0 {
			numberOfLines = g.specs.Height % numberOfLinesPerTask
		}

		log.Debugf("Create new task with start line index %d and %d lines to process", startLineIndex, numberOfLines)

		taskChannel <- NewTask(startLineIndex, numberOfLines)
	}

	log.Debug("Close task channel")
	close(taskChannel)

	valuesChannel := make(chan *[]MandelbrotValue)

	taskBarrier := &sync.WaitGroup{}
	taskBarrier.Add(numberOfTasks)
	go calculateMandelbrot(taskChannel, g.specs, valuesChannel, taskBarrier)

	quitBarrier := &sync.WaitGroup{}
	quitBarrier.Add(numberOfTasks)
	imageData := image.NewRGBA(image.Rect(0, 0, g.specs.Width - 1, g.specs.Height - 1))
	go processResults(imageData, valuesChannel, quitBarrier)

	taskBarrier.Wait()
	quitBarrier.Wait()

	return imageData
}

func processResults(imageData *image.RGBA, valuesChannel <- chan *[]MandelbrotValue, quitBarrier *sync.WaitGroup) {
	for {
		values := <-valuesChannel

		go func(values *[]MandelbrotValue, imageData *image.RGBA, quitBarrier *sync.WaitGroup) {
			defer quitBarrier.Done()

			for _, value := range *values {
				r, g, b := value.value, value.value, value.value
				imageData.SetRGBA(value.x, value.y, color.RGBA{R:r, G:g, B:b})
			}
		}(values, imageData, quitBarrier)
	}
}

func calculateMandelbrot(taskChannel <- chan *Task, specs *Specs, valuesChannel chan <- *[]MandelbrotValue, taskBarrier *sync.WaitGroup) {
	calculator := NewMandelbrotCalculator(specs.MaximumNumberOfIterations)
	scaler := NewCoordinateScaler(specs.Minimum, specs.Maximum, specs.Width, specs.Height)

	for task := range taskChannel {
		log.Debug("Start processing task with line index %d", task.startLineIndex)

		go func(task *Task, specs *Specs, scaler *CoordinateScaler, calculator *MandelbrotCalculator, valuesChannel chan <- *[]MandelbrotValue, taskBarrier *sync.WaitGroup) {
			defer taskBarrier.Done()

			values := make([]MandelbrotValue, task.numberOfLines * specs.Width)

			for y := task.startLineIndex; y < task.startLineIndex + task.numberOfLines; y++ {
				for x := 0; x < specs.Width; x++ {
					complexNumber := scaler.Scale(x, y)
					mandelbrotValue := (uint8)(calculator.FindValue(complexNumber))

					index := (y - task.startLineIndex) * specs.Width + x
					values[index] = *NewMandelbrotValue(mandelbrotValue, x, y)
				}
			}

			valuesChannel <- &values
		}(task, specs, scaler, calculator, valuesChannel, taskBarrier)
	}
}