package mandelbrot

type CoordinateScaler struct {
	minimum, maximum *ComplexNumber
	maxX, maxY       int
}

func NewCoordinateScaler(minimum, maximum *ComplexNumber, maxX, maxY int) *CoordinateScaler {
	return &CoordinateScaler{
		minimum:minimum,
		maximum:maximum,
		maxX:maxX,
		maxY:maxY}
}

func (s CoordinateScaler) Scale(x, y int) *ComplexNumber {
	realRange := s.maximum.real - s.minimum.real
	cReal := float64(x) * (realRange / float64(s.maxX)) + s.minimum.real

	imaginaryRange := s.maximum.imaginary - s.minimum.imaginary
	cImaginary := float64(y) * (imaginaryRange / float64(s.maxY)) + s.minimum.imaginary

	return &ComplexNumber{real:cReal, imaginary:cImaginary}
}