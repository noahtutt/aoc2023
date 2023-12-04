package main

import (
	"AoC/pkg/input"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func positionsToCheck(line int, match []int) (ret [][]int) {

	// lines above and below
	for i := match[0] - 1; i <= match[1]; i++ {
		ret = append(ret, []int{line - 1, i})
		ret = append(ret, []int{line + 1, i})
	}

	// current line
	ret = append(ret, []int{line, match[0] - 1})
	ret = append(ret, []int{line, match[1]})

	return
}

func checkPosition(pos []int, lines []string, numLines int, lineLength int) bool {
	if pos[0] < 0 || numLines <= pos[0] || pos[1] < 0 || lineLength <= pos[1] {
		return false
	}

	re := regexp.MustCompile(`[^\d.]`)
	return re.MatchString(string(lines[pos[0]][pos[1]]))
}

func main() {

	f, err := input.GetInputFile("3.1", "input.txt")
	if err != nil {
		panic(err)
	}
	defer input.DeferFunc(f)

	// read all the lines into memory for lookahead/lookback
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	numLines := len(lines)
	lineLength := len(lines[0])

	var sum int

	re := regexp.MustCompile(`\d+`)
	for i, line := range lines {
		matches := re.FindAllStringIndex(line, -1)
		for _, match := range matches {
			positions := positionsToCheck(i, match)
			for _, pos := range positions {
				if checkPosition(pos, lines, numLines, lineLength) {
					num, err := strconv.Atoi(line[match[0]:match[1]])
					if err != nil {
						panic(err)
					}
					sum = sum + num
					break
				}
			}
		}
	}

	fmt.Printf("finished with sum of %v", sum)
}
