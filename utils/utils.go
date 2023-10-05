package utils

import (
	"bufio"
	"fmt"
	"log"
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

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateAlphabet() map[string]int {
	alphabet := make(map[string]int)
	index := 1

	for i := 'a'; i <= 'z'; i++ {
		alphabet[string(i)] = index
		index++
	}

	for i := 'A'; i <= 'Z'; i++ {
		alphabet[string(i)] = index
		index++
	}

	return alphabet
}

func BatchSlice[T any](in []T, size int) [][]T {
	out := make([][]T, 0)

	for i := 0; i < len(in); i += size {
		j := i + size

		if j >= len(in) {
			j = len(in)
		}

		out = append(out, in[i:j])
	}

	return out
}

func Min(x1, x2 int) int {
	if x1 <= x2 {
		return x1
	}
	return x2
}

func Max(x1, x2 int) int {
	if x1 >= x2 {
		return x1
	}
	return x2
}

func PointsOverlap(x1, x2, y1, y2 int) bool {
	start, end := Max(x1, y1), Min(x2, y2)
	return start <= end
}
