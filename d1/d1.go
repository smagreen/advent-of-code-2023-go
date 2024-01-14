package main

import (
	"aoc2023/utils"
	"regexp"
	"strconv"
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

func Part2(in []string) int {
	return 0
}

func main() {
	input := utils.GetAoCInputHTTP(1)
	println(Part1(input))
	println(Part2(input))
}
