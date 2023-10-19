package main

import (
	"testing"
)

var matches [][]uint8 = parseInput("test_input.txt")

func TestPart1(t *testing.T) {
	score := part1(&matches)

	if score != 15 {
		t.Errorf("Expected 15, got %d", score)
	}
}

func TestPart2(t *testing.T) {
	score := part2(&matches)

	if score != 12 {
		t.Errorf("Expected 15, got %d", score)
	}
}