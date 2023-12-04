package main

import (
	"AoC/pkg/input"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	f, err := input.GetInputFile("2.2", "input.txt")
	if err != nil {
		panic(err)
	}
	defer input.DeferFunc(f)

	var sum int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		claims := strings.Split(line, ":")[1]
		subClaims := strings.Split(claims, ";")
		var statements []string
		for _, subClaim := range subClaims {
			statements = append(statements, strings.Split(subClaim, ",")...)
		}

		minMap := map[string]int{}
		for _, statement := range statements {
			tokens := strings.Split(statement, " ")
			color := tokens[2]
			num, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic("invalid number in" + tokens[1])
			}

			if minMap[color] < num {
				minMap[color] = num
			}
		}

		power := 1
		for _, v := range minMap {
			power = power * v
		}

		sum = sum + power
	}

	fmt.Printf("completed with %v", sum)
}
