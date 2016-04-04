package mandelbrot

type MandelbrotCalculator struct {
	maximumNumberOfIterations int
}

func NewMandelbrotCalculator(maximumNumberOfIterations int) *MandelbrotCalculator {
	return &MandelbrotCalculator{maximumNumberOfIterations:maximumNumberOfIterations}
}

func (m MandelbrotCalculator) FindValue(real, imaginary float64) int {
	i := 0
	
	nextI, nextR := imaginary, real

	for ; i < m.maximumNumberOfIterations && !(nextR * nextR + nextI * nextI >= 4.0); i++ {
		real := nextR * nextR - nextI * nextI + real
		imaginary := 2.0 * nextR * nextI + imaginary

		nextI, nextR = imaginary, real
	}

	return i
}