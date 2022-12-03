package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Result struct {
	opponent string
	player   string
}

var guaranteedPointsPerMove = map[string]uint16{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func newStrat(res Result) Result {
	move := ""

	loose := map[string]string{"A": "Z", "B": "X", "C": "Y"}
	draw := map[string]string{"A": "X", "B": "Y", "C": "Z"}
	win := map[string]string{"A": "Y", "B": "Z", "C": "X"}

	if res.player == "X" {
		move = loose[res.opponent]
	} else if res.player == "Y" {
		move = draw[res.opponent]
	} else {
		move = win[res.opponent]
	}

	return Result{res.opponent, move}
}

func gameScore(res string) uint16 {
	points := uint16(0)

	if res == "Y" {
		points = 3
	} else if res == "Z" {
		points = 6
	}

	return points
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalPoints := uint16(0)
	var currRes Result

	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		currRes.opponent = moves[0]
		currRes.player = moves[1]

		newRes := newStrat(currRes)
		totalPoints += guaranteedPointsPerMove[newRes.player]
		totalPoints += gameScore(moves[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalPoints)
}
