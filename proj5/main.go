// Header comment here!
// Author: M.Talha Adnan
// NetID: mkhan387
// Class: CS341 - Spring 2025
package main

import (
	"fmt"
	"os"
)

func Assert(cond bool, msg string) {
	if !cond {
		panic("ASSERTION FAILED: " + msg)
	}
}

var display Display

func main() {
	// Print intro message
	fmt.Println("Project 5: Geometry Using Go Interfaces")
	fmt.Println("CS 341, Spring 2025")
	fmt.Println("")
	fmt.Println("This application allows you to draw various shapes")
	fmt.Println("of different colors via interfaces in Go.")
	fmt.Println("")

	//
	// SOME PRINT STATEMENTS YOU WILL NEED CAN BE FOUND BELOW
	//
	// Ask user for dimensions for display
	// fmt.Print("Enter the width (x) that you would like the display to be: ")
	// fmt.Print("Enter the height (y) that you would like the display to be: ")
	// Enter the number of rows (x) that you would like the display to have: Enter the number of columns (y) that you would like the display to have:

	var width, height int
	scanInput[int]("Enter the number of rows (x) that you would like the display to have: ", &height)
	scanInput[int]("Enter the number of columns (y) that you would like the display to have: ", &width)

	display.initialize(width, height)

menuLoop:
	for {
		//
		// Menu options
		//
		fmt.Println("")
		fmt.Println("Select a shape to draw: ")
		fmt.Println("	 R for a rectangle")
		fmt.Println("	 T for a triangle")
		fmt.Println("	 C for a circle")
		fmt.Println(" or X to stop drawing shapes.")
		var choice string
		scanInput[string]("Your choice --> ", &choice)

		switch choice {
		case "R":
			// Rectangle
			var x0, y0, x1, y1 int
			var colour string
			scanInput[int]("Enter the X and Y values of the lower left corner of the rectangle: ", &x0, &y0)
			scanInput[int]("Enter the X and Y values of the upper right corner of the rectangle: ", &x1, &y1)
			scanInput[string]("Enter the color of the rectangle: ", &colour)

			rect := Rectangle{
				ll: Point{x0, y0},
				ur: Point{x1, y1},
				c:  Color(colour),
			}
			fmt.Println(rect.printShape())
			if err := rect.draw(&display); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Rectangle drawn successfully.")
			}

		case "T":
			// Triangle
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
				c:   Color(colour),
			}
			fmt.Println(tri.printShape())
			if err := tri.draw(&display); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Triangle drawn successfully.")
			}

		case "C":
			// Circle
			var x, y, r int
			var colour string
			scanInput[int]("Enter the X and Y values of the center of the circle: ", &x, &y)
			scanInput[int]("Enter the value of the radius of the circle: ", &r)
			scanInput[int]("Enter the color of the circle: ", &colour)

			circ := Circle{
				center: Point{x, y},
				r:      r,
				c:      Color(colour),
			}
			fmt.Println(circ.printShape())
			if err := circ.draw(&display); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Circle drawn successfully.")
			}

		case "X":
			var outName string
			scanInput[string]("Enter the name of the .ppm file in which the results should be saved: ", &outName)
			if err := display.screenShot(outName); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Done. Exiting program...")
			break menuLoop

		default:
			fmt.Println("**Error, unknown command. Try again.")
		}
	}

	//
	// Menu options
	// fmt.Println("")
	// fmt.Println("Select a shape to draw: ")
	// fmt.Println("	 R for a rectangle")
	// fmt.Println("	 T for a triangle")
	// fmt.Println("	 C for a circle")
	// fmt.Println(" or X to stop drawing shapes.")
	// fmt.Print("Your choice --> ")
	// fmt.Println("**Error, unknown command. Try again.")
	// //
	// // Drawing a rectangle
	// fmt.Print("Enter the X and Y values of the lower left corner of the rectangle: ")
	// fmt.Print("Enter the X and Y values of the upper right corner of the rectangle: ")
	// fmt.Print("Enter the color of the rectangle: ")
	// fmt.Println("Rectangle drawn successfully.")
	// //
	// // Drawing a triangle
	// fmt.Print("Enter the X and Y values of the first point of the triangle: ")
	// fmt.Print("Enter the X and Y values of the second point of the triangle: ")
	// fmt.Print("Enter the X and Y values of the third point of the triangle: ")
	// fmt.Print("Enter the color of the triangle: ")
	// fmt.Println("Triangle drawn successfully.")
	// //
	// // Drawing a circle
	// fmt.Print("Enter the X and Y values of the center of the circle: ")
	// fmt.Print("Enter the value of the radius of the circle: ")
	// fmt.Print("Enter the color of the circle: ")
	// fmt.Println("Circle drawn successfully.")
	// //
	// // Saving the results in a file
	// fmt.Print("Enter the name of the .ppm file in which the results should be saved: ")
	// //
	// // Exiting program
	// fmt.Println("Done. Exiting program...")
}

func scanInput[T any](inputPrompt string, dest ...any) {

	if inputPrompt != "" {
		fmt.Print(inputPrompt)
	}

	if len(dest) == 0 {
		panic("scanInput: no destination provided.")
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
