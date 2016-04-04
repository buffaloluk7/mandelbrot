package mandelbrot

import (
	"image"
	"image/color"
)

type MandelbrotGenerator struct {
	specs *Specs
	scaler *CoordinateScaler
	calculator *MandelbrotCalculator

}

func NewMandelbrotGenerator(specs *Specs) *MandelbrotGenerator {
	return &MandelbrotGenerator{specs:specs}
}

func (g MandelbrotGenerator) CreateMandelbrot() *image.RGBA {

	g.scaler = NewCoordinateScaler(g.specs)
	g.calculator = NewMandelbrotCalculator(g.specs.MaximumNumberOfIterations)

	values := g.pointGenerator()

	image := image.NewRGBA(image.Rect(0, 0, g.specs.Width - 1, g.specs.Height - 1))

	for value := range values {
		image.SetRGBA(value.x, value.y, color.RGBA{R:value.value, G:value.value, B:value.value})
	}

	return image
}

func (g MandelbrotGenerator) pointGenerator() <-chan MandelbrotValue {

	mandelbrotChannel := make(chan MandelbrotValue, 500)

	go func() {
		defer close(mandelbrotChannel)

		for y := 0; y < g.specs.Height; y++ {
			for x := 0; x < g.specs.Width; x++ {
				r, i := g.scaler.Scale(x, y)
				mandelbrotValue := (uint8)(g.calculator.FindValue(r, i))

				mandelbrotChannel <- MandelbrotValue{mandelbrotValue, x, y}
			}
		}
	}()

	return mandelbrotChannel
}