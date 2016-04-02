package mandelbrot

import (
	"os"
	"bufio"
	"strconv"
)

type Specs struct {
	Width, Height, MaximumNumberOfIterations int
	Minimum, Maximum *ComplexNumber
}

func NewSpecs(width, height int, minimum, maximum *ComplexNumber, maximumNumberOfIterations int) *Specs {
	return &Specs{
		Width:width,
		Height:height,
		Minimum:minimum,
		Maximum:maximum,
		MaximumNumberOfIterations:maximumNumberOfIterations}
}

func ReadFromFile(filePath string) *Specs {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arguments := make([]string, 7)
	for i := 0; i < 7 && scanner.Scan(); i++ {
		arguments[i] = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	return NewSpecs(
		parseInt(arguments[0]),
		parseInt(arguments[1]),
		NewComplexNumber(parseFloat(arguments[2]), parseFloat(arguments[3])),
		NewComplexNumber(parseFloat(arguments[4]), parseFloat(arguments[5])),
		parseInt(arguments[6]))
}

func parseInt(value string) int {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	return int(v)
}

func parseFloat(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Panic(err)
	}

	return v
}