// Author: M.Talha Adnan
// NetID: mkhan387
// Class: CS341 - Spring 2025

package main

import (
	"fmt"
	"os"
)

var display Display

func main() {
	// Print intro message
	fmt.Println("Project 5: Geometry Using Go Interfaces")
	fmt.Println("CS 341, Spring 2025")
	fmt.Println("")
	fmt.Println("This application allows you to draw various shapes")
	fmt.Println("of different colors via interfaces in Go.")
	fmt.Println("")

	// Ask user for dimensions for display
	var x, y int
	scanInput[int]("Enter the number of rows (x) that you would like the display to have: ", &x)
	scanInput[int]("Enter the number of columns (y) that you would like the display to have: ", &y)

	display.initialize(x, y)
	//
	// Menu options
	//
menuLoop:

	for {
		fmt.Println("")
		fmt.Println("Select a shape to draw: ")
		fmt.Println("	 R for a rectangle")
		fmt.Println("	 T for a triangle")
		fmt.Println("	 C for a circle")
		fmt.Println(" or X to stop drawing shapes.")

		// Accuire user input
		var choice string
		scanInput[string]("Your choice --> ", &choice)

		switch choice {
		case "X":
			var destFilename string
			scanInput[string]("Enter the name of the .ppm file in which the results should be saved: ", &destFilename)
			if err := display.screenShot(destFilename); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Done. Exiting program...")
			break menuLoop

		case "R":
			var x0, y0, x1, y1 int
			var colour string
			scanInput[int]("Enter the X and Y values of the lower left corner of the rectangle: ", &x0, &y0)
			scanInput[int]("Enter the X and Y values of the upper right corner of the rectangle: ", &x1, &y1)
			scanInput[string]("Enter the color of the rectangle: ", &colour)

			rect := Rectangle{
				ll: Point{x0, y0},
				ur: Point{x1, y1},
				c:  colorMap[colour],
			}
			fmt.Println(rect.printShape())
			if err := rect.draw(&display); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Rectangle drawn successfully.")
			}

		case "T":
			var x0, y0, x1, y1, x2, y2 int
			var colour string
			scanInput[int]("Enter the X and Y values of the first point of the triangle: ", &x0, &y0)
			scanInput[int]("Enter the X and Y values of the second point of the triangle: ", &x1, &y1)
			scanInput[int]("Enter the X and Y values of the third point of the triangle: ", &x2, &y2)
			scanInput[int]("Enter the color of the triangle: ", &colour)

			tri := Triangle{
				pt0: Point{x0, y0},
				pt1: Point{x1, y1},
				pt2: Point{x2, y2},
				c:   colorMap[colour],
			}
			fmt.Println(tri.printShape())
			if err := tri.draw(&display); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Triangle drawn successfully.")
			}

		case "C":
			var x, y, r int
			var colour string
			scanInput[int]("Enter the X and Y values of the center of the circle: ", &x, &y)
			scanInput[int]("Enter the value of the radius of the circle: ", &r)
			scanInput[int]("Enter the color of the circle: ", &colour)

			circ := Circle{
				center: Point{x, y},
				r:      r,
				c:      colorMap[colour],
			}
			fmt.Println(circ.printShape())
			if err := circ.draw(&display); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Circle drawn successfully.")
			}

		default:
			fmt.Println("**Error, unknown command. Try again.")

		}
	}
}

// Scans the console for input and puts it in the var specified and the
// The type of the var.
func scanInput[T any](inputPrompt string, dest ...any) {

	if inputPrompt != "" {
		fmt.Print(inputPrompt)
	}

	if len(dest) == 0 {
		panic("Error -> ( func ) scanInput: no destination provided.")
	}

	n, err := fmt.Scanln(dest...)
	if err != nil {
		panic(err.Error())
	}

	if n > len(dest) {
		fmt.Println("Read in more inputs than needed. Exiting program to prevent corruption.")
		os.Exit(1)
	}

}
