package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var spelledDigits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getPossibleDigits(line string) (res []string) {
	// at each position, check if we have a numeric digit OR a spelled digit
	runes := []rune(line)
	for i := 0; i < len(runes); i++ {

		if unicode.IsDigit(runes[i]) {
			res = append(res, string(runes[i]))
			continue
		}

		for spelledDigit, numericDigit := range spelledDigits {
			if len(line)-i < len(spelledDigit) {
				continue
			}

			if line[i:i+len(spelledDigit)] == spelledDigit {
				res = append(res, numericDigit)
				break
			}
		}
	}
	return
}

func main() {

	f, err := os.Open("./input/1.2/input.txt")
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
		digits := getPossibleDigits(line)
		strNum := digits[0] + digits[len(digits)-1]
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to convert %s to integer: %w", strNum, err))
		}
		sum = sum + num
	}

	fmt.Printf("finished with sum %v", sum)
}
