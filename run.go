package adventofcode

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"

	"github.com/bobg/errors"
)

var puzzleRegexp = regexp.MustCompile(`year(\d{4})/day(\d{2})/part(\d)`)

type RunFunc func(input string) (int, error)

func Run(exampleAnswer int, runFunc RunFunc) {
	if err := run(exampleAnswer, runFunc); err != nil {
		log.Fatalln(err)
	}
}

func run(exampleAnswer int, runFunc RunFunc) error {
	_, filePath, _, ok := runtime.Caller(2) // Skips `run()` and `Run()`
	if !ok {
		return errors.Errorf("getting caller")
	}

	year, day, part, err := parseFilePath(filePath)
	if err != nil {
		return errors.Wrapf(err, "parsing filepath %q", filePath)
	}

	if err := execute(year, day, part, exampleAnswer, runFunc); err != nil {
		return errors.Wrapf(err, "executing puzzle year %d day %d part %d", year, day, part)
	}

	return nil
}

func execute(year int, day int, part int, exampleAnswer int, runFunc RunFunc) error {
	pathPrefix := fmt.Sprintf("year%d/day%02d/part%d", year, day, part)

	exampleInput, err := os.ReadFile(filepath.Join(pathPrefix, "example.txt"))
	if err != nil {
		return errors.Wrapf(err, "reading example input file")
	}

	exampleResult, err := runFunc(string(exampleInput))
	if err != nil {
		return errors.Wrapf(err, "running example puzzle")
	}

	if exampleResult == exampleAnswer {
		fmt.Println("example correct")
	} else {
		return fmt.Errorf("incorrect result for example: got %d, want %d", exampleResult, exampleAnswer)
	}

	input, err := os.ReadFile(filepath.Join(pathPrefix, "input.txt"))
	if err != nil {
		return errors.Wrapf(err, "reading input file")
	}

	result, err := runFunc(string(input))
	if err != nil {
		return errors.Wrapf(err, "running puzzle")
	}

	fmt.Printf("year %d day %d part %d answer: %d\n", year, day, part, result)
	return nil
}

func parseFilePath(filePath string) (int, int, int, error) {
	matches := puzzleRegexp.FindStringSubmatch(filePath)
	if len(matches) == 0 {
		return 0, 0, 0, errors.Errorf("filename %q does not match puzzle regexp", filePath)
	}

	yearMatch, dayMatch, partMatch := matches[1], matches[2], matches[3]

	year, err := strconv.Atoi(yearMatch)
	if err != nil {
		return 0, 0, 0, errors.Wrapf(err, "parsing year %q", yearMatch)
	}

	day, err := strconv.Atoi(dayMatch)
	if err != nil {
		return 0, 0, 0, errors.Wrapf(err, "parsing day %q", dayMatch)
	}

	part, err := strconv.Atoi(partMatch)
	if err != nil {
		return 0, 0, 0, errors.Wrapf(err, "parsing part %q", partMatch)
	}

	return year, day, part, nil
}
