package tests

import (
	"testing"

	"adventofcode.com/m/problems"
)

func Test_single_digit(t *testing.T) {
	input := "threeseven286fourfour"
	expected := 34
	result := problems.Make_numbers(input)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func Test_change_input(t *testing.T) {
	input := "threeight286fourfour"
	expected := "34"
	result := problems.ChangeInput(input)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func Test_Reverse_input(t *testing.T) {
	input := "threeight286fourfour"
	expected := "ruofruof682thgieerht"
	result := problems.ReverseString(input)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
