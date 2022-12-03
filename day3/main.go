package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getBadge(group []string) rune {
	var badge rune

	for _, itemElfOne := range group[0] {
		for _, itemElfTwo := range group[1] {
			for _, itemElfThree := range group[2] {
				if itemElfOne == itemElfTwo && itemElfOne == itemElfThree {
					badge = itemElfOne
				}
			}
		}
	}

	return badge
}

func getPrioritiesTotal(commonItems []rune) uint16 {
	prioritiesTotal := uint16(0)

	for _, ci := range commonItems {
		if ci >= 'a' && ci <= 'z' {
			prioritiesTotal += uint16(ci) - 96
		} else if ci >= 'A' && ci <= 'Z' {
			prioritiesTotal += uint16(ci) - 38
		}
	}

	return prioritiesTotal
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	badges := []rune{}
	group := []string{}

	for scanner.Scan() {
		if len(group) == 3 {
			badges = append(badges, getBadge(group))
			group = []string{}
		}

		group = append(group, scanner.Text())
	}

	prioritiesTotal := getPrioritiesTotal(badges)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(prioritiesTotal)
}
