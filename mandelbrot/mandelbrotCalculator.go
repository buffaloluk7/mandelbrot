package mandelbrot

type MandelbrotCalculator struct {
	maximumNumberOfIterations int
}

func NewMandelbrotCalculator(maximumNumberOfIterations int) *MandelbrotCalculator {
	return &MandelbrotCalculator{maximumNumberOfIterations:maximumNumberOfIterations}
}

func (m MandelbrotCalculator) FindValue(complexNumber *ComplexNumber) int {
	i := 0
	cNext := NewComplexNumber(0, 0)

	for ; i < m.maximumNumberOfIterations && !doesEscape(cNext); i++ {
		real := cNext.real * cNext.real - cNext.imaginary * cNext.imaginary + complexNumber.real
		imaginary := 2.0 * cNext.real * cNext.imaginary + complexNumber.imaginary

		cNext = NewComplexNumber(real, imaginary)
	}

	return i
}

func doesEscape(complexNumber *ComplexNumber) bool {
	realPart := complexNumber.real
	imaginaryPart := complexNumber.imaginary

	return realPart * realPart + imaginaryPart * imaginaryPart >= 4.0
}