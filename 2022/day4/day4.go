package main

import (
	"fmt"
	"github.com/Falstafff/aoc/utils"
	"strconv"
	"strings"
)

func main() {
	data, err := utils.ReadFile("./2022/day4/input.txt")

	utils.Check(err)

	elvesPairs := ParseElvesPairs(data)

	var overlapCount int

	for _, elvesPair := range elvesPairs {
		if utils.PointsOverlap(elvesPair[0], elvesPair[1], elvesPair[2], elvesPair[3]) {
			overlapCount++
		}
	}

	fmt.Println(overlapCount)

}

func ParseElvesPairs(pairs []string) [][]int {
	parsedPairs := make([][]int, len(pairs))

	for i, pair := range pairs {
		rawPair := strings.Split(pair, ",")
		rawPair = append(strings.Split(rawPair[0], "-"), strings.Split(rawPair[1], "-")...)

		parsedPair := make([]int, len(rawPair))

		for i, pairItem := range rawPair {
			parsedPair[i], _ = strconv.Atoi(pairItem)
		}

		parsedPairs[i] = parsedPair
	}

	return parsedPairs
}
