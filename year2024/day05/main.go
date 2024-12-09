package main

import (
	"sort"
	"strconv"
	"strings"

	adventofcode "github.com/broothie/advent-of-code"
	"github.com/samber/lo"
)

type PageOrderingRule struct {
	Before int
	After  int
}

type Update []int

func (u Update) canBePrinted(pageOrderingRules []PageOrderingRule) bool {
	indexes := make(map[int]int)
	for index, pageNumber := range u {
		indexes[pageNumber] = index
	}

	return lo.EveryBy(pageOrderingRules, func(pageOrderingRule PageOrderingRule) bool {
		indexOfBefore, found := indexes[pageOrderingRule.Before]
		if !found {
			return true
		}

		indexOfAfter, found := indexes[pageOrderingRule.After]
		if !found {
			return true
		}

		return indexOfBefore < indexOfAfter
	})
}

func (u Update) middlePageNumber() int {
	return u[len(u)/2]
}

func main() {
	adventofcode.Part1(2024, 5, 143, func(input string) (int, error) {
		pageOrderingRules, updates, err := parse(input)
		if err != nil {
			return 0, err
		}

		printedUpdates := lo.Filter(updates, func(update Update, _ int) bool {
			return update.canBePrinted(pageOrderingRules)
		})

		return lo.SumBy(printedUpdates, func(update Update) int {
			return update.middlePageNumber()
		}), nil
	})

	adventofcode.Part2(2024, 5, 123, func(input string) (int, error) {
		pageOrderingRules, updates, err := parse(input)
		if err != nil {
			return 0, err
		}

		pageOrderingExists := make(map[[2]int]bool)
		for _, pageOrderingRule := range pageOrderingRules {
			before := pageOrderingRule.Before
			after := pageOrderingRule.After

			pageOrderingExists[[2]int{before, after}] = true
		}

		unprintedUpdates := lo.Filter(updates, func(update Update, _ int) bool {
			return !update.canBePrinted(pageOrderingRules)
		})

		for _, update := range unprintedUpdates {
			sort.Slice(update, func(i, j int) bool {
				iPageNumber := update[i]
				jPageNumber := update[j]

				return pageOrderingExists[[2]int{jPageNumber, iPageNumber}] // Purposefully swapped i and j
			})
		}

		return lo.SumBy(unprintedUpdates, func(update Update) int {
			return update.middlePageNumber()
		}), nil
	})
}

func parse(input string) ([]PageOrderingRule, []Update, error) {
	sections := strings.Split(input, "\n\n")
	pageOrderingRulesString := sections[0]
	updatesString := sections[1]

	var pageOrderingRules []PageOrderingRule
	for _, pageOrderingRuleString := range strings.Split(pageOrderingRulesString, "\n") {
		segments := strings.Split(pageOrderingRuleString, "|")

		before, err := strconv.Atoi(segments[0])
		if err != nil {
			return nil, nil, err
		}

		after, err := strconv.Atoi(segments[1])
		if err != nil {
			return nil, nil, err
		}

		pageOrderingRules = append(pageOrderingRules, PageOrderingRule{
			Before: before,
			After:  after,
		})
	}

	var updates []Update
	for _, updateString := range strings.Split(updatesString, "\n") {
		var update Update
		for _, pageNumberString := range strings.Split(updateString, ",") {
			pageNumber, err := strconv.Atoi(pageNumberString)
			if err != nil {
				return nil, nil, err
			}

			update = append(update, pageNumber)
		}

		updates = append(updates, update)
	}
	return pageOrderingRules, updates, nil
}
