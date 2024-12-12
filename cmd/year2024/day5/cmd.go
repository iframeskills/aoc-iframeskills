package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day5",
	Short: "day5",
	Long:  "day5",
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

func parseInput(input string) (map[int]map[int]bool, [][]int) {
	lines := strings.Split(input, "\n")
	var rules = make(map[int]map[int]bool)
	var updates [][]int

	// Process rules and updates
	readingRules := true
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			readingRules = false
			continue
		}
		if readingRules {
			// Parsing rules
			parts := strings.Split(line, "|")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			if _, exists := rules[x]; !exists {
				rules[x] = make(map[int]bool)
			}
			rules[x][y] = true
		} else {
			// Parsing updates
			parts := strings.Split(line, ",")
			update := make([]int, len(parts))
			for i, p := range parts {
				update[i], _ = strconv.Atoi(p)
			}
			updates = append(updates, update)
		}
	}
	return rules, updates
}

func isValidUpdate(rules map[int]map[int]bool, update []int) bool {
	// For each update, we must check the ordering constraints between pages
	pageIndex := make(map[int]int)
	for i, page := range update {
		pageIndex[page] = i
	}

	// Validate the order using the rules
	for x, successors := range rules {
		for y := range successors {
			// If both x and y are in the update, check if x comes before y
			if _, xInUpdate := pageIndex[x]; xInUpdate {
				if _, yInUpdate := pageIndex[y]; yInUpdate {
					if pageIndex[x] > pageIndex[y] {
						// Violation found, x must come before y
						return false
					}
				}
			}
		}
	}
	return true
}

func findMiddlePage(update []int) int {
	// Find the middle page number
	n := len(update)
	if n%2 == 1 {
		return update[n/2]
	}
	return update[n/2-1]
}

func part1(s string) int64 {
	// Parse input
	rules, updates := parseInput(s)

	// Process each update and find valid ones
	totalMiddlePages := 0
	for _, update := range updates {
		if isValidUpdate(rules, update) {
			middlePage := findMiddlePage(update)
			fmt.Printf("Middle page number from valid update: %d\n", middlePage)
			totalMiddlePages += middlePage
		}
	}

	// Output result
	fmt.Printf("Total of middle page numbers from valid updates: %d\n", totalMiddlePages)

	return int64(totalMiddlePages)
}

func part2(s string) int64 {
	return 0
}
