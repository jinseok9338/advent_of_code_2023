package problems

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Day3() int {
	const file_path = "./problems/inputs/day3_input.txt"

	// read line by line from the file
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		temp := []string{}
		for _, char := range line {
			temp = append(temp, string(char))
		}

		matrix = append(matrix, temp)

	}

	occurrences := FindNumbersOccurrences(matrix)
	sum := 0
	for _, occurrence := range occurrences {
		result, number := CheckSymbol(matrix, occurrence)
		if result {
			sum += number
		}
		fmt.Println(result, number)
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return 0
}

func FindNumbersOccurrences(matrix [][]string) [][][]int {
	occurrences := [][][]int{}

	for i, row := range matrix {
		start, end := -1, -1
		for j, cell := range row {
			if unicode.IsDigit(rune(cell[0])) && cell != "*" {
				// If it's the start of a new sequence
				if start == -1 {
					start = j
				}
				// Update the end index for the current sequence
				end = j
			} else {
				// If the current cell is not a digit, and there was an ongoing sequence
				if start != -1 {
					startIndex := []int{i, start}
					endIndex := []int{i, end}
					occurrences = append(occurrences, [][]int{startIndex, endIndex})
					start, end = -1, -1
				}
			}
		}

		// Check for an ongoing sequence at the end of the row
		if start != -1 {
			startIndex := []int{i, start}
			endIndex := []int{i, end}
			occurrences = append(occurrences, [][]int{startIndex, endIndex})
		}
	}

	return occurrences
}

func isValidIndex(matrix [][]string, i, j int) bool {
	return i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0])
}
func buildNumber(matrix [][]string, indexes [][]int) int {
	number := 0

	for i := indexes[0][1]; i < indexes[1][1]+1; i++ {
		indexPair := []int{indexes[0][0], i}
		i, j := indexPair[0], indexPair[1]

		// Ensure the starting index is valid
		if !isValidIndex(matrix, i, j) {
			continue
		}

		// Iterate over the columns starting from the specified index

		// Parse the digit and concatenate it to the number
		if digit, err := strconv.Atoi(matrix[i][j]); err == nil {
			number = number*10 + digit
		}

	}

	return number
}

func CheckSymbol(matrix [][]string, indexes [][]int) (bool, int) {
	// Define a function to check if a character is a symbol
	isSymbol := func(char string) bool {
		// symbol is a digit
		return char == "!" || char == "@" || char == "#" || char == "$" || char == "%" || char == "^" || char == "&" || char == "*" || char == "(" || char == ")" || char == "-" || char == "+" || char == "=" || char == "_" || char == "/" || char == "~"
	}

	for _, indexPair := range indexes {
		for i := indexPair[0] - 1; i <= indexPair[0]+1; i++ {
			for j := indexPair[1] - 1; j <= indexPair[1]+1; j++ {
				if isValidIndex(matrix, i, j) && isSymbol(matrix[i][j]) {
					return true, buildNumber(matrix, indexes)
				}
			}
		}
	}

	return false, buildNumber(matrix, indexes)
}
