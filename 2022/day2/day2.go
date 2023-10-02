package main

import (
	"fmt"
	"github.com/Falstafff/aoc/utils"
	"strings"
)

func main() {
	data, err := utils.ReadFile("./2022/day2/input.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	rounds, err := ParseRounds(data, map[string]Gesture{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	})

	if err != nil {
		fmt.Println("Round parsing error", err)
		return
	}

	game := Game{
		outcome: &PlayerParrotOutcome{},
	}

	game.Start(rounds)

	fmt.Println(game.GetScore())
}

type Gesture int

const (
	Rock Gesture = iota + 1
	Paper
	Scissors
)

type Outcome int

const (
	Lost Outcome = 0
	Draw         = 3
	Won          = 6
)

const (
	MustLose = Rock
	MustDraw = Paper
	MustWin  = Scissors
)

type Rounds [][]Gesture

type OutcomeStrategy interface {
	GetOutcome(round *[]Gesture) Outcome
}

type ClassicOutcome struct{}

func (o *ClassicOutcome) GetOutcome(round *[]Gesture) Outcome {
	enemy, player := (*round)[0], (*round)[1]

	if enemy == player {
		return Draw
	}

	if (player == Rock && enemy == Scissors) || (player == Scissors && enemy == Paper) || (player == Paper && enemy == Rock) {
		return Won
	}

	return Lost
}

type PlayerParrotOutcome struct{}

func (o *PlayerParrotOutcome) GetOutcome(round *[]Gesture) Outcome {
	enemy, player := (*round)[0], (*round)[1]

	switch player {
	case MustLose:
		if enemy == Rock {
			player = Scissors
		} else if enemy == Paper {
			player = Rock
		} else {
			player = Paper
		}

		(*round)[1] = player

		return Lost
	case MustDraw:
		player = enemy

		(*round)[1] = player

		return Draw
	case MustWin:
		if enemy == Rock {
			player = Paper
		} else if enemy == Paper {
			player = Scissors
		} else {
			player = Rock
		}

		(*round)[1] = player

		return Won
	default:
		panic("Wrong player type")
	}

}

type Game struct {
	total   int
	outcome OutcomeStrategy
}

func (g *Game) Start(rounds Rounds) {
	for _, round := range rounds {
		outcome := g.outcome.GetOutcome(&round)
		g.total += g.CalcScore(round, outcome)
	}
}

func (g *Game) GetScore() int {
	return g.total
}

func (g *Game) CalcScore(round []Gesture, outcome Outcome) int {
	player := round[1]
	return int(player) + int(outcome)
}

func ParseRounds(items []string, guessToGesture map[string]Gesture) (Rounds, error) {
	rounds := make(Rounds, len(items))

	for index, item := range items {
		guesses := strings.Split(item, " ")

		gestures := make([]Gesture, len(guesses))

		for index, guess := range guesses {
			gesture, ok := guessToGesture[guess]

			if !ok {
				return nil, fmt.Errorf("unknown %s guess in the map", guess)
			}

			gestures[index] = gesture
		}

		rounds[index] = gestures
	}

	return rounds, nil
}
