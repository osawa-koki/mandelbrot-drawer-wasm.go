package main

import (
	"bytes"
	"encoding/base64"
  "fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"syscall/js"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	iterations             = 200
	contrast               = 15
)

func drawMandelbrot(_this js.Value, _args []js.Value) any {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	bytes := image2bytes(img)
	return bytes2base64(bytes)
}

func image2bytes(img image.Image) []byte {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func bytes2base64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

func mandelbrot(z complex128) color.Color {
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func main() {
  fmt.Println("Program started!")
	js.Global().Set("drawMandelbrot", js.FuncOf(drawMandelbrot))
	fmt.Println("Program finished!")
  select {}
}
