package year2023

import (
	"aoc/cmd/year2023/day1"
	"aoc/cmd/year2023/day10"
	"aoc/cmd/year2023/day11"
	"aoc/cmd/year2023/day12"
	"aoc/cmd/year2023/day13"
	"aoc/cmd/year2023/day14"
	"aoc/cmd/year2023/day15"
	"aoc/cmd/year2023/day16"
	"aoc/cmd/year2023/day17"
	"aoc/cmd/year2023/day18"
	"aoc/cmd/year2023/day19"
	"aoc/cmd/year2023/day2"
	"aoc/cmd/year2023/day20"
	"aoc/cmd/year2023/day21"
	"aoc/cmd/year2023/day22"
	"aoc/cmd/year2023/day23"
	"aoc/cmd/year2023/day24"
	"aoc/cmd/year2023/day25"
	"aoc/cmd/year2023/day4"
	"aoc/cmd/year2023/day5"
	"aoc/cmd/year2023/day6"
	"aoc/cmd/year2023/day7"
	"aoc/cmd/year2023/day8"
	"aoc/cmd/year2023/day9"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2023",
	Short: "2023",
	Long:  "2023 is a command line utility for solving Advent of Code 2023 puzzles.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here
	},
}

func init() {
	Cmd.AddCommand(day1.Cmd)
	Cmd.AddCommand(day2.Cmd)
	Cmd.AddCommand(day4.Cmd)
	Cmd.AddCommand(day5.Cmd)
	Cmd.AddCommand(day6.Cmd)
	Cmd.AddCommand(day7.Cmd)
	Cmd.AddCommand(day8.Cmd)
	Cmd.AddCommand(day9.Cmd)
	Cmd.AddCommand(day10.Cmd)
	Cmd.AddCommand(day11.Cmd)
	Cmd.AddCommand(day12.Cmd)
	Cmd.AddCommand(day13.Cmd)
	Cmd.AddCommand(day14.Cmd)
	Cmd.AddCommand(day15.Cmd)
	Cmd.AddCommand(day16.Cmd)
	Cmd.AddCommand(day17.Cmd)
	Cmd.AddCommand(day18.Cmd)
	Cmd.AddCommand(day19.Cmd)
	Cmd.AddCommand(day20.Cmd)
	Cmd.AddCommand(day21.Cmd)
	Cmd.AddCommand(day22.Cmd)
	Cmd.AddCommand(day23.Cmd)
	Cmd.AddCommand(day24.Cmd)
	Cmd.AddCommand(day25.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
