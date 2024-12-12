package day4

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day4",
	Short: "day4",
	Long:  "day4",
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

func countXMAS(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// check if XMAS exists in a specific direction
	checkDirection := func(x, y, dx, dy int) bool {
		word := ""
		for i := 0; i < 4; i++ { // "XMAS" is 4 characters
			if x < 0 || y < 0 || x >= rows || y >= cols {
				return false
			}
			word += string(grid[x][y])
			x += dx
			y += dy
		}
		return word == "XMAS"
	}

	// Directions: (dx, dy) -> (row change, column change)
	directions := [][2]int{
		{0, 1},   // Right
		{0, -1},  // Left
		{1, 0},   // Down
		{-1, 0},  // Up
		{1, 1},   // Diagonal Down-Right
		{1, -1},  // Diagonal Down-Left
		{-1, 1},  // Diagonal Up-Right
		{-1, -1}, // Diagonal Up-Left
	}

	// Iterate over each cell in the grid
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if checkDirection(x, y, dx, dy) {
					count++
				}
			}
		}
	}

	return count
}

func part1(s string) int64 {
	return int64(countXMAS(strings.Split(s, "\n")))
}

func part2(s string) int64 {
	return 0
}
