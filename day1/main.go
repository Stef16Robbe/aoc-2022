package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	highest := uint64(0)

	for scanner.Scan() {
		num, err := strconv.ParseUint(scanner.Text(), 10, 0)

		if err == nil {
			elfTotal += num
		} else {
			// we have hit a newline
			if elfTotal > highest {
				highest = elfTotal
			}
			elfTotal = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(highest)
}
