package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// https://golangdocs.com/structs-in-golang
type Result struct {
	opponent string
	player   string
}

var guaranteedPointsPerMove = map[string]uint16{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var pointsPerResult = map[Result]uint16{
	{opponent: "A", player: "X"}: 3,
	{opponent: "B", player: "X"}: 0,
	{opponent: "C", player: "X"}: 6,
	{opponent: "A", player: "Y"}: 6,
	{opponent: "B", player: "Y"}: 3,
	{opponent: "C", player: "Y"}: 0,
	{opponent: "A", player: "Z"}: 0,
	{opponent: "B", player: "Z"}: 6,
	{opponent: "C", player: "Z"}: 3,
}

func determineResultScore(res Result) uint16 {
	totalScore := uint16(0)

	totalScore += guaranteedPointsPerMove[res.player]
	totalScore += pointsPerResult[res]

	return uint16(totalScore)
}

// https://gosamples.dev/read-file/
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

		totalPoints += determineResultScore(currRes)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalPoints)
}
