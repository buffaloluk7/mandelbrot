package specs

import (
	"os"
	"bufio"
	"strconv"
	"strings"
)

type Specs struct {
	Width, Height             int
	MinR, MinI                float64
	MaxR, MaxI                float64
	MaximumNumberOfIterations int
	NumberOfLinesPerTask      int
	InitialSharpnessFactor    int
}

func NewSpecs(width, height int, minR, minI, maxR, maxI float64, maximumNumberOfIterations, maximumNumberOfLinesPerTask, initialSharpnessFactor int) *Specs {
	return &Specs{
		Width:width,
		Height:height,
		MinR:minR,
		MinI:minI,
		MaxR:maxR,
		MaxI:maxI,
		MaximumNumberOfIterations:maximumNumberOfIterations,
		NumberOfLinesPerTask:maximumNumberOfLinesPerTask,
		InitialSharpnessFactor:initialSharpnessFactor}
}

func ReadFromString(specs string) *Specs {
	return fromString(specs)
}

func ReadFromFile(filePath string) *Specs {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arguments string
	for scanner.Scan() {
		arguments = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return fromString(arguments)
}

func fromString(arguments string) *Specs {
	splittedArguments := strings.Split(arguments, ";")

	return NewSpecs(
		(int)(parseFloat(splittedArguments[0])),
		(int)(parseFloat(splittedArguments[1])),
		parseFloat(splittedArguments[2]),
		parseFloat(splittedArguments[3]),
		parseFloat(splittedArguments[4]),
		parseFloat(splittedArguments[5]),
		(int)(parseFloat(splittedArguments[6])),
		5,
		8)
}

func parseFloat(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	return v
}