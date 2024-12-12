package day5

import (
	"container/list"
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
	// Create a map for the in-degrees of each page
	inDegree := make(map[int]int)
	for page := range rules {
		inDegree[page] = 0
	}

	// Count in-degrees based on the rules
	for _, successors := range rules {
		for successor := range successors {
			inDegree[successor]++
		}
	}

	// Now, perform Kahn's algorithm (Topological sort)
	queue := list.New()
	for page, degree := range inDegree {
		if degree == 0 {
			queue.PushBack(page)
		}
	}

	sortedPages := []int{}
	for queue.Len() > 0 {
		page := queue.Remove(queue.Front()).(int)
		// Ensure the page is in the update
		if contains(update, page) {
			sortedPages = append(sortedPages, page)
		}

		// Reduce in-degree of successors
		if successors, exists := rules[page]; exists {
			for successor := range successors {
				inDegree[successor]--
				if inDegree[successor] == 0 {
					queue.PushBack(successor)
				}
			}
		}
	}

	// Check if we processed all pages in the update
	if len(sortedPages) == len(update) {
		// Ensure that the update respects the topological order
		pageIndex := make(map[int]int)
		for i, page := range sortedPages {
			pageIndex[page] = i
		}

		// Validate the update's order against the topological order
		for i := 0; i < len(update)-1; i++ {
			if pageIndex[update[i]] > pageIndex[update[i+1]] {
				return false
			}
		}

		return true
	}
	return false
}

func contains(arr []int, x int) bool {
	for _, val := range arr {
		if val == x {
			return true
		}
	}
	return false
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
