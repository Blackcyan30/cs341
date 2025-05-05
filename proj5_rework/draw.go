// Author: M.Talha Adnan
// NetID: mkhan387
// Class: CS341 - Spring 2025

package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sync"
)

// ----------------------------------------------------------------------
//
// Define Color type
// depending on your implementation
// Color struct to store the values of color and name
type Color struct {
	name    string
	r, g, b uint
}

// Num of total worker routines
const workers = 4

// Built to retieve colors by name.
var colorMap = map[string]Color{
	"red":    {"red", 255, 0, 0},
	"green":  {"green", 0, 255, 0},
	"blue":   {"blue", 0, 0, 255},
	"yellow": {"yellow", 255, 255, 0},
	"orange": {"orange", 255, 164, 0},
	"purple": {"purple", 128, 0, 128},
	"brown":  {"brown", 165, 42, 42},
	"black":  {"black", 0, 0, 0},
	"white":  {"white", 255, 255, 255},
}

// Point struct
//
//	x: the x position of the point
//	y: the y position of the point
type Point struct {
	x, y int
}

// Rectangle struct
//
//	ll: a Point representing the lower-left corner of the rectangle
//	ur: a Point representing the upper-right corner of the rectangle
//	c: the color of the rectangle
type Rectangle struct {
	ll, ur Point
	c      Color
}

// Triangle struct
//
//	pt0: a Point representing the first point of the triangle
//	pt1: a Point representing the second point of the triangle
//	pt2: a Point representing the third point of the triangle
//	c: the color of the triangle
type Triangle struct {
	pt0, pt1, pt2 Point
	c             Color
}

// Circle struct
//
//	center: a Point representing the center of the circle
//	r: an integer representing the radius of the circle
//	c: the color of the circle
type Circle struct {
	center Point
	r      int
	c      Color
}

// Display struct
//
//	maxX: the dimension of the display on the X-axis
//	maxY: the dimension of the display on the Y-axis
//	matrix: a slice of slices in which each element represents
//			the current color on the display for the given pixel
type Display struct {
	maxX, maxY int
	matrix     [][]Color
}

// Geometry interface
// Each type that implements this interface should implement the following methods:
//
//	draw() - Draw the filled-in shape on the screen
//	printShape() - Print the type of the object
type geometry interface {
	draw(scn screen) (err error)
	printShape() (s string)
}

// Screen interface
// Each type that implements this interface should implement the following methods:
//
//	initialize() - 	Initialize a screen of maxX times maxY
//				   	Called before any other method in screen interface
//				   	Clears the screen so that it is all white
//	getMaxXY() - 	Retrieve and return the maxX and maxY dimensions of the screen
//	drawPixel() -	Draw the pixel with a given color at a given location
//	getPixel() - 	Retrieve and return the color of the pixel at a given location
//	clearScreen() - Clear the whole screen by setting each pixel to white
//	screenShot() - 	Write the screen to a .ppm file
//					(which you can then view graphically with an image viewer)
type screen interface {
	initialize(x, y int)
	getMaxXY() (x, y int)
	drawPixel(x, y int, c Color) (err error)
	getPixel(x, y int) (c Color, err error)
	clearScreen()
	screenShot(f string) (err error)
}

// Package level variables for errors
//
//	outOfBoundsErr - Thrown when the figure is out of bounds of the screen
//	colorUnknownErr - Thrown when a Color is not valid
var outOfBoundsErr error
var colorUnknownErr error

// ----------------------------------------------------------------------
//
// init()
// Called before main(), sets the package level variables for errors
func init() {
	outOfBoundsErr = errors.New("**Error: Attempt to draw a figure out of bounds of the screen.")
	colorUnknownErr = errors.New("**Error: Attempt to use an invalid color.")
}

// ----------------------------------------------------------------------
//
// outOfBounds()
// Check if a given point would go out of bounds of the screen (return true)
// or not (return false)
func outOfBounds(p Point, scn screen) bool {
	x, y := scn.getMaxXY()
	return p.x < 0 || p.y < 0 || p.x >= x || p.y >= y
}

// ----------------------------------------------------------------------
//
// colorUnknown()
// Check if a given color is unknown (return true) or known (return false)
func colorUnknown(c Color) bool {
	_, ok := colorMap[c.name]
	return !ok
}

// ----------------------------------------------------------------------
//
// draw() method with Rectangle receiver
// Draws a filled in rectangle
func (rect Rectangle) draw(scn screen) (err error) {
	var ll, ur, ul, lr Point
	ll, ur = rect.ll, rect.ur
	ul, lr = Point{ll.x, ur.y}, Point{ur.x, ll.y}

	if outOfBounds(ll, scn) ||
		outOfBounds(ur, scn) ||
		outOfBounds(ul, scn) ||
		outOfBounds(lr, scn) {
		return outOfBoundsErr
	}
	if colorUnknown(rect.c) {
		return colorUnknownErr
	}

	for row := ll.y; row < ur.y; row++ {
		for col := ll.x; col < ur.x; col++ {
			err := scn.drawPixel(col, row, rect.c)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ----------------------------------------------------------------------
//
// printShape() method with Rectangle receiver
// Prints the type (a Rectangle)
func (rect Rectangle) printShape() string {
	var ll, ur Point = rect.ll, rect.ur
	return fmt.Sprintf("Rectangle: (%d,%d) to (%d,%d)",
		ll.x, ll.y, ur.x, ur.y)
}

// ----------------------------------------------------------------------
//
// draw() method with Triangle receiver
// Draws a filled in triangle
// Reference: https://gabrielgambetta.com/computer-graphics-from-scratch/07-filled-triangles.html
//
// interpolate() is a helper function
func interpolate(l0, d0, l1, d1 int) (values []int) {
	a := float64(d1-d0) / float64(l1-l0)
	d := float64(d0)

	count := l1 - l0 + 1
	for ; count > 0; count-- {
		values = append(values, int(d))
		d = d + a
	}
	return
}
func (tri Triangle) draw(scn screen) (err error) {
	// Check if drawing this triangle would cause either error
	if outOfBounds(tri.pt0, scn) || outOfBounds(tri.pt1, scn) || outOfBounds(tri.pt2, scn) {
		return outOfBoundsErr
	}
	if colorUnknown(tri.c) {
		return colorUnknownErr
	}

	y0 := tri.pt0.y
	y1 := tri.pt1.y
	y2 := tri.pt2.y

	// Sort the points so that y0 <= y1 <= y2
	if y1 < y0 {
		tri.pt1, tri.pt0 = tri.pt0, tri.pt1
	}
	if y2 < y0 {
		tri.pt2, tri.pt0 = tri.pt0, tri.pt2
	}
	if y2 < y1 {
		tri.pt2, tri.pt1 = tri.pt1, tri.pt2
	}
	x0, y0, x1, y1, x2, y2 := tri.pt0.x, tri.pt0.y, tri.pt1.x, tri.pt1.y, tri.pt2.x, tri.pt2.y

	x01 := interpolate(y0, x0, y1, x1)
	x12 := interpolate(y1, x1, y2, x2)
	x02 := interpolate(y0, x0, y2, x2)

	// Concatenate the short sides
	x012 := append(x01[:len(x01)-1], x12...)

	// Determine which is left and which is right
	var x_left, x_right []int
	m := len(x012) / 2
	if x02[m] < x012[m] {
		x_left = x02
		x_right = x012
	} else {
		x_left = x012
		x_right = x02
	}

	// Draw the horizontal segments
	for y := y0; y <= y2; y++ {
		for x := x_left[y-y0]; x <= x_right[y-y0]; x++ {
			scn.drawPixel(x, y, tri.c)
		}
	}
	return
}

// ----------------------------------------------------------------------
//
// printShape() method with Triangle receiver
// Returns the type (a Triangle) and it's specs
func (tri Triangle) printShape() string {
	var x1, x2, x3 int
	var y1, y2, y3 int

	x1, x2, x3 = tri.pt0.x, tri.pt1.x, tri.pt2.x
	y1, y2, y3 = tri.pt0.y, tri.pt1.y, tri.pt2.y

	return fmt.Sprintf("Triangle: (%d,%d), (%d,%d), (%d,%d)",
		x1, y1, x2, y2, x3, y3)
}

// ----------------------------------------------------------------------
//
// draw() method with Circle receiver
// Draws a filled in circle
// Reference: https://www.redblobgames.com/grids/circle-drawing/
//
// insideCircle() is a helper function
func insideCircle(center, tile Point, r float64) (inside bool) {
	var dx float64 = float64(center.x - tile.x)
	var dy float64 = float64(center.y - tile.y)
	var distance float64 = math.Sqrt(dx*dx + dy*dy)
	return distance <= r
}
func (circ Circle) draw(scn screen) (err error) {
	// TO DO: Check if drawing this circle would cause either error
	if outOfBounds(Point{circ.center.x + circ.r, circ.center.y}, scn) ||
		outOfBounds(Point{circ.center.x - circ.r, circ.center.y}, scn) ||
		outOfBounds(Point{circ.center.x, circ.center.y + circ.r}, scn) ||
		outOfBounds(Point{circ.center.x, circ.center.y - circ.r}, scn) {
		return outOfBoundsErr
	}
	if colorUnknown(circ.c) {
		return colorUnknownErr
	}

	height := circ.center.y + circ.r
	width := circ.center.x + circ.r
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if insideCircle(circ.center, Point{x, y}, float64(circ.r)) {
				scn.drawPixel(x, y, circ.c)
			}
		}
	}
	return
}

// ----------------------------------------------------------------------
//
// printShape() method with Circle receiver
// returns the type (a Circle) and it's specs
func (circ Circle) printShape() string {
	var c Point
	var r int
	c, r = circ.center, circ.r

	return fmt.Sprintf("Circle: centered around (%d,%d) with radius %d",
		c.x, c.y, r)
}

// ----------------------------------------------------------------------
//
// clearScreen() method with Display pointer receiver
// Clears the screen by resetting it to white
// It does this in parallel, using at most 4 worker goroutines,
// as my computer has 4 cores.
func (display *Display) clearScreen() {
	var rows, cols int = display.maxX, display.maxY

	var numWorkers int = workers
	if numWorkers > rows {
		numWorkers = rows
	}

	var chunkSize int = (rows + numWorkers - 1) / numWorkers

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	var white Color = colorMap["white"]

	for w := 0; w < numWorkers; w++ {
		startRow := w * chunkSize
		endRow := startRow + chunkSize

		if endRow > rows {
			endRow = rows
		}

		go func(from int, to int) {
			defer wg.Done()
			for row := from; row < to; row++ {
				for col := 0; col < cols; col++ {
					display.matrix[row][col] = white
				}
			}
		}(startRow, endRow)
	}

	wg.Wait()
}

// ----------------------------------------------------------------------
//
// initialize() method with Display pointer receiver
// Initializes a screen of maxX times maxY
// Clears the screen so that it is all white
// It does this in parallel, using at most 4 worker goroutines,
// as my computer has 4 cores.
func (display *Display) initialize(x, y int) {
	display.maxX, display.maxY = x, y
	display.matrix = make([][]Color, x)

	numWorkers := workers
	if x < numWorkers {
		numWorkers = x
	}

	chunkSize := (x + numWorkers - 1) / numWorkers
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for w := 0; w < numWorkers; w++ {
		startRow := w * chunkSize
		endRow := startRow + chunkSize

		if endRow > x {
			endRow = x
		}

		go func(from int, to int) {
			defer wg.Done()
			for row := from; row < to; row++ {
				display.matrix[row] = make([]Color, y)
			}
		}(startRow, endRow)
	}

	wg.Wait()
	display.clearScreen()
}

// ----------------------------------------------------------------------
//
// getMaxXY() method with Display pointer receiver
// Retrieve and return the maxX and maxY dimensions of the screen
func (display *Display) getMaxXY() (x, y int) {
	return display.maxX, display.maxY
}

// ----------------------------------------------------------------------
//
// drawPixel() method with Display pointer receiver
// Draw the pixel with a given color at a given location
func (display *Display) drawPixel(x, y int, c Color) (err error) {
	if outOfBounds(Point{x, y}, display) {
		return outOfBoundsErr
	}
	if colorUnknown(c) {
		return colorUnknownErr
	}

	display.matrix[x][y] = c
	return nil
}

// ----------------------------------------------------------------------
//
// getPixel() method with Display pointer receiver
// Retrieve and return the color of the pixel at a given location
func (display *Display) getPixel(x, y int) (c Color, err error) {
	if outOfBounds(Point{x, y}, display) {
		return Color{}, outOfBoundsErr
	}
	return display.matrix[x][y], nil
}

// ----------------------------------------------------------------------
//
// screenShot() method with Display pointer receiver
// Write the screen to a .ppm file
// (which you can then view graphically with an image viewer)
func (display *Display) screenShot(f string) (err error) {

	file, e := os.Create(f + ".ppm")
	err = e
	if err != nil {
		fmt.Println("**Error creating ppm file: ", err)
		return
	}

	var x, y int = display.getMaxXY()
	fmt.Fprintln(file, "P3")
	fmt.Fprintf(file, "%d %d\n", x, y)
	fmt.Fprintln(file, "255")

	for row := 0; row < x; row++ {
		for col := 0; col < y; col++ {
			var p Color = display.matrix[row][col]
			fmt.Fprintf(file, "%d %d %d ", p.r, p.g, p.b)
		}
		fmt.Fprintln(file)
	}

	return nil
}
