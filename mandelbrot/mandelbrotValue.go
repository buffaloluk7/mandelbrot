package mandelbrot

type MandelbrotValue struct {
	value uint8
	x, y  int
}

func NewMandelbrotValue(value uint8, x, y int) *MandelbrotValue {
	return &MandelbrotValue{value:value, x:x, y:y}
}