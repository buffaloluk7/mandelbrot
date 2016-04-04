package main

import (
	"github.com/buffaloluk7/mandelbrot/mandelbrot"
	"github.com/op/go-logging"
	"os"
	"image/jpeg"
	"time"
	"golang.org/x/net/websocket"
	"net/http"
	"bytes"
	"encoding/base64"
)

var log = logging.MustGetLogger("main")

func main() {
	// Setup logging
	var backend = logging.NewLogBackend(os.Stdout, "", 0)
	var backendLeveled = logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backendLeveled)

	// Parse console arguments
	/*if len(os.Args) != 2 {
		log.Panic("Invalid number of arguments. Expected: 1")
		return
	}

	specs := mandelbrot.ReadFromFile(os.Args[1])*/

	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir("client")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func echoHandler(ws *websocket.Conn) {
	msg := make([]byte, 512)
	if _, err := ws.Read(msg); err != nil {
		log.Fatal(err)
	}

	specs := mandelbrot.ReadFromFile("data/mb0.spec")
	generator := mandelbrot.NewMandelbrotGenerator(specs)

	initialSharpnessFactor := 8

	for i := 0; i < initialSharpnessFactor; i++ {
		start := time.Now()
		imageData := generator.CreateMandelbrot(initialSharpnessFactor - i)
		log.Infof("Took %s to create mandelbrot set.", time.Since(start))

		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, imageData, nil); err != nil {
			log.Debug("unable to encode image.")
		}

		data := base64.StdEncoding.EncodeToString([]byte(buffer.Bytes()))

		m, err := ws.Write([]byte(data))
		if err != nil {
			log.Fatal(err)
		}
		log.Debug("Send: %d", m)
	}
}