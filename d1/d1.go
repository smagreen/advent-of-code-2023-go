package main

import (
	"aoc2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func concatenateFirstAndLastNumbers(input []string) (result []int) {

	re := regexp.MustCompile(`\d`)
	for _, line := range input {
		if digits := re.FindAllString(line, -1); len(digits) > 0 {
			this := digits[0]

			if len(digits) == 1 {
				this += digits[0]
			} else {
				this += digits[len(digits)-1]
			}

			value, _ := strconv.Atoi(this)
			result = append(result, value)
		}
	}

	return result
}

func Part1(lines []string) int {
	result := concatenateFirstAndLastNumbers(lines)
	sum := 0
	for _, num := range result {
		sum += num
	}
	return sum
}

func concatenateFirstAndLastDigitsNumber(lines []string) (result []int) {
	type Positions struct {
		firstIndex  int
		firstNumber int
		lastIndex   int
		lastNumber  int
	}

	nums := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	for _, line := range lines {
		pos := Positions{firstIndex: len(line), lastIndex: -1}
		for k, v := range nums {

			if firstDigit := strings.Index(line, strconv.Itoa(v)); firstDigit >= 0 && firstDigit < pos.firstIndex {
				pos.firstIndex = firstDigit
				pos.firstNumber = v
			}

			if firstWord := strings.Index(line, k); firstWord >= 0 && firstWord < pos.firstIndex {
				pos.firstIndex = firstWord
				pos.firstNumber = v
			}

			if lastDigit := strings.LastIndex(line, k); lastDigit >= 0 && lastDigit > pos.lastIndex {
				pos.lastIndex = lastDigit
				pos.lastNumber = v
			}

			if lastWord := strings.LastIndex(line, strconv.Itoa(v)); lastWord >= 0 && lastWord > pos.lastIndex {
				pos.lastIndex = lastWord
				pos.lastNumber = v
			}
		}
		digits, _ := strconv.Atoi(fmt.Sprintf("%d%d", pos.firstNumber, pos.lastNumber))
		result = append(result, digits)
	}
	return result
}

func Part2(lines []string) int {
	result := concatenateFirstAndLastDigitsNumber(lines)
	sum := 0
	for _, num := range result {
		sum += num
	}
	return sum
}

func main() {
	input := utils.GetAoCInputHTTP(1)
	println(Part1(input))
	println(Part2(input))
}
