package shared

import (
	"errors"
	"fmt"
	"log"
)

func ConvertBinaryNumberStringsArrayToInts(string_numbers []string) ([]int, error) {
	int_numbers := []int{}

	for _, string_element := range string_numbers {
		int_element, err := ConvertBinaryNumberStringToInt(string_element)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		int_numbers = append(int_numbers, int_element)
	}

	return int_numbers, nil
}

func test_ConvertBinaryNumberStringsArrayToInts() {
	test_input := []string{"101", "10001", "111", "10", "0", "1"}

	expected_output := []int{5, 17, 7, 2, 0, 1}
	actual_output, _ := ConvertBinaryNumberStringsArrayToInts(test_input)

	fmt.Printf("Input = [%v] (type = %T)\n", test_input, test_input)
	fmt.Printf("Expected output = [%v] (type = %T)\n", expected_output, expected_output)
	fmt.Printf("Actual output = [%v] (type = %T)\n", actual_output, actual_output)
}

func ConvertBinaryNumberStringToInt(binary_number_as_string string) (int, error) {
	// Check if the string is valid, i.e. only contains no other characters but "0" and "1".
	for _, char := range binary_number_as_string {
		if char != '1' && char != '0' {
			return -1, errors.New("Invalid input '%v'! Must only contain '0' or '1'.")
		}
	}

	result := 0
	bit_count := len(binary_number_as_string)

	for i, char := range binary_number_as_string {

		if char == '1' {
			result = result | 1<<(bit_count-1-i)
		}
	}

	return result, nil
}

func test_ConvertBinaryNumberStringToInt() {
	test_input := "101"
	expected_output := 5
	output, _ := ConvertBinaryNumberStringToInt(test_input)

	fmt.Printf("test_input = '%v' (type = %T), expected_output = %v (type = %T), output = '%v'", test_input, test_input, expected_output, expected_output, output)
}
