package main

import (
	"math/rand"
	"testing"
)

var input = "4d20H3+3d4L1+12-3"

func EqualStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestParseCommands(t *testing.T) {
	commands, _ := parseCommands(input)
	expectedCommands := []string{"4d20H3", "+3d4L1", "+12", "-3"}
	if !EqualStringSlices(commands, expectedCommands) {
		t.Errorf("Commands parsed incorrectly. Got: %s. Expected %s", commands, expectedCommands)
	}
}

func TestParseCountSizeOP(t *testing.T) {
	count, size, operation := parseCountSizeOp("4d20H3")

	if count != 4 {
		t.Errorf("Count parsed incorrectly. Got: %d. Expected %d.", count, 4)
	}
	if size[1] != "20" {
		t.Errorf("Size parsed incorrectly. Got: %s. Expected %d.", size[1], 20)
	}
	if operation != "+" {
		t.Errorf("Operation parsed incorrectly. Got: %s. Expected %s.", operation, "+")
	}
}

func TestRollDice(t *testing.T) {
	for i := 0; i < 10; i++ {
		count := rand.Intn(9) + 1
		size := []int{2, 4, 6, 8, 10, 12, 20, 100}[rand.Intn(8)]
		rolls, err := rollDice(count, size, 0)

		if err != nil {
			t.Errorf(err.Error())
		}

		if len(rolls) != count {
			t.Errorf("Number of rolls is incorect got: %d, needed: %d", len(rolls), count)
		}

		for i := 0; i < len(rolls); i++ {
			if rolls[i] <= 0 || rolls[i] > size {
				t.Errorf("Roll value out of bounds. Got: %d. Expected a value between 1 and %d.", rolls[i], size)
			}
		}

	}
}

type RollSpecificDiceData struct {
	count     int
	size      int
	rolls     []int
    seed      int64
	shouldErr bool
}

func TestRollSpecificDice(t *testing.T) {
	items := []RollSpecificDiceData{
		{4, 20, []int{2, 8, 8, 20}, 1, false},
		{4,  2, []int{1, 1, 1, 1}, 2, false},
		{3, 10, []int{9, 8, 7}, 3, false},
		{1, 2, []int{2}, 4, false},
		{3, 11, []int{}, 5, true},
		{2, 13, []int{}, 6, true},
		{5, 12, []int{}, 0, false},
	}

	for _, item := range items {
		rolls, err := rollDice(item.count, item.size, item.seed)

		if item.seed == 0 {
			if len(rolls) != item.count{
				t.Errorf("Incorrect number of rolls. Got: %d. Expected %d.", len(rolls), item.count)
			}
			for _, roll := range rolls {
				if roll > item.size || roll <= 0 {
					t.Errorf("Roll out of bounds. Got: %d. Expected 0 < x < %d.", roll, item.size)
				}
			}
		} else if !item.shouldErr && err != nil {
			t.Errorf(err.Error())
		} else if !item.shouldErr && err == nil {

			if !EqualIntSlices(rolls, item.rolls) {
				t.Errorf("Rolls incorrect. Got: %d. Expected %d.", rolls, item.rolls)
			}
		}
	}

}

type HighLowData struct {
	command   string
	rolls     []int
	highInt   int
	lowInt    int
	sum       int
	kept      []int
	unkept    []int
	shouldErr bool
}

func TestHighLow(t *testing.T) {
	items := []HighLowData{
		{"4d10H2", []int{1, 2, 3, 4}, 2, 0, 7, []int{3, 4}, []int{1, 2}, false},
		// {"4d10H2L1", []int{6, 2, 10, 4}, 2, 1, 18, []int{6, 10, 2}, []int{4}, false},
		// {"6d10H1L2", []int{1, 2, 3, 4}, 1, 2, 7, []int{4, 1, 2}, []int{3}, false},
		{"3d8", []int{8, 7, 3, 5}, 0, 0, 23, []int{8, 7, 3, 5}, []int{}, false},
		// {"2d10HL3", []int{1, 2}, 0, 0, 0, []int{}, []int{}, true},
		// {"2d10H2L2", []int{1, 2}, 0, 0, 0, []int{}, []int{}, true},
	}

	for _, item := range items {
		highInt, lowInt, kept, sum, unkept, err := highLow(item.command, item.rolls)

		if !item.shouldErr && err != nil {
			t.Errorf(err.Error())
		} else if !item.shouldErr && err == nil {
			if highInt != item.highInt {
				t.Errorf("highInt incorrect. Got: %d. Expected %d.", highInt, item.highInt)
			}

			if lowInt != item.lowInt {
				t.Errorf("lowInt incorrect. Got: %d. Expected %d.", lowInt, item.lowInt)
			}

			if sum != item.sum {
				t.Errorf("sum incorrect. Got: %d. Expected %d.", sum, item.sum)
			}

			if !EqualIntSlices(kept, item.kept) {
				t.Errorf("kept incorrect. Got: %d. Expected %d.", kept, item.kept)
			}

			if !EqualIntSlices(unkept, item.unkept) {
				t.Errorf("unkept incorrect. Got: %d. Expected %d.", unkept, item.unkept)
			}
		}
	}
}
