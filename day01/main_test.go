package main

import (
	"testing"
)

var backpacks = parseInput("test_input.txt")

func TestPart1(t *testing.T) {
	// Call the function
	result := part1(&backpacks)

	// Check the result
	if result != 24_000 {
		t.Errorf("Expected 24000, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	// Call the function
	result := part2(&backpacks)

	// Check the result
	if result != 45_000 {
		t.Errorf("Expected 45000, got %d", result)
	}
}
