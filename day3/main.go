package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func contains(elems []rune, v rune) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}

	return false
}

func getSacks(line string) ([]rune, []rune) {
	itemCount := len(line)

	// sackOne := strings.Split(line[:itemCount/2], "")
	sackOne := []rune(line[:itemCount/2])
	// sackTwo := strings.Split(line[itemCount/2:], "")
	sackTwo := []rune(line[itemCount/2:])

	return sackOne, sackTwo
}

func getCommonItems(sackOne []rune, sackTwo []rune) []rune {
	commonItems := []rune{}

	for _, sOneItem := range sackOne {
		for _, sTwoitem := range sackTwo {
			if sOneItem == sTwoitem && !contains(commonItems, sOneItem) {
				commonItems = append(commonItems, sOneItem)
			}
		}
	}

	return commonItems
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

// 1. read line content into two different arrays (sack 1, sack 2) [X]
// 2. find the common item between the two, add it to a seperate slice [X]
// 3. find a smart way to get the priorities without hard coding [X]
// 4. find priorities for all items in slice, add them up together [X]
// 5. ???
// 6. profit [X]
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	commonItems := []rune{}
	for scanner.Scan() {
		sackOne, sackTwo := getSacks(scanner.Text())
		commonItems = append(commonItems, getCommonItems(sackOne, sackTwo)...)
	}

	prioritiesTotal := getPrioritiesTotal(commonItems)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(prioritiesTotal)
}
