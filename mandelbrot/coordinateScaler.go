package mandelbrot

import "github.com/buffaloluk7/mandelbrot/specs"

type CoordinateScaler struct {
	specs *specs.Specs
}

func NewCoordinateScaler(specs *specs.Specs) *CoordinateScaler {
	return &CoordinateScaler{specs}
}

func (s CoordinateScaler) Scale(x, y int) (float64, float64) {
	realRange := s.specs.MaxR - s.specs.MinR
	cReal := float64(x) * (realRange / float64(s.specs.Width)) + s.specs.MinR

	imaginaryRange := s.specs.MaxI - s.specs.MinI
	cImaginary := float64(y) * (imaginaryRange / float64(s.specs.Height)) + s.specs.MinI

	return cReal, cImaginary
}