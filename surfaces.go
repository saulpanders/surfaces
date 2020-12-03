//@saulpanders
// need to track down the reference for this... (sorry!)
// computes SVG rendering of a 3-D Manifold (surface)
//TODO: compute colors based on height (i.e. red valleys blue hils)
//TODO: add cmd-line args, add equation parsing
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            //canvas size (px)
	cells         = 100                 //# of grid cells
	xyrange       = 2.0                 //axis ranges (+-xyrange)
	xyscale       = width / 2 / xyrange //px per x or y unit
	zscale        = height * 0.4        //px per z unit
	angle         = math.Pi / 6         //angle of x, y axes
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30), cos(30)

func main() {
	printSurface()
}

func corner(i, j int) (float64, float64) {
	//find (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//compute height
	z := f(x, y)

	//Project (x,y,z) isometrically onto 2d SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	//ripple sin(r)/r
	//r := math.Hypot(x, y) //dist from (0,0)
	//ripple := math.Sin(r) / r
	// ^ change xyrange to 30

	//monkey saddle, change xyrange to 2.0
	//saddle := math.Pow(x, 3) - 3*x*math.Pow(y, 2)
	//return saddle
	pitchfork := math.Pow(x, 3) - 3*x
	return pitchfork
}

func printSurface() {
	fmt.Printf("<svg xmlns = 'http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width= '%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Printf("</svg>")
}
