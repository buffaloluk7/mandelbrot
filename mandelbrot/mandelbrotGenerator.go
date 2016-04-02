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

type Task struct {
	startLineIndex, numberOfLines int
}

func NewTask(startLineIndex, numberOfLines int) *Task {
	return &Task{startLineIndex:startLineIndex, numberOfLines:numberOfLines}
}

type MandelbrotValue struct {
	value uint8
	x, y int
}

func NewMandelbrotValue(value uint8, x, y int) *MandelbrotValue {
	return &MandelbrotValue{value:value, x:x, y:y}
}

func (g MandelbrotGenerator) CreateImage() *image.RGBA {
	numberOfLinesPerTask := 2
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

	imageData := image.NewRGBA(image.Rect(0, 0, g.specs.Width - 1, g.specs.Height - 1))
	value := make(chan *MandelbrotValue)
	done := make(chan bool)

	go func(imageData *image.RGBA, value <-chan *MandelbrotValue, done <-chan bool) {
		for {
			select {
			case value := <-value:
				r, g, b := value.value, value.value, value.value
				imageData.SetRGBA(value.x, value.y, color.RGBA{R:r, G:g, B:b})
			case <-done:
				return
			}
		}
	}(imageData, value, done)

	calculator := NewMandelbrotCalculator(g.specs.MaximumNumberOfIterations)
	scaler := NewCoordinateScaler(g.specs.Minimum, g.specs.Maximum, g.specs.Width, g.specs.Height)
	barrier := &sync.WaitGroup{}

	for task := range taskChannel {
		barrier.Add(1)

		log.Infof("Start processing task with line index %d", task.startLineIndex)

		go func(task *Task, scaler *CoordinateScaler, calculator *MandelbrotCalculator, value chan <- *MandelbrotValue) {
			defer barrier.Done()

			for y := task.startLineIndex; y < task.startLineIndex + task.numberOfLines; y++ {
				for x := 0; x < g.specs.Width; x++ {
					complexNumber := scaler.Scale(x, y)
					mandelbrotValue := (uint8)(calculator.FindValue(complexNumber))

					value <- NewMandelbrotValue(mandelbrotValue, x, y)
				}
			}


		}(task, scaler, calculator, value)
	}

	barrier.Wait()
	done <- true

	return imageData
}