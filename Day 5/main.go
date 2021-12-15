package main

import (
	"fmt"
	"strconv"
	"strings"

	"adventofcode.com/shared"
)

func main() {

	main_program()
}

type point struct {
	x int
	y int
}

// Assumes point_string contains two numbers separated by a comma, like "123,432"
func newPoint(point_string string) *point {
	coords := strings.Split(point_string, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])

	return &point{x, y}
}

type line struct {
	a *point
	b *point
}

func newLine(point_a_string string, point_b_string string) *line {
	a := newPoint(point_a_string)
	b := newPoint(point_b_string)

	return &line{a, b}
}

func main_program() {
	input_strings := shared.GetInputFromFile("input.txt")
	fmt.Println("Hello")

	lines := parseInput(input_strings)

	fmt.Println("Printing all lines...")
	printLines(lines)

	//problem1(lines)

	problem2(lines)

}

func problem1(lines []*line) {
	lines1 := getLinesForPart1(lines)

	fmt.Println("Printing all lines for part 1...")
	printLines(lines1)

	//sortLinePoints(lines1)
	fmt.Println("Printing all lines after sorting their points...")
	printLines(lines1)

	// Testing getPointsOnLine
	points := getPointsOnLine(lines1[0])
	printPoints(points)

	_map := intersectLines(lines1)

	printPointMap(_map, 10, 10)

	intersections_count := 0
	for _, v := range _map {
		if v > 1 {
			intersections_count += 1
		}
	}

	fmt.Printf("intersections_count = %v\n", intersections_count)
}

// This turned out to be exactly the same code (minus the test sections) as problem1()
func problem2(lines []*line) {
	// Testing getPointsOnLine for a diagonal line
	points2 := getPointsOnLine(&line{&point{8, 0}, &point{0, 8}})
	fmt.Println("Printing all points on the line (8,0) -> (0,8)...")
	printPoints(points2)

	_map := intersectLines(lines)
	intersections_count := 0
	for _, v := range _map {
		if v > 1 {
			intersections_count += 1
		}
	}

	fmt.Printf("intersections_count = %v\n", intersections_count)
}

func parseInput(input_strings []string) []*line {
	result := []*line{}

	for _, element := range input_strings {
		point_strings := strings.Split(element, " -> ")
		result = append(result, newLine(point_strings[0], point_strings[1]))
	}

	return result
}

func printLines(lines []*line) {
	for _, l := range lines {
		printLine(l)
	}
}

func printLine(l *line) {
	printPoint(l.a)
	fmt.Print(" -> ")
	printPoint(l.b)
	fmt.Println()
}

func printPoints(points []*point) {
	for _, p := range points {
		printPoint(p)
		fmt.Println()
	}
}

func printPoint(p *point) {
	fmt.Printf("(%v,%v)", p.x, p.y)
}

func printPointMap(_map map[point]int, x_count int, y_count int) {
	for y := 0; y < y_count; y++ {
		for x := 0; x < x_count; x++ {
			fmt.Printf(" %v", _map[point{x, y}])
		}
		fmt.Println()
	}

}

func intersectLines(lines []*line) map[point]int {
	_map := make(map[point]int)

	for _, l := range lines {
		for _, p := range getPointsOnLine(l) {
			_map[point{p.x, p.y}] += 1
		}
	}

	return _map
}

// Returns all horizontal and vertical lines (where either x1=x2 or y1=y2)
func getLinesForPart1(lines []*line) []*line {
	result := []*line{}

	for _, l := range lines {
		if l.a.x == l.b.x || l.a.y == l.b.y {
			result = append(result, l)
		}

	}

	return result
}

// We want point a to have the lower value of x or y in each line
func sortLinePoints(lines []*line) {
	for _, l := range lines {
		if l.a.x > l.b.x || l.a.y > l.b.y {
			// Swap point a and b in the line
			p := l.a
			l.a = l.b
			l.b = p
		}
	}
}

func getPointsOnLine(l *line) []*point {
	result := []*point{}

	x_inc := 0
	if l.a.x > l.b.x {
		x_inc = -1
	} else if l.a.x < l.b.x {
		x_inc = 1
	}

	y_inc := 0
	if l.a.y > l.b.y {
		y_inc = -1
	} else if l.a.y < l.b.y {
		y_inc = 1
	}

	var n int
	if l.a.x == l.b.x {
		n = computeNumberOfPoints_VerticalLine(l)
	} else if l.a.y == l.b.y {
		n = computeNumberOfPoints_HorizontalLine(l)
	} else {
		n = computeNumberOfPoints_DiagonalLine(l)
	}

	x, y := l.a.x, l.a.y

	for i := 0; i < n; i++ {
		result = append(result, &point{x, y})
		x += x_inc
		y += y_inc
	}

	return result
}

func computeNumberOfPoints_HorizontalLine(l *line) int {
	num_points := l.a.x - l.b.x
	if num_points < 0 {
		num_points *= -1
	}
	return num_points + 1
}

func computeNumberOfPoints_VerticalLine(l *line) int {
	num_points := l.a.y - l.b.y
	if num_points < 0 {
		num_points *= -1
	}
	return num_points + 1
}

// Knowing that the line is 45 degrees, the number of points is easy to compute
func computeNumberOfPoints_DiagonalLine(l *line) int {
	return computeNumberOfPoints_HorizontalLine(l)
}
