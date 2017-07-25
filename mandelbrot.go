// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"

	"github.com/lucasb-eyer/go-colorful"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 2048, 2048
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			// Image point (px, py) represents complex value z.
			// Super sample: each corner
			var samples [4]color.Color
			var i = 0
			for sy := 0.; sy < 1.; sy += 0.5 {
				y := (float64(py)+sy)/height*(ymax-ymin) + +ymin
				for sx := 0.; sx < 1.; sx += 0.5 {
					x := (float64(px)+sx)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					//samples[i] = mandelbrot(z)
					samples[i] = newton(z)
					i++
				}
			}
			// Average the samples and set the single pixel
			var r, g, b uint32
			for _, c := range samples {
				_r, _g, _b, _ := c.RGBA()
				r += uint32(_r / 257)
				g += uint32(_g / 257)
				b += uint32(_b / 257)
			}
			r /= 4
			g /= 4
			b /= 4
			img.Set(px, py, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 25

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colorize(v, n)
		}
	}
	return color.Black
}

func colorize(z complex128, n uint8) color.Color {
	hue := float64(n+1) - math.Log(math.Log(cmplx.Abs(z)))/math.Log(2.)
	hue = 0.95 + 20.0*hue
	return colorful.Hsv(hue, 0.8, 1.0)
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 120
	const contrast = 6
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			hue := float64(360-contrast*int(i)) + 270
			// Put hue in correct range
			for ; hue > 360; hue -= 360 {
			}
			return colorful.Hsv(hue, 0.75, 1.0)
		}
	}
	return color.Black
}
