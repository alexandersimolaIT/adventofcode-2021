package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	main_program()
}

func main_program() {
	input := GetInputFromFile("input.txt")
	horizontal_pos, depth := ComputeFinalPosition(input)
	
	fmt.Println("Running main program...")
	fmt.Printf("\tFinal position: {horizontal position = %v, depth = %v}\n", horizontal_pos, depth)
	fmt.Printf("\t(Horizontal pos) * (depth) = %v", horizontal_pos * depth)
}

func test_arrayContainsString() {
	array := []string{"asdf", "foo", "b123"}
	fmt.Printf("array %v contains %v = %v\n", array, "foo", arrayContainsString(array, "foo"))
	fmt.Printf("array %v contains %v = %v\n", array, "bar", arrayContainsString(array, "bar"))
}

func test_GetInputFromFile() {
	file_name := "test_input.txt"

	fmt.Printf("Running GetInputFromFile('%v')...\n", file_name)
	input := GetInputFromFile(file_name)
	fmt.Printf("\tinput = %v\n", input)
}

func test_ComputeFinalPosition() {
	input := GetInputFromFile("test_input.txt")
	expected_horizontal_pos := 15
	expected_depth := 60

	fmt.Printf("Running ComputeFinalPosition(input)...\n")
	fmt.Printf("\tinput = %v\n", input)
	fmt.Printf("\tExpected horizontal position = %v\n", expected_horizontal_pos)
	fmt.Printf("\tExpected depth = %v\n", expected_depth)
	horizontal_pos, depth := ComputeFinalPosition(input)
	
	fmt.Printf("\nResult:\n\tHorizontal position = %v\n\tDepth = %v\n", horizontal_pos, depth)
}

func GetInputFromFile(file string) []string {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := []string{} 

	for scanner.Scan() {
		row := scanner.Text()

		result = append(result, row)
	}

	return result
}

func ComputeFinalPosition(input []string) (int, int) {
	valid_directions := []string{"forward", "up", "down"}
	aim := 0
	horizontal_pos := 0
	depth := 0

	for _, row := range input {
		words := strings.Split(row, " ")
		direction, stepsAsString := words[0], words[1]
		x, err := strconv.Atoi(stepsAsString)

		// Handle if string conversion failed
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Printf("Direction = %v, steps = %v\n", direction, steps)
		
		if !arrayContainsString(valid_directions, direction) {
			
			fmt.Printf("'%v' is not a valid direction...", direction)
			return -1, -1
		}

		switch direction {
		case "forward":
			horizontal_pos += x
			depth += aim * x
		case "up":
			aim -= x
		case "down":
			aim += x
		default:
			fmt.Printf("Somehow, direction = '%v', which is undefined...", direction)
		}
	}

	return horizontal_pos, depth
	//return horizontal_pos, depth
	
}

func arrayContainsString(array []string, _string string) bool {
	for _, element := range array {
		if element == _string {
			return true
		}
	}
	return false
}