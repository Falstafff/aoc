package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	elfNotes, err := ReadFile("./2023/day1/input.txt.txt")

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

	fmt.Println(Sum(calories[0:3]))
}

func ReadFile(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(file)

	if err != nil {
		return lines, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
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

func FindMax(values []int) int {
	max := values[0]

	for _, value := range values {
		if value > max {
			max = value
		}
	}

	return max
}

func Sum(values []int) int {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum
}
