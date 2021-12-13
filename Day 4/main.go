package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"adventofcode.com/shared"
	"golang.org/x/text/runes"
)

func main() {
	main_program()

}

func main_program() {
	input := shared.GetInputFromFile("input.txt")

	rolled_numbers_str := strings.Split(input[0], ",")
	rolled_numbers := ConvertStringNumbersToIntArray(rolled_numbers_str)
	fmt.Printf("rolled_numbers = %v (type = %T)\n", rolled_numbers, rolled_numbers)

	bingo_boards_str := input[2:]

	// Remove empty elements
	bingo_boards := make([][][]int, 1)
	board_index := 0

	// Parse the string numbers into int arrays; one int array per board
	for _, element := range bingo_boards_str {
		if element == "" {
			board_index += 1
			bingo_boards = append(bingo_boards, [][]int{})
			continue
		}

		bingo_board := bingo_boards[board_index]
		next_row_str := RemoveBadElements(strings.Split(element, " "))

		next_row := ConvertStringNumbersToIntArray(next_row_str)
		bingo_boards[board_index] = append(bingo_board, next_row)
	}

	// Print all bingo boards
	for _, board := range bingo_boards {
		printBingoBoard(board)
		fmt.Println()
	}

	// 1st part
	winner_board, rolls, last_number := PlayBingo(rolled_numbers, bingo_boards)
	fmt.Println("Winning board:")
	printBingoBoard(winner_board)
	fmt.Printf("\nLast number = %v\n", last_number)
	fmt.Printf("Rolls = %v\n", rolls)

	unmarked_numbers_sum := 0

	is_rolled := make(map[int]bool)
	for _, e := range rolls {
		is_rolled[e] = true
	}

	for _, row := range winner_board {
		for _, number := range row {
			if !is_rolled[number] {
				unmarked_numbers_sum += number
			}
		}
	}

	fmt.Printf("Final answer (part 1):\nunmarked_numbers_sum * last_number = %v", unmarked_numbers_sum*last_number)

	// 2nd part
	loser_board, rolls2, last_number2 := PlayReverseBingo(rolled_numbers, bingo_boards)
	fmt.Println("\n\nLoser board:")
	printBingoBoard(loser_board)
	fmt.Printf("\nLast number = %v\n", last_number2)
	fmt.Printf("Rolls = %v\n", rolls2)

	unmarked_numbers_sum = 0

	is_rolled2 := make(map[int]bool)
	for _, e := range rolls2 {
		is_rolled2[e] = true
	}

	for _, row := range loser_board {
		for _, number := range row {
			if !is_rolled2[number] {
				unmarked_numbers_sum += number
			}
		}
	}

	fmt.Printf("Final answer (part 2):\nunmarked_numbers_sum * last_number = %v", unmarked_numbers_sum*last_number2)
}

func PlayBingo(all_rolls []int, bingo_boards [][][]int) ([][]int, []int, int) {
	current_rolls := []int{}
	var latest_roll int = -1

	for i := 0; i < len(all_rolls); i++ {
		latest_roll = all_rolls[i]
		current_rolls = append(current_rolls, latest_roll)

		for _, board := range bingo_boards {
			for i := 0; i < len(board); i++ {
				// Check row i
				row := board[i]
				if ArrayContains(current_rolls, row) {
					return board, current_rolls, latest_roll
				}

				// Check column i
				col := []int{}
				for _, row := range board {
					col = append(col, row[i])
				}
				if ArrayContains(current_rolls, col) {
					return board, current_rolls, latest_roll
				}

			}
		}

	}
	return nil, current_rolls, latest_roll
}

// For the second half of this puzzle
// Assumes that at some point there will be exactly one board left
func PlayReverseBingo(all_rolls []int, bingo_boards [][][]int) ([][]int, []int, int) {
	current_rolls := []int{}
	var latest_roll int = -1

	for i := 0; i < len(all_rolls); i++ {
		latest_roll = all_rolls[i]
		current_rolls = append(current_rolls, latest_roll)

		should_remove := make(map[int]bool)

		for j, board := range bingo_boards {
			//should_remove[j] = false

			for i := 0; i < len(board); i++ {
				// Check row i
				row := board[i]
				if ArrayContains(current_rolls, row) {
					should_remove[j] = true
				}

				// Check column i
				col := []int{}
				for _, row := range board {
					col = append(col, row[i])
				}

				if ArrayContains(current_rolls, col) {
					should_remove[j] = true
				}
			}
		}

		// Remove bingo boards that have won
		kept_boards := [][][]int{}
		for j, board := range bingo_boards {
			if !should_remove[j] {
				kept_boards = append(kept_boards, board)
			}
		}

		if len(kept_boards) == 0 {
			return bingo_boards[0], current_rolls, latest_roll
		}

		bingo_boards = kept_boards
	}
	return nil, current_rolls, latest_roll
}

func ArrayContains(container []int, containee []int) bool {
	contains := make(map[int]bool)
	for _, e := range container {
		contains[e] = true
	}

	for _, e := range containee {
		if !contains[e] {
			return false
		}
	}

	return true
}

func printBingoBoard(board [][]int) {
	width := len(board[0])

	for _, row := range board {
		for i, number := range row {
			if i%width == 0 {
				fmt.Println()
			}
			if number < 10 {
				fmt.Print(" ")
			}
			fmt.Printf(" %v", number)
		}
	}
}

func ConvertStringNumbersToIntArray(string_numbers []string) []int {

	numbers := []int{}

	for _, element := range string_numbers {
		number, err := strconv.Atoi(element)

		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, number)
	}

	return numbers
}

func RemoveBadElements(string_array []string) []string {
	result := []string{}
	for _, element := range string_array {
		if StringIsNumber(element) {
			result = append(result, element)
		}
	}
	return result
}

func StringIsNumber(input string) bool {
	if len(input) == 0 {
		return false
	}

	not_digits := runes.NotIn(unicode.Number)
	for _, char := range input {

		if not_digits.Contains(char) {
			return false
		}
	}

	return true
}
