package tests

import (
	"testing"

	"adventofcode.com/m/problems"
)

func Test_empty_matrix(t *testing.T) {
	matrix := [][]string{{".", "1", "4", "9", "*", "2", "2", "7"}, {".", "1", "4", "9", "*", "2", "2", "7"}}

	occurrences := problems.FindNumbersOccurrences(matrix)

	for _, occurrence := range occurrences {
		_, number := problems.CheckSymbol(matrix, occurrence)

		if number != 0 {
			t.Errorf("Expected 0, got %d", number)
		}
	}

}
