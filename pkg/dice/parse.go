package dice

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var groupReg = regexp.MustCompile(`^\d*[d]\d*(?:[H]\d*){0,1}$`)
var countReg = regexp.MustCompile(`([+-]?)(\d*)`)
var sizeReg = regexp.MustCompile(`(?:[dD])(\d*)`)
var highReg = regexp.MustCompile(`(?:[hH])(\d*)`)

// 4d20H3+3d4L1
func ParseStringInput(input string) (RollRequest, error) {
	req := RollRequest{}

	groups := strings.Split(input, "+")

	for _, group := range groups {
		die := DieRequest{}

		// check if group is valid
		if !groupReg.MatchString(group) {
			return req, errors.New("invalid input")
		}
		var err error

		// get the die count
		count := countReg.FindString(group)
		die.Count, err = strconv.Atoi(count)
		if err != nil {
			return req, err
		}

		// get the die size
		size := sizeReg.FindStringSubmatch(group)[1]
		die.Size, err = strconv.Atoi(size)
		if err != nil {
			return req, err
		}

		// get the number of highest kept
		highest := highReg.FindStringSubmatch(group)[1]
		if highest == "" {
			die.Highest = die.Count
		} else {
			die.Highest, err = strconv.Atoi(highest)
			if err != nil {
				return req, err
			}
		}

		req.Dice = append(req.Dice, die)
	}

	return req, nil
}
