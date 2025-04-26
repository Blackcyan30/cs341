// Name: Talha Adnana
// Course: CS 341, Prof. Kidane
// Homework: 7
// NetID: mkhan387

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Assert(cond bool, msg string) {
	if !cond {
		panic("ASSERTION FAILED: " + msg)
	}
}

func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

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

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var input string
var err error

// ----------------------------------------------------------------------
//
// main() function, program starts running here
func main() {
	passwordMap = make(map[string]EntrySlice)

	// Ask user if they want to initialize the map using a file
	fmt.Print("Enter a filename if you would like to initialize the map using a file\n")
	fmt.Print("(or enter N/A if the map should start as empty): ")

	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		os.Exit(1)
		return
	}
	input = strings.TrimSpace(input)

	// Reading the file
	if input != "N/A" {
		fmt.Println("\nInitializing map using file...")
		var file *os.File
		var err error
		file, err = os.Open(input)

		if err != nil {
			fmt.Println("**Error opening file. Exiting program...")
			file.Close()
			os.Exit(1)
		}
		// Trim the input to remove any leading/trailing whitespace
		input = strings.TrimSpace(input)

		var scanner *bufio.Scanner = bufio.NewScanner(file)
		Assert(scanner != nil, "Scanner should not be nil")

		var fileEntry string
		var fileEntryComponents []string
		var passwordMapEntry Entry

		for scanner.Scan() {
			fileEntry = scanner.Text()
			Assert(fileEntry != "", "File entry should not be empty")
			fileEntryComponents = strings.Split(fileEntry, " ")
			Assert(len(fileEntryComponents) == 3, "Invalid file format: File entry should have exactly 3 components")

			var site string = fileEntryComponents[0]
			var user string = fileEntryComponents[1]
			var password string = fileEntryComponents[2]

			passwordMapEntry = Entry{site: site, user: user, password: password}

			var prevLen int = len(passwordMap[site])
			passwordMap[site] = append(passwordMap[site], passwordMapEntry)

			Assert(len(passwordMap[site]) == prevLen+1, "Failed to add entry to map")
		}

		fmt.Println("Done reading in file.")
		file.Close()
	}

	//
	// Menu options
	//
	printMenu()

	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		os.Exit(1)
		return
	}
	input = strings.TrimSpace(input)

	for input != "X" {
		switch input {
		case "L":
			listPasswordMap()
		case "A":
			addEntry()
		case "R":
			removeEntry()
		default:
			fmt.Println("**Error, unknown command. Try again.")
		}
		printMenu()
		// fmt.Scanln(&input)
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input.")
			os.Exit(1)
			return
		}
		input = strings.TrimSpace(input)
	}

	fmt.Println("Exiting program.")
}

func printMenu() {
	fmt.Println("")
	fmt.Println("Select a menu option: ")
	fmt.Println("	 L to list the contents of the map")
	fmt.Println("	 A to add a new entry to the map")
	fmt.Println("	 R to remove a website and/or user")
	fmt.Println(" or X to exit the program.")
	fmt.Print("Your choice --> ")
}

func listPasswordMap() {
	if len(passwordMap) == 0 {
		return
	}

	for site, entries := range passwordMap {
		fmt.Println("Website:", site)
		for _, entry := range entries {
			fmt.Printf("\t %s \t %s\n", entry.user, entry.password)
		}
	}
}

func addEntry() {
	fmt.Printf("Enter the site, username, and password (separated by spaces): ")

	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error reading input.")
		return
	}

	// Tokenize the input
	input = strings.TrimSpace(input)
	var components []string = strings.Split(input, " ")

	assert(len(components) == 3, "Invalid input format: Expected 3 components: site, username, and password.")

	var site string = components[0]
	var username string = components[1]
	var password string = components[2]

	// Checking for duplicate entry
	entryArr, ok := passwordMap[site]
	if ok {
		for _, entry := range entryArr {
			if entry.user == username {
				fmt.Printf("**Error: Attempting to add a duplicate entry. Try again.")
				return
			}
		}
	}

	// Add new entry to the map
	var newEntry Entry = Entry{site: site, user: username, password: password}
	var prevLen int = len(passwordMap[site])
	passwordMap[site] = append(passwordMap[site], newEntry)

	Assert(len(passwordMap[site]) == prevLen+1, "Failed to add entry to map")
}

func removeEntry() {
	fmt.Print("Enter the site and username (separated by spaces, username optional): ")

	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}

	// Tokenize the input
	input = strings.TrimSpace(input)
	var components []string = strings.Split(input, " ")

	assert(len(components) >= 1, "Invalid input format: Expected at least 1 component: site.")

	var site string = components[0]

	// Check if website exists
	var entryArr EntrySlice
	var ok bool
	entryArr, ok = passwordMap[site]

	if !ok {
		fmt.Println("**Error: Attempt to remove a website that does not exist in the map. Try again.")
		return
	}

	// If only site provided, remove entire site
	if len(components) == 1 {
		if len(entryArr) == 1 {
			delete(passwordMap, site)
			return
		}
		fmt.Println("**Error: Attempt to remove multiple users. Try again.")
		return
	}

	var newEntries EntrySlice
	// If site and username provided, remove specific user
	if len(components) == 2 {
		var username string = components[1]
		var found bool = false

		for i := range entryArr {
			if entryArr[i].user == username {
				newEntries = append(entryArr[:i], entryArr[i+1:]...)
				found = true
				break
			}
		}

		if found {
			if len(newEntries) == 0 {
				delete(passwordMap, site)
			} else {
				passwordMap[site] = newEntries
			}
		} else {
			fmt.Println("**Error: Attempt to remove a username that does not exist in the map. Try again.")
			return
		}

		return
	}
}
