package mandelbrot

type ComplexNumber struct {
	real, imaginary float64
}

func NewComplexNumber(real, imaginary float64) *ComplexNumber {
	return &ComplexNumber{real:real, imaginary: imaginary}
}