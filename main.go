package main

import "os"
import "fmt"
import "strings"
import "io/ioutil"
import "encoding/json"

var validchars = " .:T"

func main() {

	// If no arguments were provided, show a helpful message and then abort
	if len(os.Args) == 1 {
		fmt.Println("Welcome to Avalanche. Please specify a json file as an argument:")
		fmt.Println("> ./avalanche examples/example.json")
		os.Exit(0)
	}

	// Read the file path to the json file
	filepath := os.Args[1]

	// Read the json into an Arena struct
	original := createArenaFromArray(readJSONFile(filepath))

	// Create a clone of the arena struct, for "before" and "after" output
	completed := original.duplicate()

	// Print the original arena struct
	fmt.Println("Start:")
	fmt.Print(original.toString())

	// Run the steps needed to compute the fallen rock positions
	rounds := completed.solve()

	// Print the completed struct to show the solution
	fmt.Println("End:")
	fmt.Print(completed.toString())

	// How many rounds of moves did it take?
	fmt.Printf("Rounds: %d\n", rounds)
}

// Read the contents of a properly formated JSON file into an array of strings
func readJSONFile(filepath string) [][]string {

	// Read the raw file contents
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Decode the json string into an array of string arrays
	var input [][]string
	json.Unmarshal(raw, &input)

	// Return the decoded arrays
	return input
}

// Convert an array of string arrays into an Arena struct instance
func createArenaFromArray(input [][]string) Arena {

	// Create a struct instance
	var arena Arena

	// How many rows of input are there?
	rows := len(input)

	// If no rows are present we should reject the input
	if rows == 0 {
		fmt.Println("No valid input was provided")
		os.Exit(1)
	}

	// How many columns of input are there?
	var cols = len(input[0])

	// If no columns are present, we should reject the input
	if cols == 0 {
		fmt.Println("No valid input was provided")
		os.Exit(1)
	}

	// Establish empty Column structs for each known column
	for c := 0; c < cols; c++ {
		arena.columns = append(arena.columns, Column{})
	}

	// Read the columnar data from the array and sort it into the new struct
	var char string
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			// Read the input character
			char = input[r][c]

			// Ensure the character is white listed
			if !strings.Contains(validchars, input[r][c]) {
				char = " "
			}

			// Copy the validated character into its column slice
			arena.columns[c].cells = append(arena.columns[c].cells, char)
		}
	}

	return arena
}

/******************************************************************************
  Arenas
 ******************************************************************************/

// Arena Structure
type Arena struct {
	columns []Column
}

// Move all arena cell contents until they can't move any more
func (a *Arena) solve() int {
	ticks := 0

	// Perform one round of moves in each column
	movement := false
	for i := range a.columns {
		if a.columns[i].doTick() {
			movement = true
		}
	}

	// Did a move just occur?
	if movement {
		ticks++
	}

	round := false

	// If moves were made, we can go into our loop
	for movement == true {

		// Perform ticks on each column in the arena
		for i := range a.columns {
			if a.columns[i].doTick() {
				round = true
			}
		}
		// fmt.Println(a.toString())
		movement = round

		// Did a movement occur in this round?
		if round {
			ticks++
		}

		// reset our round movement check
		round = false
	}

	return ticks
}

// Convert an Arena to a string for ease of printing
func (a *Arena) toString() string {

	// Initialize our return string
	output := ""

	// How many columns are there?
	cols := len(a.columns)

	// If there are no columns, return an empty string
	if cols == 0 {
		return output
	}

	// How many rows are there?
	rows := len(a.columns[0].cells)

	// If there is only one column, we will skip the border ascii
	if cols == 1 {
		for r := 0; r < rows; r++ {
			output += a.columns[0].cells[r] + "\n"
		}

		return output
	}

	// Build our ascii cap line
	var cap = "+"
	for c := 0; c < cols; c++ {
		cap += "-"
	}
	cap += "+\n"

	// Add the top cap
	output += cap

	// Print the arena contents, row by row
	var line string
	for r := 0; r < rows; r++ {
		line = "|"
		for c := 0; c < cols; c++ {
			line += a.columns[c].cells[r]
		}
		line += "|\n"
		output += line
	}

	// Add the end cap
	output += cap

	return output

}

// A helper method to clone Arena structs
func (a *Arena) duplicate() Arena {

	// Initialize our new struct
	duplicate := Arena{}

	// Create duplicates of each original column and add them to the new Arena
	for i := 0; i < a.numberOfColumns(); i++ {
		c := make([]string, a.numberOfRows())
		copy(c, a.columns[i].cells)

		duplicate.columns = append(duplicate.columns, Column{cells: c})
	}

	return duplicate
}

// Determine the number of columns in the Arena
func (a *Arena) numberOfColumns() int {
	return len(a.columns)
}

// Determine the number of rows in the Arena
func (a *Arena) numberOfRows() int {
	if len(a.columns) == 0 {
		return 0
	}

	return len(a.columns[0].cells)
}

/******************************************************************************
  Columns
 ******************************************************************************/

// Column structure
type Column struct {
	cells []string
}

// Perform one movement step for this column
func (c *Column) doTick() bool {

	// We can only move pieces if there are two or more spaces available
	if len(c.cells) < 2 {
		return false // No movement occurred
	}

	// Find the index of the last cell
	end := len(c.cells) - 1

	// Loop through each possible "target" cell in this column and
	// perform movements if possible
	movement := false
	for i := end; i > 0; i-- {
		if c.move(i-1, i) {
			movement = true
		}
	}

	// Return an indicator of whether or not a move occurred
	return movement
}

// Determine if a given index has any room to move below it
func (c *Column) isIndexBlocked(index int) bool {

	// How many cells in this column?
	length := len(c.cells)

	// An empty set of cells is deemed blocked,
	// though this should probably return an error
	if length == 0 {
		return true
	}

	// A negative index is deemed blocked,
	// though this should probably return an error
	if index < 0 {
		return true
	}

	// What is the index of the last cell?
	end := len(c.cells) - 1

	// An out of bounds index is deemed blocked,
	// though this should probably return an error
	if index >= end {
		return true
	}

	// Find the index of the next "T", aka "blocking", cell below this one
	limiter := find(c.cells[index:], "T")

	if limiter == 0 {
		// If no limiter was found, the limiter is the end of the column
		limiter = end + 1
	} else {
		// Otherwise, calculate the true index of the limiter using the offset
		limiter += (index + 1)
	}

	// If there is space in the cells between the index and the limiter,
	// this cell is not blocked
	if contains(c.cells[index+1:limiter], " ") {
		return false
	}

	// We shouldn't ever get to this point.
	return true
}

// Attempt to move the contents of one index into another location. This is
// intended for use between adacent cells. Return true if a move occurred
func (c *Column) move(i int, j int) bool {

	// How many cells in this column?
	length := len(c.cells)

	// If there are no cells, abort the move
	if length == 0 {
		return false
	}

	// Abort on negative indices
	if i < 0 || j < 0 {
		return false
	}

	// Determine the last index of the cell slice
	end := len(c.cells) - 1

	// Abort on indices beyond the end of the slice
	if i > end || j > end {
		return false
	}

	// Abort if the indices point to the same location
	if i == j {
		return false
	}

	// Abort if either the source or the target are tables
	if c.cells[i] == "T" || c.cells[j] == "T" {
		return false
	}

	// Determine if our target cell has room to move itself
	targetBlocked := c.isIndexBlocked(j)

	// If the target is blocked, abort if the indices contain the same values
	if !targetBlocked && c.cells[i] == c.cells[j] {
		return false
	}

	// If the target is blocked, check to see if we can combine two rocks
	if targetBlocked && c.cells[j] == "." && c.cells[i] == "." {
		c.cells[i] = " "
		c.cells[j] = ":"
		return true
	}

	// Abort if the target cell is blocked and already has a set of collapsed rocks
	if targetBlocked && c.cells[j] == ":" {
		return false
	}

	// If the target cell is empty, proceed with the move
	if c.cells[j] == " " && c.cells[i] != " " {
		c.cells[j] = c.cells[i]
		c.cells[i] = " "

		return true
	}

	// Lets not assume we can move without meeting specific known criteria
	return false
}

// A helper method to see if a string slice haystack contains a string needle
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// A helper method to find the index of a string needing in a string slice haystack
func find(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return 0
}

// A helper method to check the equality of two string slices
// https://stackoverflow.com/a/15312097
func slicesAreEqual(a []string, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
