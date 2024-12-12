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
	// Ensure the update is not empty
	if len(update) == 0 {
		panic("Attempted to find the middle page of an empty update")
	}

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

func contains(arr []int, x int) bool {
	for _, val := range arr {
		if val == x {
			return true
		}
	}
	return false
}

// TODO:  print out the status of the queue,
// in-degree of each page, and the current state of sortedPages
// to identify where things are going wrong.

func topologicalSort(rules map[int]map[int]bool, update []int) []int {
	// Step 1: Calculate the in-degree of each page
	inDegree := make(map[int]int)
	for page := range rules {
		inDegree[page] = 0 // Initialize in-degree for all pages
	}

	// Update in-degree for each page based on the rules
	for _, successors := range rules {
		for successor := range successors {
			inDegree[successor]++
		}
	}

	// Step 2: Create a queue for pages with in-degree 0 (can be printed first)
	queue := list.New()
	for page, degree := range inDegree {
		if degree == 0 && contains(update, page) {
			queue.PushBack(page)
		}
	}

	// Step 3: Perform the topological sort using Kahn's algorithm
	sortedPages := []int{}
	for queue.Len() > 0 {
		page := queue.Remove(queue.Front()).(int)

		// Ensure that the page is part of the update
		if contains(update, page) {
			sortedPages = append(sortedPages, page)
		}

		// Reduce in-degree of neighbors and add them to the queue if their in-degree becomes 0
		if successors, exists := rules[page]; exists {
			for successor := range successors {
				inDegree[successor]--
				if inDegree[successor] == 0 && contains(update, successor) {
					queue.PushBack(successor)
				}
			}
		}
	}

	if len(sortedPages) != len(update) {
		panic("not all pages were sorted")
	}

	// Return sorted pages (valid order)
	return sortedPages
}

func part2(s string) int64 {
	// Parse input
	rules, updates := parseInput(s)

	// Process each update and reorder if necessary
	totalMiddlePages := 0
	for _, update := range updates {
		// If the update is not valid, reorder it using topological sort
		if !isValidUpdate(rules, update) {
			orderedUpdate := topologicalSort(rules, update)

			middlePage := findMiddlePage(orderedUpdate)
			fmt.Printf("Middle page number from correctly ordered update: %d\n", middlePage)
			totalMiddlePages += middlePage
		}
	}

	// Output result
	fmt.Printf("Total of middle page numbers from reordered updates: %d\n", totalMiddlePages)

	return int64(totalMiddlePages)
}
