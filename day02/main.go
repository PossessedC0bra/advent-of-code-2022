package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var valueMap = map[string]uint8{
	"A": 1, // Rock
	"B": 2, // Paper
	"C": 3, // Scissors
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissors
}

func main() {
	matches := parseInput("input.txt")

	fmt.Println("Part 1: ", part1(&matches))
	fmt.Println("Part 2: ", part2(&matches))
}

func parseInput(filename string) [][]uint8 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	matches := make([][]uint8, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		match := make([]uint8, 0)
		for _, token := range tokens {
			match = append(match, valueMap[token])
		}
		matches = append(matches, match)
	}

	return matches
}

// Part 1: given a list of matches calculate the score of player 2.
// The score is calculated as follows:
// - losses are worth 0 points
// - draws are worth 3 points
// - wins are worth 6 points
// Additionally the used piece grants additional points:
// - Rock: 1 point
// - Paper: 2 points
// - Scissors: 3 points
func part1(matches *[][]uint8) uint32 {
	score := uint32(0)
	for _, match := range *matches {
		var res int = int(match[0]) - int(match[1])

		score += uint32(match[1])

		if res == 0 { // draw
			score += 3
		} else if res == -1 || res == 2 { // win
			score += 6
		}
	}

	return score
}

// Part 2: given a list of used pieces by player 1 and the expected outcome of
// the match calculate the score of player 2.
// The expected outcome is defined as follows:
// - X: player 1 wins
// - Y: draw
// - Z: player 2 wins
// The score is calculated as follows:
// - losses are worth 0 points
// - draws are worth 3 points
// - wins are worth 6 points
// Additionally the used piece grants additional points:
// - Rock: 1 point
// - Paper: 2 points
// - Scissors: 3 points
func part2(matches *[][]uint8) uint32 {
	score := uint32(0)
	for _, match := range *matches {
		// expect a draw
		if match[1] == 2 {
			// 3 points for a draw + points for the used piece (same piece as player 1)
			score += 3 + uint32(match[0])
		} else if match[1] == 3 { // expect a win
			score += 6
			// in order to win choose piece with value 1 higher than player 1
			nextPiece := match[0] + 1
			if nextPiece == 4 {
				score += 1
			} else {
				score += uint32(nextPiece)
			}
		} else { // expect a loss
			// in order to lose choose piece with value 1 lower than player 1
			previousPiece := match[0] - 1
			if previousPiece == 0 {
				score += 3
			} else {
				score += uint32(previousPiece)
			}
		}
	}

	return score
}
