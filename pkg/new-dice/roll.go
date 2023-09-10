package newdice

import (
	"math/rand"
	"sort"
)

func RollDice(req RollRequest) (RollResponse, error) {
	res := RollResponse{}

	for _, die := range req.Dice {
		outcome := die.RollDie()
		res.Dice = append(res.Dice, outcome)
	}

	return res, nil
}

func (die DieRequest) RollDie() DieResponse {
	res := DieResponse{
		DieRequest: die,
	}

	res.Rolls = []int{}

	for i := 0; i < die.Count; i++ {
		res.Rolls = append(res.Rolls, rand.Intn(die.Size)+1)
	}

	sorted := sort.IntSlice(res.Rolls)
	res.HighestKept = sorted[die.Count-die.Highest : die.Count]
	res.LowestKept = sorted[0:die.Lowest]
	res.Unkept = sorted[die.Lowest : die.Count-die.Highest]

	res.Subtotal = 0
	for _, roll := range res.HighestKept {
		res.Subtotal += roll
	}
	for _, roll := range res.LowestKept {
		res.Subtotal += roll
	}

	return res
}
