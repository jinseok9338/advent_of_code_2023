package problems

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func Day1() int {
	const file_path = "./problems/inputs/day1_input.txt"

	// read line by line from the file
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Process each line here
		number := Make_numbers(line)
		result += number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	// combine first number and last number to make 2 digit number
	return result
}

func Make_numbers(input string) int {
	var firstNum, lastNum int
	numCount := 0

	// if the line contains {"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",}
	// replace the number with it's corresponding value
	// threehqv2 -> 3hqv2
	// eightwothree -> 8wo3
	input = ChangeInput(input)
	for _, char := range input {
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			numCount++
			if numCount == 1 {
				firstNum = num
			}
			lastNum = num
		}
	}

	if numCount == 1 {
		// If there is only one number, use it to create a 2-digit number
		return firstNum*10 + firstNum
	} else if numCount >= 2 {
		// If there are two or more numbers, use the first and last numbers
		return firstNum*10 + lastNum
	}

	// Return 0 if no numbers are found
	return 0
}

func ChangeInput(input string) string {
	re := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9|0`)

	re_reverse := regexp.MustCompile(`eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|1|2|3|4|5|6|7|8|9|0`)

	var first string
	var last string

	matches := re.FindAllStringIndex(input, -1)

	for _, match := range matches {
		// if the match is a number add the number
		// if the match is number in string then convert it to the wordToDigit then append

		var match_string = input[match[0]:match[1]]
		if unicode.IsDigit(rune(match_string[0])) {
			first = match_string
			break
		} else {
			first = wordToDigit(match_string)
			break
		}
	}

	// reverse input
	reverse_input := ReverseString(input)
	matches_reverse := re_reverse.FindAllStringIndex(reverse_input, -1)

	for _, match_reverse := range matches_reverse {
		// if the match is a number add the number
		// if the match is number in string then convert it to the wordToDigit then append
		var match_string = reverse_input[match_reverse[0]:match_reverse[1]]
		if unicode.IsDigit(rune(match_string[0])) {
			last = match_string
			break
		} else {
			last = wordToDigit(match_string)
			break
		}
	}

	return first + last
}

func wordToDigit(word string) string {
	switch word {
	case "eno":
		return "1"
	case "owt":
		return "2"
	case "eerht":
		return "3"
	case "ruof":
		return "4"
	case "evif":
		return "5"
	case "xis":
		return "6"
	case "neves":
		return "7"
	case "thgie":
		return "8"
	case "enin":
		return "9"
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return "0"

	}

}

func ReverseString(input string) string {
	// Convert the string to a rune slice
	runes := []rune(input)

	// Get the length of the rune slice
	length := len(runes)

	// Reverse the order of the runes
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the reversed rune slice back to a string
	return string(runes)
}
