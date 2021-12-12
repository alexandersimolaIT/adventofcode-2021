package main

import (
	"fmt" // To print to the console
	"os" // To open a text file
	"log" // To print error message
	"bufio" // To iterate over all rows in a file
	"strconv" // To convert a string to int
)
func main() {
	main_program()
}

func main_program() {
	var input = GetInputFromFile("input.txt")

	ComputeNumberOfIncreases2(input)
}

func test_ComputeNumberOfIncreases() {
	test_input := []int{2, 8, 4, 6, 2, 2, 7}
	fmt.Println("Expected output = 2")
	ComputeNumberOfIncreases(test_input)
}

func test_ComputeNumberOfIncreases2() {
	test_input := []int{2, 8, 4, 6, 2, 2, 7}
	fmt.Println("Expected output = 2")
	ComputeNumberOfIncreases2(test_input)
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

	if len(input) <= 3 {
		fmt.Print("The input is not large enough.")
		return
	}

	windowSize := 3

	var prevSum int
	var thisSum int

	prevSum = sum(input[:3])
	
	// Count the number of "input windows" that has a greater sum than the previous windows's sum
	for i:=windowSize; i<len(input); i++ {
		// Compute the sum of the elements [i-2,i]
		thisSum = sum(input[i-windowSize+1:i+1]) 
		
		if thisSum > prevSum {
			count += 1
		}

		prevSum = thisSum
	}

	fmt.Printf("Number of increases = %v\n", count)
}

// There is no need to compute the sum of the entire window actually.
// You only need to compute element i and i-3, since element i-2 and i-1 is included in both sums.
func ComputeNumberOfIncreases2(input []int) {
	fmt.Printf("Number of rows = %v\n", len(input))
	count := 0

	if len(input) <= 3 {
		fmt.Print("The input is not large enough.")
		return
	}

	windowSize := 3
	
	// Count the number of "input windows" that has a greater sum than the previous windows's sum
	for i:=windowSize; i<len(input); i++ {		
		if input[i] > input[i-windowSize] {
			count += 1
		}
	}

	fmt.Printf("Number of increases = %v\n", count)
}

func sum(array []int) int {  
	result := 0  
	for _, v := range array {  
	 result += v  
	}  
	return result  
   }