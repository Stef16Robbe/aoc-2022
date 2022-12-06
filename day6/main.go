package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func removeDuplicateValues(runeSlice []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range runeSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func checkMarker(m []rune) bool {
	isMarker := true
	if len(removeDuplicateValues(m)) != 14 {
		isMarker = false
	}

	return isMarker
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	processedChars := 0
	for scanner.Scan() {
		line := scanner.Text()
		chars := []rune(line)
		from := 0
		to := 14

		for i := 0; i < len(line)-1; i++ {
			marker := chars[from:to]

			if checkMarker(marker) {
				processedChars = to
				break
			}
			from++
			to++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(processedChars)
}
