package main

import (
	"aoc2023/utils"
	"testing"
)

func TestDay1(t *testing.T) {
	t.Run("Test Part 1 Example", func(t *testing.T) {
		input := utils.GetAoCInputFile("part1-example.txt")
		got := Part1(input)

		if got != 142 {
			t.Errorf("expected 142 got %d ", got)
		}
	})

	t.Run("Test Part 1 Solve", func(t *testing.T) {
		input := utils.GetAoCInputFile("input.txt")
		got := Part1(input)

		if got != 55712 {
			t.Errorf("expected 55712 got %d ", got)
		}
	})

	t.Run("Test Part 2 Example", func(t *testing.T) {
		input := utils.GetAoCInputFile("part2-example.txt")
		got := Part2(input)

		if got != 281 {
			t.Errorf("expected 281 got %d ", got)
		}
	})
}
