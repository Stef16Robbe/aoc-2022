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

// https://github.com/tiri/aoc22/blob/main/04.go#L14
func (p Section) len() int {
	return p.to - p.from
}

func getPair(line string) (Section, Section) {
	var firstElf, secondElf Section

	// https://github.com/tiri/aoc22/blob/main/04.go#L25
	fmt.Sscanf(line, "%d-%d,%d-%d", &firstElf.from, &firstElf.to, &secondElf.from, &secondElf.to)

	if firstElf.len() < secondElf.len() {
		tmp := firstElf
		firstElf = secondElf
		secondElf = tmp
	}

	return firstElf, secondElf
}

func comparePair(firstElf Section, secondElf Section) bool {
	overlaps := true

	if firstElf.to < secondElf.from || secondElf.to < firstElf.from {
		overlaps = false
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
