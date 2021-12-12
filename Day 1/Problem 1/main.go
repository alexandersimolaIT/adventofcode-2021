package main

import (
	"fmt" // To print to the console
	"os" // To open a text file
	"log" // To print error message
	"bufio" // To iterate over all rows in a file
	"strconv" // To convert a string to int
)
func main() {
	var input = GetInputFromFile("input.txt")
	//test_input := []int{2, 8, 4, 6, 2, 2, 5}

	ComputeNumberOfIncreases(input)
}

func GetInputFromFile(file string) []int {
	// Get a "file descriptor" for the file "input.txt"
	f, err := os.Open(file)

	// Handle if we fail to open the file "input.txt"
    if err != nil {
        log.Fatal(err)
    }

	// Close the "file descriptor"
	defer f.Close() 

	// Create a file scanner
	scanner := bufio.NewScanner(f)

	// Make a slice to insert the number on each row in the file
	input := []int{}

	// Scan() makes a list of "tokens", one for each row
	// The loop iterates over each in order
	for scanner.Scan() {
		// Get the string for this row
		row := scanner.Text()

		// Convert the string to an int
		rowValAsInt, err := strconv.Atoi(row)
		
		// Handle if the conversion fails
		if err != nil {
			log.Fatal(err)
		}

		// Add the number on the row to the return list
		input = append(input, rowValAsInt)
	}

	// If Scan() fails, print the error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func ComputeNumberOfIncreases(input []int) {
	fmt.Printf("Number of rows = %v\n", len(input))
	count := 0

	// Count the number of lines that contain a number that is greater than the previous row's number
	for i:=1; i<len(input); i++ {
		if input[i] > input[i-1] {
			count += 1
		}
	}

	fmt.Printf("Number of increases = %v\n", count)
}