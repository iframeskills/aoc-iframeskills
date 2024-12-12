package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day3",
	Short: "day3",
	Long:  "day3",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {

	logrus.Infof("hi")

	b, err := os.ReadFile(fmt.Sprintf("cmd/year%s/%s/1.txt", parent, command))
	if err != nil {
		logrus.Fatalf("error reading input: %v", err)
	}

	logrus.Infof("score part1: %d", part1(string(b)))
	logrus.Infof("score part2: %d", part2(string(b)))
}

func mul(a, b int) int {
	return a * b
}

func getValidInstructions(input string) []string {
	// Define the regex pattern
	pattern := `mul\(\d+,\d+\)`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Find all matches
	return re.FindAllString(input, -1)
}

func getScore(input string) int64 {
	var score int64
	for _, line := range strings.Split(input, "\n") {
		fmt.Println(line)
		validInstructions := getValidInstructions(line)

		// Loop over each match
		for _, instruction := range validInstructions {

			// Define the regex to extract numbers
			pattern := `mul\((\d+),(\d+)\)`
			re := regexp.MustCompile(pattern)

			// Extract numbers using capturing groups
			matches := re.FindStringSubmatch(instruction)

			if matches != nil {
				// Convert the captured strings to integers
				x, _ := strconv.Atoi(matches[1]) // First number
				y, _ := strconv.Atoi(matches[2]) // Second number

				// Perform multiplication
				multipliedresult := mul(x, y)
				score += int64(multipliedresult)

				fmt.Printf("Multiplication of %d and %d is %d\n", x, y, multipliedresult)
			} else {
				fmt.Println("Invalid input format.")
			}
		}

	}
	return score
}

func part1(s string) int64 {
	return getScore(s)
}

func part2(s string) int64 {
	// Step 1: Concatenate all strings
	concatenated := strings.Join(strings.Split(s, "\n"), "")

	// Step 2: Process the instructions
	var result strings.Builder
	include := true // Start in "include" mode

	for i := 0; i < len(concatenated); i++ {
		if strings.HasPrefix(concatenated[i:], "do()") {
			// Found "do()" - switch to "include" mode and append "do()"
			include = true
			result.WriteString("do()")
			i += len("do()") - 1 // Skip the length of "do()"
		} else if strings.HasPrefix(concatenated[i:], "don't()") {
			// Found "don't()" - switch to "exclude" mode and append "don't()"
			include = false
			result.WriteString("don't()")
			i += len("don't()") - 1 // Skip the length of "don't()"
		} else if include {
			// Only append characters if in "include" mode
			result.WriteByte(concatenated[i])
		}
	}

	return getScore(result.String())
}
