package day13

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day13",
	Short: "day13",
	Long:  "day13",
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

// Machine represents a claw machine's configuration.
type Machine struct {
	Ax, Ay, Bx, By, Px, Py int64
}

// SolveMachine tries to find the minimum cost to win the prize for a given machine.
// Returns cost, and whether it's possible to win.
func SolveMachine(machine Machine, maxPresses int64) (int64, bool) {
	minCost := int64(math.MaxInt64)
	possible := false

	// Iterate over all possible presses for button A (n)
	for n := int64(0); n <= maxPresses; n++ {
		// Iterate over all possible presses for button B (m)
		for m := int64(0); m <= maxPresses; m++ {
			// Calculate resulting X and Y positions
			x := n*machine.Ax + m*machine.Bx
			y := n*machine.Ay + m*machine.By

			// Check if this aligns with the prize
			if x == machine.Px && y == machine.Py {
				// Calculate cost: 3 tokens for A presses, 1 token for B presses
				cost := 3*n + 1*m
				if cost < minCost {
					minCost = cost
					possible = true
				}
			}
		}
	}

	return minCost, possible
}

// ParseInput reads the input string and converts it into a slice of Machines.
func ParseInput(input string) ([]Machine, error) {
	re := regexp.MustCompile(`[-+]?\d+`)
	var machines []Machine
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		lineA := scanner.Text()
		if !strings.HasPrefix(lineA, "Button A") {
			continue
		}
		// Parse Button A
		valuesA := re.FindAllString(lineA, -1)
		ax, _ := strconv.ParseInt(valuesA[0], 10, 64)
		ay, _ := strconv.ParseInt(valuesA[1], 10, 64)

		scanner.Scan()
		lineB := scanner.Text()
		// Parse Button B
		valuesB := re.FindAllString(lineB, -1)
		bx, _ := strconv.ParseInt(valuesB[0], 10, 64)
		by, _ := strconv.ParseInt(valuesB[1], 10, 64)

		scanner.Scan()
		lineP := scanner.Text()
		// Parse Prize
		valuesP := re.FindAllString(lineP, -1)
		px, _ := strconv.ParseInt(valuesP[0], 10, 64)
		py, _ := strconv.ParseInt(valuesP[1], 10, 64)

		machines = append(machines, Machine{Ax: ax, Ay: ay, Bx: bx, By: by, Px: px, Py: py})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return machines, nil
}

func part1(s string) int64 {

	machines, err := ParseInput(s)
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		panic(err)
	}

	maxPresses := int64(100)
	totalCost := int64(0)
	numPrizes := 0

	// Solve for each machine
	for i, machine := range machines {
		cost, possible := SolveMachine(machine, maxPresses)
		if possible {
			fmt.Printf("Machine %d: Prize won with cost %d\n", i+1, cost)
			totalCost += cost
			numPrizes++
		} else {
			fmt.Printf("Machine %d: Prize cannot be won\n", i+1)
		}
	}

	fmt.Printf("\nTotal prizes won: %d\n", numPrizes)
	fmt.Printf("Minimum total cost: %d\n", totalCost)

	return int64(totalCost)
}

// AdjustPrizes adds an offset to the X and Y coordinates of each prize.
func AdjustPrizes(machines []Machine, offset int64) []Machine {
	adjustedMachines := make([]Machine, len(machines))
	for i, machine := range machines {
		adjustedMachines[i] = Machine{
			Ax: machine.Ax, Ay: machine.Ay,
			Bx: machine.Bx, By: machine.By,
			Px: int64(machine.Px) + offset, Py: int64(machine.Py) + offset,
		}

		// log.Println(fmt.Sprintf("Machine{Ax: %d, Ay: %d, Bx: %d, By: %d, Px: %d, Py: %d}", machine.Ax, machine.Ay, machine.Bx, machine.By, machine.Px+offset, machine.Py+offset))
	}
	return adjustedMachines
}

func part2(s string) int64 {

	machines, err := ParseInput(s)
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		panic(err)
	}

	maxPresses := int64(10000000000000)
	totalCost := int64(100)
	numPrizes := 0

	// TODO: obviously way too inefficient, needs to be optimised...
	adjustedMachines := AdjustPrizes(machines, 10000000000000)
	// Solve for each machine
	for i, machine := range adjustedMachines {
		cost, possible := SolveMachine(machine, maxPresses)
		if possible {
			fmt.Printf("Machine %d: Prize won with cost %d\n", i+1, cost)
			totalCost += cost
			numPrizes++
		} else {
			fmt.Printf("Machine %d: Prize cannot be won\n", i+1)
		}
	}

	fmt.Printf("\nTotal prizes won: %d\n", numPrizes)
	fmt.Printf("Minimum total cost: %d\n", totalCost)

	return int64(totalCost)
}
