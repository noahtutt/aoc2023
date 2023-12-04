package main

import (
	string2 "AoC/pkg/string"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func getFirstNumber(line string) (string, error) {

	for _, c := range line {
		if unicode.IsDigit(c) {
			return string(c), nil
		}
	}

	return "", fmt.Errorf("no digits were found")
}

func main() {

	f, err := os.Open("./input/1.1/input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to open file: %w", err))
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("failed to close file: %w", err))
		}
	}(f)

	var sum int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		digit1, err := getFirstNumber(line)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to get first digit for line %s: %w", line, err))
		}

		digit2, err := getFirstNumber(string2.Reverse(line))
		if err != nil {
			fmt.Println(fmt.Errorf("failed to get second digit for line %s: %w", line, err))
		}

		res, err := strconv.Atoi(digit1 + digit2)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to covert %s to integer", digit1+digit2))
		}
		sum = sum + res
	}

	fmt.Printf("finished with sum %v\n", sum)
}
