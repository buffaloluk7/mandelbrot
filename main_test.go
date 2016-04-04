package main

import (
	"testing"
	"github.com/buffaloluk7/mandelbrot/mandelbrot"
	"github.com/buffaloluk7/mandelbrot/specs"
)

func BenchmarkMandelbrot(b *testing.B) {
	specs := specs.ReadFromFile("data/mb0.spec")
	generator := mandelbrot.NewMandelbrotGenerator(specs)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		generator.CreateMandelbrot(1)
	}
}

func BenchmarkMandelbrotParallel(b *testing.B) {
	specs := specs.ReadFromFile("data/mb0.spec")
	generator := mandelbrot.NewMandelbrotGenerator(specs)
	b.ResetTimer()

	b.RunParallel(func (pb *testing.PB) {
		for pb.Next() {
			generator.CreateMandelbrot(1)
		}
	})
}