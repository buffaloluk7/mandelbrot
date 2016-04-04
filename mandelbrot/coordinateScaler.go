package mandelbrot

type CoordinateScaler struct {
	specs *Specs
}

func NewCoordinateScaler(specs *Specs) *CoordinateScaler {
	return &CoordinateScaler{specs}
}

func (s CoordinateScaler) Scale(x, y int) (float64, float64) {
	cReal := float64(x) * ((s.specs.MaxR - s.specs.MinR) / float64(s.specs.Width)) + s.specs.MinR
	cImaginary := float64(y) * ((s.specs.MaxI - s.specs.MinI) / float64(s.specs.Height)) + s.specs.MinI

	return cReal, cImaginary
}