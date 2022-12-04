package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Section struct {
	from int
	to   int
}

func getPair(line string) (Section, Section) {
	var firstElf, secondElf Section

	// https://github.com/tiri/aoc22/blob/main/04.go#L25
	fmt.Sscanf(line, "%d-%d,%d-%d", &firstElf.from, &firstElf.to, &secondElf.from, &secondElf.to)

	return firstElf, secondElf
}

func comparePair(firstElf Section, secondElf Section) bool {
	overlaps := false
	if firstElf.from <= secondElf.from && firstElf.to >= secondElf.to {
		overlaps = true
	}
	if secondElf.from <= firstElf.from && secondElf.to >= firstElf.to {
		overlaps = true
	}

	return overlaps
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalCompleteOverlaps := 0

	for scanner.Scan() {
		firstElf, secondElf := getPair(scanner.Text())

		if comparePair(firstElf, secondElf) {
			totalCompleteOverlaps += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalCompleteOverlaps)
}
