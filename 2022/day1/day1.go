package main

import (
	"fmt"
	"github.com/Falstafff/aoc/utils"
	"sort"
	"strconv"
)

func main() {
	elfNotes, err := utils.ReadFile("./2022/day1/input.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	calories, err := ReduceCalories(elfNotes)

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	fmt.Println(utils.Sum(calories[0:3]))
}

func ReduceCalories(lines []string) ([]int, error) {
	var calories []int
	index := 0

	for _, line := range lines {
		if len(calories) != index+1 {
			calories = append(calories, 0)
		}

		if line == "" {
			index++
			continue
		}

		calorie, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println("Error converting to int", err)
			return calories, err
		}

		calories[index] += calorie
	}

	return calories, nil
}
