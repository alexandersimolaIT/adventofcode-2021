package shared

import (
	"bufio"
	"log"
	"os"
)

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
