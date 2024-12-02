package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day2",
	Short: "day2",
	Long:  "day2",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf("cmd/year%s/%s/1.txt", parent, command))
	if err != nil {
		logrus.Fatalf("error reading input: %v", err)
	}

	logrus.Infof("score part1: %d", part1(string(b)))
	logrus.Infof("score part2: %d", part2(string(b)))
}

// Check if a report is safe based on the given rules:
// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
func isSafe(report []int) bool {
	var direction int = 0
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff < -3 || diff > 3 || diff == 0 {
			// Any invalid difference immediately makes the report unsafe
			return false
		}
		if direction == 0 {
			// Set the direction
			direction = diff
		} else if (direction > 0 && diff < 0) || (direction < 0 && diff > 0) {
			// Direction mismatch makes the report unsafe
			return false
		}
	}
	return true
}

// Check if removing a single level makes the report safe
func isSafeWithRemoval(report []int) bool {

	// Iterate through each level and check if removing it makes the report safe
	for i := 0; i < len(report); i++ {
		// Create a copy of the report with the i-th element removed
		modifiedReport := append(report[:i:i], report[i+1:]...)

		// Check if the modified report is safe
		if isSafe(modifiedReport) {
			return true
		}
	}

	// If no single removal makes the report safe, return false
	return false
}
func part1(s string) int64 {
	var score int64
	for _, line := range strings.Split(s, "\n") {
		// Parse the line into a slice of integers
		fields := strings.Fields(line)
		report := make([]int, len(fields))
		for i, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Printf("Invalid number in report: %s\n", field)
			}
			report[i] = num
		}

		if isSafe(report) {
			score++
		}
	}
	return score
}

func part2(s string) int64 {

	var score int64
	for _, line := range strings.Split(s, "\n") {
		// Parse the line into a slice of integers
		fields := strings.Fields(line)
		report := make([]int, len(fields))
		for i, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Printf("Invalid number in report: %s\n", field)
			}
			report[i] = num
		}

		if isSafeWithRemoval(report) {
			score++
		}
	}
	return score
}
