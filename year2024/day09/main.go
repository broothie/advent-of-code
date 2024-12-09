package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bobg/errors"
	adventofcode "github.com/broothie/advent-of-code"
	"github.com/samber/lo"
)

func main() {
	adventofcode.Part1(2024, 9, 1928, func(input string) (int, error) {
		blocks, err := parsePart1(input)
		if err != nil {
			return 0, errors.Wrap(err, "parsing")
		}

		cursor := 0
		fileBlockCursor := len(blocks) - 1
		for cursor < fileBlockCursor {
			for {
				if blocks[cursor] == -1 {
					break
				}

				cursor++
			}

			for {
				if blocks[fileBlockCursor] != -1 {
					break
				}

				fileBlockCursor--
			}

			if cursor >= fileBlockCursor {
				break
			}

			blocks[cursor] = blocks[fileBlockCursor]
			blocks[fileBlockCursor] = -1
		}

		sum := 0
		for index, fileID := range blocks {
			if fileID != -1 {
				sum += index * fileID
			}
		}

		return sum, nil
	})

	adventofcode.Part2(2024, 9, 2858, func(input string) (int, error) {
		blocks, err := parsePart2(input)
		if err != nil {
			return 0, errors.Wrap(err, "parsing")
		}

		for fileBlockCursor := len(blocks) - 1; fileBlockCursor > 0; fileBlockCursor-- {
			fileBlock := blocks[fileBlockCursor]
			if len(fileBlock) == 0 || fileBlock[0] == -1 {
				continue
			}

			var freeBlock []int
			freeBlockFound := false
			for freeBlockCursor := 0; freeBlockCursor < fileBlockCursor; freeBlockCursor++ {
				block := blocks[freeBlockCursor]
				if lo.Count(block, -1) >= len(fileBlock) {
					freeBlock = block
					freeBlockFound = true
					break
				}
			}
			if !freeBlockFound {
				continue
			}

			copy(freeBlock[lo.IndexOf(freeBlock, -1):], fileBlock)
			blocks[fileBlockCursor] = filledSlice(-1, len(fileBlock))
		}

		sum := 0
		for index, fileID := range lo.Flatten(blocks) {
			if fileID != -1 {
				sum += index * fileID
			}
		}

		return sum, nil
	})
}

func parsePart1(input string) ([]int, error) {
	freeSpace := 0
	fileID := 0
	var blocks []int
	for index, character := range input {
		blockSize, err := strconv.Atoi(string(character))
		if err != nil {
			return nil, errors.Wrap(err, "converting character")
		}

		if index%2 == 0 {
			// even: file
			for i := 0; i < blockSize; i++ {
				blocks = append(blocks, fileID)
			}

			fileID++
		} else {
			// odd: free space
			for i := 0; i < blockSize; i++ {
				blocks = append(blocks, -1)
				freeSpace++
			}
		}
	}

	return blocks, nil
}

func parsePart2(input string) ([][]int, error) {
	fileID := 0
	var blocks [][]int
	for index, character := range input {
		blockSize, err := strconv.Atoi(string(character))
		if err != nil {
			return nil, errors.Wrap(err, "converting character")
		}

		if index%2 == 0 {
			// even: file
			blocks = append(blocks, filledSlice(fileID, blockSize))

			fileID++
		} else {
			// odd: free space
			blocks = append(blocks, filledSlice(-1, blockSize))
		}
	}

	return blocks, nil
}

func filledSlice(value int, length int) []int {
	result := make([]int, length)
	for index := range result {
		result[index] = value
	}

	return result
}

func printBlocks(blocks []int) {
	strs := lo.Map(blocks, func(fileID int, _ int) string {
		if fileID == -1 {
			return "."
		}

		return strconv.Itoa(fileID)
	})

	fmt.Println(strings.Join(strs, ""))
}
