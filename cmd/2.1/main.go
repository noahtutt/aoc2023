package main

import (
	"AoC/pkg/input"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var maxByColor = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {

	f, err := input.GetInputFile("2.1", "input.txt")
	if err != nil {
		panic(err)
	}
	defer input.DeferFunc(f)

	var sum int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		splitLine := strings.Split(line, ":")
		id, err := strconv.Atoi(strings.Split(splitLine[0], " ")[1])
		if err != nil {
			panic("invalid id in " + splitLine[0])
		}

		claims := splitLine[1]
		subClaims := strings.Split(claims, ";")
		var statements []string
		for _, subClaim := range subClaims {
			statements = append(statements, strings.Split(subClaim, ",")...)
		}

		valid := true
		for _, statement := range statements {
			tokens := strings.Split(statement, " ")
			num, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic(err)
			}
			maxVal, ok := maxByColor[tokens[2]]
			if !ok {
				panic("invalid color " + tokens[2])
			}

			if num > maxVal {
				valid = false
				break
			}
		}

		if valid {
			sum = sum + id
		}
	}

	fmt.Printf("completed with %v", sum)
}
