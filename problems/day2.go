package problems

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Cubes struct {
	gameId int
	red    int
	blue   int
	green  int
}

func Day2() int {
	const file_path = "./problems/inputs/day2_input.txt"

	// read line by line from the file
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		power := ParseLine(line, &sum)
		sum += power

	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return 0

}

func ParseLine(line string, sum *int) int {
	// Use regular expression to extract the gameId and cube information
	re := regexp.MustCompile(`Game (\d+): (.+)$`)
	matches := re.FindAllStringSubmatch(line, -1)

	if len(matches) == 0 {
		return 0
	}

	//gameId, _ := strconv.Atoi(matches[0][1])

	// Split the cube information by ';'
	rounds := strings.Split(matches[0][2], ";")
	var red, blue, green int

	// Iterate through each round and accumulate cube counts
	for _, round := range rounds {
		round = strings.TrimSpace(round)
		if round != "" {
			// Split each round by ','
			cubeCounts := strings.Split(round, ",")
			// each round
			for _, cubeCount := range cubeCounts {
				cubeCount = strings.TrimSpace(cubeCount)

				// Extract the color and count
				parts := strings.SplitN(cubeCount, " ", 2)
				if len(parts) == 2 {
					count, _ := strconv.Atoi(parts[0])
					color := strings.ToLower(strings.TrimSpace(parts[1]))

					// valid := verifyIfOverflows(color, count)

					// if !valid {
					// 	return 0
					// }

					switch color {
					case "red":
						red = max(red, count)
					case "blue":
						blue = max(blue, count)
					case "green":
						green = max(green, count)
					}

				}

			}

		}
	}

	return red * blue * green
}

// 12 red cubes, 13 green cubes, and 14 blue cubes
func verifyIfOverflows(color string, count int) bool {
	if color == "red" && count > 12 {
		return false
	}

	if color == "blue" && count > 14 {
		return false
	}

	if color == "green" && count > 13 {
		return false
	}

	return true
}
