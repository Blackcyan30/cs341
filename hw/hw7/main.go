// Header comment here!

package main

import (
	"fmt"
	// "io"
	// "os"
)

// ----------------------------------------------------------------------
//
// Entry struct
// Consists of:
//   - Website
//   - Username
//   - Password
type Entry struct {
	site, user, password string
}

// EntrySlice
// Each website contains a slice of password entries.
// Create a new type to make the code more readable.
type EntrySlice []Entry

// passwordMap
// Package level variable
// Map with
//
//	key: website (type: string)
//	value: slice of password entries (type: EntrySlice)
var passwordMap map[string]EntrySlice

// ----------------------------------------------------------------------
//
// main() function, program starts running here
func main() {
	passwordMap = make(map[string]EntrySlice)

	// Ask user if they want to initialize the map using a file
	fmt.Print("Enter a filename if you would like to initialize the map using a file\n")
	fmt.Print("(or enter N/A if the map should start as empty): ")
	var userChoice string
	fmt.Scanln(&userChoice)

	//
	// HERE ARE ALL THE PRINT STATEMENTS YOU NEED
	//
	// Reading in the file
	fmt.Println("Initializing map using file...")
	// os.Open() can be used to open a file for reading
	fmt.Println("**Error opening file. Exiting program...")
	// use os.Exit(1) if there is an error
	// compare to io.EOF to check if you have reached the end of the file
	fmt.Println("Done reading in file.")
	//
	// Menu options
	fmt.Println("")
	fmt.Println("Select a menu option: ")
	fmt.Println("	 L to list the contents of the map")
	fmt.Println("	 A to add a new entry to the map")
	fmt.Println("	 R to remove a website and/or user")
	fmt.Println(" or X to exit the program.")
	fmt.Print("Your choice --> ")
	fmt.Println("**Error, unknown command. Try again.")
	//
	// Listing the contents of the map
	fmt.Printf("Website: %s\n", "google.com")
	fmt.Printf("\t %s \t %s\n", "someUserName", "doNotUseThisP@ssw0rd!")
	//
	// Adding an entry
	fmt.Print("Enter the site, username, and password (separated by spaces): ")
	fmt.Println("**Error: Attempting to add a duplicate entry. Try again.")
	//
	// Removing a website / user
	fmt.Print("Enter the site and username (separated by spaces, username optional): ")
	fmt.Println("**Error: Attempt to remove a website that does not exist in the map. Try again.")
	fmt.Println("**Error: Attempt to remove a username that does not exist in the map. Try again.")
	fmt.Println("**Error: Attempt to remove multiple users. Try again.")
	//
	// End of program
	fmt.Println("Exiting program.")
}
