// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	zmax          = 1                   // high value of z where colour saturates
	zmin          = -1                  // low value of z where colour saturates
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func writeSvg(w io.Writer, height, width int) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, color, oka := corner(i+1, j, height, width)
			bx, by, _, okb := corner(i, j, height, width)
			cx, cy, _, okc := corner(i, j+1, height, width)
			dx, dy, _, okd := corner(i+1, j+1, height, width)
			if oka && okb && okc && okd {
				fmt.Fprintf(w, "<polygon fill='#%02X%02X%02X' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					color.R, color.G, color.B, ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j, height, width int) (float64, float64, color.RGBA, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale

	// Figure out the colour of the point
	z = math.Min(zmax, math.Max(zmin, z))
	redc := uint8((z - zmin) / (zmax - zmin) * 255)
	bluec := uint8((-z - zmin) / (zmax - zmin) * 255)

	return sx, sy, color.RGBA{redc, 0, bluec, 0xff}, !math.IsInf(z, 0)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var err error
		var myheight int
		var mywidth int
		r.ParseForm()
		w.Header().Set("Content-Type", "image/svg+xml")
		heightarg, ok := r.Form["height"]
		if !ok {
			myheight = height
		} else {
			myheight, err = strconv.Atoi(heightarg[0])
			if err != nil {
				myheight = height
			}
		}
		weightarg, ok := r.Form["width"]
		if !ok {
			mywidth = width
		} else {
			mywidth, err = strconv.Atoi(weightarg[0])
			if err != nil {
				mywidth = width
			}
		}
		writeSvg(w, myheight, mywidth)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-
