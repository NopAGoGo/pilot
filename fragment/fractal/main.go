package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	xMin, yMin, xMax, yMax = -2, -2, +2, +2
	width, height          = 1 << 13, 1 << 13
)

func main() {
	// http://localhost:7070/?cycles=10
	handler := func(w http.ResponseWriter, r *http.Request) {
		cycles, err := strconv.Atoi(r.FormValue("cycles"))
		if err != nil {
			cycles = 5
		}
		lissajous(w, cycles)
	}
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":7070", nil)
	fmt.Println(err)
}

func lissajous(out io.Writer, cycles int) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(yMax-yMin) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xMax-xMin) + xMin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(out, img) // NOTE: ignoring errors

}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
