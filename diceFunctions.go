package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
)

var modifierReg = regexp.MustCompile(`(?:[+-])(\d*)`)
var commandReg = regexp.MustCompile(`[+-]*[\d\w]*`)
var countReg = regexp.MustCompile(`([+-]?)(\d*)`)
var sizeReg = regexp.MustCompile(`(?:[dD])(\d*)`)
var highReg = regexp.MustCompile(`(?:[hH])(\d*)`)
var lowReg = regexp.MustCompile(`(?:[lL])(\d*)`)

func performCommands(input string) Result {

	commands := parseCommands(input)
	fmt.Println(commands)
	result := Result{}
	result.Total = 0
	dice := []Group{}

	for i := 0; i < len(commands); i++ {
		command := commands[i]
		group, isDie := makeGroup(command)
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

func parseCommands(input string) []string {
	return commandReg.FindAllString(input, -1)
}

func makeGroup(command string) (Group, bool) {

	group := Group{}

	countInt, size, operation := parseCountSizeOp(command)

	if len(size) > 1 {
		sizeInt, _ := strconv.Atoi(size[1])
		group.Size = sizeInt

		group.Count = countInt
		group.Operation = operation

		rolls := rollDice(countInt, sizeInt, 0)
		group.Rolls = rolls

		highInt, lowInt, kept, sum := highLow(command, rolls)

		group.Kept = kept
		group.High = highInt
		group.Low = lowInt

		if operation == "-" {
			group.Subtotal = 0 - sum
		} else {
			group.Subtotal = sum
		}

		return group, true

	}

	modifier := modifierReg.FindStringSubmatch(command)
	modifierInt, _ := strconv.Atoi(modifier[1])
	if group.Operation == "-" {
		modifierInt = 0 - modifierInt
	}
	return Group{Subtotal: modifierInt}, false

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

func rollDice(countInt int, sizeInt int, override_seed int64) []int {

	if override_seed != 0 {
		rand.Seed(override_seed)
	}
	rolls := []int{}

	for j := 0; j < countInt; j++ {
		rolls = append(rolls, rand.Intn(sizeInt-1)+1)
	}

	return rolls
}

func highLow(command string, rolls []int) (int, int, []int, int) {
	high := highReg.FindStringSubmatch(command)
	low := lowReg.FindStringSubmatch(command)

	length := len(rolls)
	kept := []int{}
	highInt := 0
	lowInt := 0

	if len(high) > 0 {
		highInt, _ = strconv.Atoi(high[1])
		if highInt > 0 {
			sort.Ints(rolls)
			kept = append(kept, rolls[length-highInt:length]...)
		}
	}
	if len(low) > 0 {
		lowInt, _ = strconv.Atoi(low[1])

		if lowInt > 0 {
			sort.Ints(rolls)
			kept = append(kept, rolls[0:lowInt]...)
		}
	}
	if len(low) <= 0 && len(high) <= 0 {
		kept = rolls
	}

	sum := 0

	for _, i := range kept {
		sum += i
	}

	return highInt, lowInt, kept, sum
}
