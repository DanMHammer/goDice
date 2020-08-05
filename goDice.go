package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type Result struct {
	Input         string     `json:"input"`
	Valid         bool       `json:"valid"`
	Dice          []DieGroup `json:"dice"`
	RollTotal     int        `json:"roll_total"`
	Modifiers     []int      `json:"modifiers"`
	ModifierTotal int        `json:"modifier_total"`
	Total         int        `json:"total"`
	Image         string     `json:"image"`
}

type DieGroup struct {
	Size       int    `json:"size"`
	Count      int    `json:"count"`
	Operation  string `json: "operation"`
	High       int    `json:"high"`
	Low        int    `json:"low"`
	Multiplier int    `json:"multiplier"`
	Rolls      []int  `json:"rolls"`
	Kept       []int  `json:"kept"`
	Subtotal   int    `json:"subtotal"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/roll", roll)
	log.Fatal(http.ListenAndServe(":8080", router))
}

//4d20H3+3d4L1+12-3
func roll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	input := r.URL.Query().Get("string")
	out := parseInput(input)
	out.Valid = true
	out.Input = input
	json.NewEncoder(w).Encode(out)
}

func parseInput(input string) Result {
	commandReg := regexp.MustCompile(`[+-]*[\d\w]*`)
	countReg := regexp.MustCompile(`([+-]?)(\d*)`)
	sizeReg := regexp.MustCompile(`(?:[dD])(\d*)`)
	highReg := regexp.MustCompile(`(?:[hH])(\d*)`)
	lowReg := regexp.MustCompile(`(?:[lL])(\d*)`)
	modifierReg := regexp.MustCompile(`(?:[+-])(\d*)`)

	commands := commandReg.FindAllString(input, -1)
	fmt.Println(commands)
	result := Result{}
	result.Total = 0
	dice := []DieGroup{}
	for i := 0; i < len(commands); i++ {
		command := commands[i]
		group := DieGroup{}

		count := countReg.FindStringSubmatch(command)

		if count[1] == "" || count[1] == "+" {
			group.Operation = "+"
		} else {
			group.Operation = count[1]
		}

		countInt, _ := strconv.Atoi(count[2])

		size := sizeReg.FindStringSubmatch(command)

		if len(size) > 1 {
			sizeInt, _ := strconv.Atoi(size[1])
			group.Size = sizeInt

			if countInt == 0 && group.Size != 0 {
				group.Count = 1
			} else {
				group.Count = countInt
			}

			rolls := []int{}

			for j := 0; j < group.Count; j++ {
				rolls = append(rolls, rand.Intn(sizeInt-1)+1)
			}

			group.Rolls = rolls

			length := len(rolls)

			high := highReg.FindStringSubmatch(command)
			low := lowReg.FindStringSubmatch(command)

			if len(high) > 0 {
				highInt, _ := strconv.Atoi(high[1])
				group.High = highInt

				if highInt > 0 {
					sort.Ints(rolls)
					kept := rolls[length-highInt : length]
					group.Kept = kept
				}
			} else if len(low) > 0 {
				lowInt, _ := strconv.Atoi(low[1])
				group.Low = lowInt

				if lowInt > 0 {
					sort.Ints(rolls)
					kept := rolls[0:lowInt]
					group.Kept = kept
				}
			} else {
				group.Kept = group.Rolls
			}

			subtotal := 0

			for _, i := range group.Kept {
				subtotal += i
			}

			group.Subtotal = subtotal

			if group.Operation == "-" {
				group.Subtotal = 0 - group.Subtotal
			}

			result.Total = result.Total + group.Subtotal
			result.RollTotal = result.RollTotal + group.Subtotal

			dice = append(dice, group)

		} else {
			fmt.Println(modifierReg.FindStringSubmatch(command))
			modifier := modifierReg.FindStringSubmatch(command)
			modifierInt, _ := strconv.Atoi(modifier[1])
			if group.Operation == "-" {
				modifierInt = 0 - modifierInt
			}
			result.Modifiers = append(result.Modifiers, modifierInt)
			result.Total = result.Total + modifierInt
			result.ModifierTotal = result.ModifierTotal + modifierInt
		}

	}
	result.Dice = dice
	return result
}
