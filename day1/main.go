package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// https://gosamples.dev/read-file/
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	elfTotal := uint64(0)
	totals := []uint64{}

	for scanner.Scan() {
		num, err := strconv.ParseUint(scanner.Text(), 10, 0)

		if err == nil {
			elfTotal += num
		} else {
			// we have hit a newline
			totals = append(totals, elfTotal)
			elfTotal = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// https://stackoverflow.com/a/40932847/10503012
	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})

	fmt.Println(totals[0] + totals[1] + totals[2])
}
