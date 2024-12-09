package adventofcode

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bobg/errors"
	"github.com/fatih/color"
)

type RunFunc func(input string) (int, error)

func Part1(year, day int, exampleAnswer int, runFunc RunFunc) {
	Run(year, day, 1, exampleAnswer, runFunc)
}

func Part2(year, day int, exampleAnswer int, runFunc RunFunc) {
	Run(year, day, 2, exampleAnswer, runFunc)
}

func Run(year, day, part int, exampleAnswer int, runFunc RunFunc) {
	if err := run(year, day, part, exampleAnswer, runFunc); err != nil {
		color.Red("error: %v", err)
		os.Exit(1)
	}
}

func run(year, day, part int, exampleAnswer int, runFunc RunFunc) error {
	yearDirName := fmt.Sprintf("year%d", year)
	dayDirName := fmt.Sprintf("day%02d", day)
	directoryPath := filepath.Join(yearDirName, dayDirName)

	var exampleInput []byte
	var err error
	if part == 2 {
		if exampleInput, err = os.ReadFile(filepath.Join(directoryPath, "part2.example.input.txt")); err != nil {
			if !os.IsNotExist(err) {
				return errors.Wrapf(err, "reading example input for part %d", part)
			}
		}
	}

	if exampleInput == nil {
		if exampleInput, err = os.ReadFile(filepath.Join(directoryPath, "example.input.txt")); err != nil {
			return errors.Wrapf(err, "reading example input for part %d", part)
		}
	}

	exampleResult, err := runFunc(strings.TrimSuffix(string(exampleInput), "\n"))
	if err != nil {
		return errors.Wrapf(err, "running part %d example implementation", part)
	}

	if exampleAnswer != exampleResult {
		return fmt.Errorf("incorrect result for part %d: got %d, want %d", part, exampleResult, exampleAnswer)
	}

	input, err := os.ReadFile(filepath.Join(directoryPath, "input.txt"))
	if err != nil {
		return errors.Wrapf(err, "reading example input for part %d", part)
	}

	result, err := runFunc(strings.TrimSuffix(string(input), "\n"))
	if err != nil {
		return errors.Wrapf(err, "running part %d implementation", part)
	}

	color.Green("part %d answer: %d\n", part, result)
	return nil
}
