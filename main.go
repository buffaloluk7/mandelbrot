package main

import (
	"github.com/buffaloluk7/mandelbrot/mandelbrot"
	"github.com/buffaloluk7/mandelbrot/specs"
	"image/jpeg"
	"time"
	"golang.org/x/net/websocket"
	"net/http"
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {
	http.Handle("/mandelbrot", websocket.Handler(mandelbrotHandler))
	http.Handle("/", http.FileServer(http.Dir("client")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to open socket on port 8080 with error. %s.\n", err.Error())
	}
}

func mandelbrotHandler(ws *websocket.Conn) {
	msg := make([]byte, 512)
	if _, err := ws.Read(msg); err != nil {
		fmt.Printf("Failed to read message with error: %s.\n", err.Error())
	}

	specs := specs.ReadFromString(string(msg))
	generator := mandelbrot.NewMandelbrotGenerator(specs)
	imageSize := specs.Height * specs.Width

	for sharpnessFactor := specs.InitialSharpnessFactor; sharpnessFactor > 0; sharpnessFactor /= 2 {
		calculateMandelbrot(ws, generator, sharpnessFactor, imageSize)
	}
}

func calculateMandelbrot(ws *websocket.Conn, generator *mandelbrot.MandelbrotGenerator, sharpnessFactor, imageSize int) {
	start := time.Now()
	imageData := generator.CreateMandelbrot(sharpnessFactor)
	fmt.Printf("Took %s to create mandelbrot set with sharpness factor %d.\n", time.Since(start), sharpnessFactor)

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, imageData, nil); err != nil {
		fmt.Printf("Failed to encode image with error: %s.\n", err.Error())
	}

	output := make([]byte, imageSize)
	base64.StdEncoding.Encode(output, buffer.Bytes())

	if _, err := ws.Write(output); err != nil {
		fmt.Printf("Failed to send message with error: %s.\n", err.Error())
	}
}