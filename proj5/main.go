// Header comment here!

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

	var width, height int
	scanInput[int]("Enter the width (x) that you would like the display to be: ", &width)
	scanInput[int]("Enter the height (y) that you would like the display to be: ", &height)

	display.initialize(width, height)
	// for input != "X" {
	// 	switch input {
	// 	case "L":
	// 		listPasswordMap()
	// 	case "A":
	// 		addEntry()
	// 	case "R":
	// 		removeEntry()
	// 	default:
	// 		fmt.Println("**Error, unknown command. Try again.")
	// 	}
	// 	printMenu()
	// 	// fmt.Scanln(&input)
	// 	input, err = reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println("Error reading input.")
	// 		os.Exit(1)
	// 		return
	// 	}
	// 	input = strings.TrimSpace(input)
	// }

	var userInput string
	scanInput[string]("Your choice --> ", &userInput)
mainloop:
	for {
		switch userInput {
		case "X":
			break mainloop
		case "R":
			var llc, urc Point
			var color string
			scanInput[int]("Enter the X and Y values of the lower left corner of the rectangle: ", &llc.x, &llc.y)
			scanInput[int]("Enter the X and Y values of the upper right corner of the rectangle: ", &urc.x, &urc.y)
			scanInput[string]("Enter the color of the rectangle: ", &color)
			fmt.Println("Rectangle drawn successfully.")

		case "T":
			var p1, p2, p3 Point
			var color string
			scanInput[int]("Enter the X and Y values of the first point of the triangle: ", &p1.x, &p1.y)
			scanInput[int]("Enter the X and Y values of the second point of the triangle: ", &p2.x, &p2.y)
			scanInput[int]("Enter the X and Y values of the third point of the triangle: ", &p3.x, &p3.y)
			scanInput[int]("Enter the color of the triangle: ", &color)
			fmt.Println("Triangle drawn successfully.")

		case "C":
			var center Point
			fmt.Print("Enter the X and Y values of the center of the circle: ")
			scanInput[int]("Enter the X and Y values of the first point of the triangle: ", &p1.x, &p1.y)
			fmt.Print("Enter the value of the radius of the circle: ")
			fmt.Print("Enter the color of the circle: ")
			fmt.Println("Circle drawn successfully.")

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

	n, err := fmt.Scanln(dest...)
	Assert(err == nil, err.Error())

	if n == len(dest) {
		fmt.Println("Read in more inputs than needed. Exiting program to prevent corruption.")
		os.Exit(1)
	}

}
