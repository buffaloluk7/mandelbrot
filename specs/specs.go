package specs

import (
	"os"
	"bufio"
	"strconv"
	"io"
	"strings"
)

type Specs struct {
	Width, Height, MaximumNumberOfIterations, InitialSharpnessFactor int
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
		InitialSharpnessFactor:8,
		MaximumNumberOfIterations:maximumNumberOfIterations}
}

func ReadFromString(specs string) *Specs {
	r := strings.NewReader(specs)

	return fromReader(r)
}

func ReadFromFile(filePath string) *Specs {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return fromReader(file)
}

func fromReader(r io.Reader) *Specs{
	scanner := bufio.NewScanner(r)
	var arguments []string
	for scanner.Scan() {
		arguments =  strings.Split(scanner.Text(), ";")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return NewSpecs(
		(int)(parseFloat(arguments[0])),
		(int)(parseFloat(arguments[1])),
		parseFloat(arguments[2]),
		parseFloat(arguments[3]),
		parseFloat(arguments[4]),
		parseFloat(arguments[5]),
		(int)(parseFloat(arguments[6])))
}

func parseFloat(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	return v
}