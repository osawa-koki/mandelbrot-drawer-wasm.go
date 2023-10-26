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

func drawMandelbrot(_this js.Value, _args []js.Value) any {
	xmin := _args[0].Float()
	ymin := _args[1].Float()
	xmax := _args[2].Float()
	ymax := _args[3].Float()
	width := _args[4].Int()
	height := _args[5].Int()
	iterations := _args[6].Int()
	contrast := _args[7].Int()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, (func() color.Color {
				var v complex128
				for n := int(0); n < iterations; n++ {
					v = v*v + z
					if cmplx.Abs(v) > 2 {
						return color.Gray{255 - uint8(contrast*n)}
					}
				}
				return color.Black
			})())
		}
	}

	return bytes2base64(image2bytes(img))
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

func main() {
  fmt.Println("Program started!")
	js.Global().Set("drawMandelbrot", js.FuncOf(drawMandelbrot))
	fmt.Println("Program finished!")
  select {}
}
