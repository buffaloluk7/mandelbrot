package mandelbrot

import (
	"os"
	"bufio"
	"strconv"
)

type Specs struct {
	Width, Height, MaximumNumberOfIterations int
	MinR, MinI float64
	MaxR, MaxI float64
}

func NewSpecs(width, height int, minR, minI, maxR, maxI float64, maximumNumberOfIterations int) *Specs {
	return &Specs{
		Width:width,
		Height:height,
		MinR:minR,
		MinI:minI,
		MaxR:maxR,
		MaxI:maxI,
		MaximumNumberOfIterations:maximumNumberOfIterations}
}

func ReadFromFile(filePath string) *Specs {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arguments := make([]string, 7)
	for i := 0; i < 7 && scanner.Scan(); i++ {
		arguments[i] = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return NewSpecs(
		parseInt(arguments[0]),
		parseInt(arguments[1]),
		parseFloat(arguments[2]),
		parseFloat(arguments[3]),
		parseFloat(arguments[4]),
		parseFloat(arguments[5]),
		parseInt(arguments[6]))
}

func parseInt(value string) int {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}

	return int(v)
}

func parseFloat(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	return v
}