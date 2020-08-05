package main

import (
	"net/http"

	svg "github.com/ajstarks/svgo"
)

func generate(w http.ResponseWriter, result Result) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	x := 0
	y := 0
	for i := 0; i < len(result.Dice); i++ {
		operation(s, result.Dice[i], x, y)
		x += 50
		die(s, result.Dice[i], x, y)
		x += 101
	}
	s.End()
}

func operation(s *svg.SVG, die Group, x int, y int) {
	s.Text(x, y, "+", "font-size:30px;fill:blue")
}

func die(s *svg.SVG, die Group, x int, y int) {
	switch die.Size {
	case 6:
		d6(s, die, x, y)
	}
}

func d6(s *svg.SVG, die Group, x int, y int) {
	s.Rect(x, y, 100, 100, "fill:none;stroke:black")
}
