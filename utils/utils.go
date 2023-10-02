package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(name string) ([]string, error) {
	var lines []string

	file, err := os.Open(name)

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
