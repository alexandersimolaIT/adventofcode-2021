package main

import (
	"fmt"

	"adventofcode.com/shared"
)

func main() {
	main_program()
}

func main_program() {
	gamma, epsilon := problem1()
	fmt.Printf("gamma = %v, epsilon = %v\n", gamma, epsilon)
	fmt.Printf("gamma * epsilon = %v\n", gamma*epsilon)

	oxygen_rating := problem2(false)
	co2_rating := problem2(true)
	fmt.Printf("Oxygen generator rating = %v\n", oxygen_rating)
	fmt.Printf("CO2 scrubbing rating = %v\n", co2_rating)
	fmt.Printf("{Oxygen generator rating} * {CO2 scrubbing rating} = %v\n", oxygen_rating*co2_rating)
}

func problem1() (int, int) {
	input_strings := shared.GetInputFromFile("input.txt")
	input, _ := shared.ConvertBinaryNumberStringsArrayToInts(input_strings)

	input_length := len(input)
	bit_count := len(input_strings[0]) // Assume that all strings contain the same number of characters

	bit_counts := make([]int, bit_count)

	gamma := 0

	for j := 0; j < bit_count; j++ {
		mask := 1 << j
		// Iterate over all strings, one for each row in the new matrix
		for _, element := range input {
			bit_counts[bit_count-j-1] += (mask & element) >> j
		}
	}

	// Compute gamma
	for i := 0; i < bit_count; i++ {
		// If majority of bits at index i is 1, then make i'th bit in gamma 1
		if float32(bit_counts[bit_count-i-1]) >= float32(input_length/2) {
			gamma = gamma | (1 << i)
		}
	}

	//epsilon := (2 ^ bit_count - 1) & (^gamma)
	epsilon := (1<<bit_count - 1) - gamma

	return gamma, epsilon
}

// Assumes all elements in the input string-array have the same length
// Returns an array of bit counts where element[0] represents the number of 1-bits in the leftmost-character
func Count1BitsAtEachPosition(input []string) []int {
	bit_count := len(input[0]) // Assume that all strings contain the same number of characters
	bit_counts := make([]int, bit_count)

	for _, element := range input {

		for j := 0; j < bit_count; j++ {
			if element[j] == byte('1') {
				bit_counts[j] += 1
			}
		}

	}

	return bit_counts
}

func test_Count1BitsAtEachPosition() {
	test_input := shared.GetInputFromFile("test_input.txt")

	expected_bit_counts := []int{7, 5, 8, 7, 5}
	actual_bit_counts := Count1BitsAtEachPosition(test_input)

	fmt.Printf("test_input = %v, (type = %T)\n", test_input, test_input)
	fmt.Printf("expected_bit_counts = %v, (type = %T)\n", expected_bit_counts, expected_bit_counts)
	fmt.Printf("actual_bit_counts = %v, (type = %T)\n", actual_bit_counts, actual_bit_counts)
}

func problem2(look_for_least_common_bit bool) int {

	input := shared.GetInputFromFile("input.txt")
	input_width := len(input[0])
	var answer string

	for i := 0; i < input_width && len(input) > 1; i++ {
		filtered_input := []string{}

		bit_counts := Count1BitsAtEachPosition(input)
		num_ones := bit_counts[i]

		var keep_ones bool
		if look_for_least_common_bit {
			keep_ones = !(2*num_ones >= len(input))
		} else {
			keep_ones = (2*num_ones >= len(input))
		}

		var keep_bit byte
		if keep_ones {
			keep_bit = '1'
		} else {
			keep_bit = '0'
		}

		// Fill the filtered input with all entries that match the most common bit at position i from the left (where i=0 is leftmost)
		for _, element := range input {
			if element[i] == keep_bit {
				filtered_input = append(filtered_input, element)
			}
		}

		input = filtered_input
		answer = filtered_input[0]
	}

	answer_int, _ := shared.ConvertBinaryNumberStringToInt(answer)

	return answer_int
}

func problem2_with_remove(bit_counts []int) (int, int) {

	filtered_input := shared.GetInputFromFile("test_input.txt")

	for i := 0; i < len(bit_counts); i++ {
		input := []string{}
		copy(filtered_input, input)

		num_ones := bit_counts[i]

		var most_common_bit byte
		if num_ones >= len(input) {
			most_common_bit = '1'
		} else {
			most_common_bit = '0'
		}

		for j := len(input) - 1; j >= 0; j-- {

			element := input[j]

			if element[i] != most_common_bit {
				filtered_input = removeSliceElementAtIndex(input, j)
			}
		}
	}

	return 0, 0
}

func removeSliceElementAtIndex(slice []string, index int) []string {
	if index == 0 {
		return slice[1:]
	} else if index == len(slice)-1 {
		return slice[:index]
	} else {
		return append(slice[:index], slice[index+1:]...)
	}
}
