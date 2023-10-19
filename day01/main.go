package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	backpacks := parseInput("input.txt")

	fmt.Println("Part 1: ", part1(&backpacks))
	fmt.Println("Part 2: ", part2(&backpacks))
}

func parseInput(filename string) [][]uint32 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	backpacks := make([][]uint32, 0)

	backpack := make([]uint32, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			backpacks = append(backpacks, backpack)
			backpack = make([]uint32, 0)
			continue
		}

		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		backpack = append(backpack, uint32(val))
	}
	backpacks = append(backpacks, backpack)

	return backpacks
}

func part1(backpacks *[][]uint32) uint32 {
	max := uint32(0)
	for _, backpack := range *backpacks {
		groupTotal := uint32(0)
		for _, item := range backpack {
			groupTotal += item
		}

		if groupTotal > max {
			max = groupTotal
		}
	}

	return max
}

func part2(backpacks *[][]uint32) uint32 {
	// max1 >= max2 >= max3
	max1 := uint32(0)
	max2 := uint32(0)
	max3 := uint32(0)

	for _, backpack := range *backpacks {
		backpackTotal := uint32(0)
		for _, item := range backpack {
			backpackTotal += item
		}

		if backpackTotal >= max1 {
			max3 = max2
			max2 = max1
			max1 = backpackTotal
		} else if backpackTotal >= max2 {
			max3 = max2
			max2 = backpackTotal
		} else if backpackTotal >= max3 {
			max3 = backpackTotal
		}
	}

	return max1 + max2 + max3
}
