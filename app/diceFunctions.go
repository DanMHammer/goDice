package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var modifierReg = regexp.MustCompile(`(?:[+-])(\d*)`)
var commandReg = regexp.MustCompile(`[+-]*[\d\w]*`)
var countReg = regexp.MustCompile(`([+-]?)(\d*)`)
var sizeReg = regexp.MustCompile(`(?:[dD])(\d*)`)
var highReg = regexp.MustCompile(`(?:[hH])(\d*)`)
var lowReg = regexp.MustCompile(`(?:[lL])(\d*)`)
var allowedSizes = []int{2, 4, 6, 8, 10, 12, 20, 100}

func performCommands(input string) Result {

	commands, err := parseCommands(input)
	if err != nil {
		return Result{Valid: false, Errors: err.Error()}
	}
	result := Result{}
	result.Total = 0
	dice := []Group{}

	for i := 0; i < len(commands); i++ {
		command := commands[i]
		group, isDie, err := makeGroup(command)
		if err != nil {
			result.Errors = err.Error()
			result.Valid = false
			return result
		}

		if isDie {
			dice = append(dice, group)
			result.RollTotal = result.RollTotal + group.Subtotal
			result.Total = result.Total + group.Subtotal
		} else {
			result.Modifiers = append(result.Modifiers, group.Subtotal)
			result.ModifierTotal = result.ModifierTotal + group.Subtotal
			result.Total = result.Total + group.Subtotal
		}
	}
	result.Dice = dice
	return result
}

func parseCommands(input string) ([]string, error) {
	commands := commandReg.FindAllString(input, -1)
	commandsJoined := strings.Join(commands, "")

	if commandsJoined != input {
		return commands, fmt.Errorf("Commands did not parse correctly. Input: %s. Parsed Commands: %s", input, commands)
	}
	return commands, nil
}

func contains(arr []int, integer int) bool {
	for _, a := range arr {
		if a == integer {
			return true
		}
	}
	return false
}

func makeGroup(command string) (Group, bool, error) {

	group := Group{}

	countInt, size, operation := parseCountSizeOp(command)

	if len(size) > 1 {
		sizeInt, _ := strconv.Atoi(size[1])

		if !contains(allowedSizes, sizeInt) {
			return group, true, fmt.Errorf("Size is not valid. Size: %d is not a standard dice with 2, 4, 6, 8, 10, 12, 20, or 100 sides", sizeInt)
		}
		group.Size = sizeInt

		group.Count = countInt
		group.Operation = operation

		rolls, err := rollDice(countInt, sizeInt, 0)

		if err != nil {
			return group, true, err
		}

		group.Rolls = rolls

		highInt, lowInt, kept, sum, unkept, err := highLow(command, rolls)

		if err != nil {
			return group, true, err
		}

		group.Kept = kept
		group.High = highInt
		group.Low = lowInt
		group.Unkept = unkept

		if operation == "-" {
			group.Subtotal = 0 - sum
		} else {
			group.Subtotal = sum
		}

		return group, true, nil

	}

	modifier := modifierReg.FindStringSubmatch(command)
	modifierInt, _ := strconv.Atoi(modifier[0])

	return Group{Subtotal: modifierInt}, false, nil

}

func parseCountSizeOp(command string) (int, []string, string) {
	count := countReg.FindStringSubmatch(command)
	countInt, _ := strconv.Atoi(count[2])

	size := sizeReg.FindStringSubmatch(command)

	operation := ""
	if count[1] == "" || count[1] == "+" {
		operation = "+"
	} else {
		operation = count[1]
	}

	if countInt == 0 {
		countInt = 1
	}

	return countInt, size, operation
}

func rollDice(countInt int, sizeInt int, overrideSeed int64) ([]int, error) {

	if overrideSeed != 0 {
		rand.Seed(overrideSeed)
	}
	rolls := []int{}

	if !contains(allowedSizes, sizeInt) {
		return rolls, fmt.Errorf("Size is not valid. Size: %d is not a standard dice with 2, 4, 6, 8, 10, 12, 20, or 100 sides", sizeInt)
	}

	for j := 0; j < countInt; j++ {
		rolls = append(rolls, rand.Intn(sizeInt-1)+1)
	}

	return rolls, nil
}

func highLow(command string, rolls []int) (int, int, []int, int, []int, error) {
	high := highReg.FindStringSubmatch(command)
	low := lowReg.FindStringSubmatch(command)

	length := len(rolls)
	kept := []int{}
	unkept := []int{}
	highInt := 0
	lowInt := 0
	sum := 0

	if len(high) > 0 {
		highInt, _ = strconv.Atoi(high[1])
		if highInt > 0 {
			if highInt > length {
				return highInt, 0, kept, sum, unkept, fmt.Errorf("Too many high dice requested. %d dice rolled. %d high dice requested", length, highInt)
			}
			sort.Ints(rolls)
			kept = append(kept, rolls[length-highInt:length]...)
			unkept = rolls[0 : length-highInt]
		}
	}
	if len(low) > 0 {
		lowInt, _ = strconv.Atoi(low[1])

		if lowInt > 0 {
			if lowInt > length {
				return highInt, lowInt, kept, sum, unkept, fmt.Errorf("Too many low dice requested. %d dice rolled. %d low dice requested", length, lowInt)
			} else if lowInt+highInt > length {
				return highInt, lowInt, kept, sum, unkept, fmt.Errorf("Too many dice requested. %d dice rolled. %d low dice and %d high dice requested", length, lowInt, highInt)
			}
			sort.Ints(unkept)
			kept = append(kept, unkept[0:lowInt-1]...)
			unkept = kept

			for i, keep := range kept {
				for _, roll := range rolls {
					if roll == keep {
						unkept[i] = -1
					}
				}
			}

			unkeptFinal := unkept
			for _, unkeptRoll := range unkept {
				if unkeptRoll != -1 {
					unkeptFinal = append(unkeptFinal, unkeptRoll)
				}
			}

			unkept = unkeptFinal
		}
	}
	if len(low) <= 0 && len(high) <= 0 {
		kept = rolls
	}

	for _, i := range kept {
		sum += i
	}

	return highInt, lowInt, kept, sum, unkept, nil
}
