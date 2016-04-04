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
		panic("ListenAndServe: " + err.Error())
	}
}

func mandelbrotHandler(ws *websocket.Conn) {
	msg := make([]byte, 512)
	if _, err := ws.Read(msg); err != nil {
		panic(err)
	}

	specs := specs.ReadFromString(string(msg))
	generator := mandelbrot.NewMandelbrotGenerator(specs)

	initialSharpnessFactor := 8

	for i := 0; i < initialSharpnessFactor; i++ {
		start := time.Now()
		imageData := generator.CreateMandelbrot(initialSharpnessFactor - i)
		fmt.Printf("Took %s to create mandelbrot set.", time.Since(start))

		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, imageData, nil); err != nil {
			panic("unable to encode image.")
		}

		data := base64.StdEncoding.EncodeToString([]byte(buffer.Bytes()))

		_, err := ws.Write([]byte(data))
		if err != nil {
			panic(err)
		}
	}
}