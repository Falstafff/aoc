package main

import (
	"fmt"
	"github.com/Falstafff/aoc/utils"
	"log"
)

func main() {
	data, err := utils.ReadFile("./2022/day3/input.txt")

	utils.Check(err)

	rucksacks := NewRucksackCollection(data)

	fmt.Println(rucksacks.FindBatchedBadgesSum(3))
}

type Rucksack struct {
	items string
}

func (r *Rucksack) FindCommonItem() (string, error) {
	uniqueItems := make(map[string]bool)
	length := len(r.items)
	mid := length / 2

	for i := 0; i < mid; i++ {
		item := string(r.items[i])

		if _, ok := uniqueItems[item]; !ok {
			uniqueItems[item] = true
		}
	}

	for i := mid; i < length; i++ {
		item := string(r.items[i])

		if _, ok := uniqueItems[item]; ok {
			return item, nil
		}

	}

	return "", fmt.Errorf("not found")
}

func NewRucksack(items string) Rucksack {
	return Rucksack{
		items: items,
	}
}

type RucksackCollection []Rucksack

func (rc *RucksackCollection) FindMostCommonItemSum() int {
	var sum int

	priorityByItem := utils.GenerateAlphabet()

	for _, rucksack := range *rc {
		item, err := rucksack.FindCommonItem()

		if err != nil {
			continue
		}

		priority, ok := priorityByItem[item]

		if !ok {
			log.Fatalf("%s item  priority not found", item)
		}

		sum += priority
	}

	return sum
}

func (rc *RucksackCollection) FindBatchedBadgesSum(batchSize int) int {
	var sum int

	priorityByItem := utils.GenerateAlphabet()

	for i := 0; i < len(*rc); i += batchSize {
		j := i + batchSize

		if j >= len(*rc) {
			j = len(*rc)
		}

		rucksacks := (*rc)[i:j]

		itemBadge, err := rc.FindCommonBadge(rucksacks)

		utils.Check(err)

		sum += priorityByItem[itemBadge]
	}

	return sum
}

func (rc *RucksackCollection) FindCommonBadge(rucksacks []Rucksack) (string, error) {
	itemsCount := make(map[rune]int)

	for _, rucksack := range rucksacks {
		uniqueItems := make(map[rune]bool)

		for _, item := range rucksack.items {
			if _, ok := uniqueItems[item]; !ok {
				uniqueItems[item] = true
				itemsCount[item]++
			}
		}
	}

	for item, count := range itemsCount {
		if count == 3 {
			return string(item), nil
		}
	}

	return "", fmt.Errorf("badge is not found")
}

func NewRucksackCollection(rucksacks []string) RucksackCollection {
	newRucksacks := make(RucksackCollection, len(rucksacks))

	for index, rucksack := range rucksacks {
		newRucksacks[index] = NewRucksack(rucksack)
	}

	return newRucksacks
}
