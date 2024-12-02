package day21

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day21",
	Short: "day21",
	Long:  "day21",
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

func part1(s string) int64 {
	var score int64
	for _, line := range strings.Split(s, "\n") {
		fmt.Println(line)
	}
	return score
}

func part2(s string) int64 {
	return 0
}
