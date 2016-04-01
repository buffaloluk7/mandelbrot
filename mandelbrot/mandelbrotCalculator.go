package mandelbrot

type MandelbrotCalculator struct {
	maximumNumberOfIterations int
}

func NewMandelbrotCalculator(maximumNumberOfIterations int) *MandelbrotCalculator {
	return &MandelbrotCalculator{maximumNumberOfIterations:maximumNumberOfIterations}
}

func (m MandelbrotCalculator) FindValue(complexNumber *ComplexNumber) int {
	i := 0
	cNext := &ComplexNumber{real:0, imaginary:0}

	for ; i < m.maximumNumberOfIterations && !doesEscape(cNext); i++ {
		real := cNext.real * cNext.real - cNext.imaginary * cNext.imaginary + complexNumber.real
		imaginary := 2.0 * cNext.real * cNext.imaginary + complexNumber.imaginary

		cNext = &ComplexNumber{real:real, imaginary: imaginary}
	}

	return i
}

func doesEscape(complexNumber *ComplexNumber) bool {
	realPart := complexNumber.real
	imaginaryPart := complexNumber.imaginary

	return realPart * realPart + imaginaryPart * imaginaryPart >= 4.0
}