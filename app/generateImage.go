package main

import (
	"fmt"
	"net/http"
	"strconv"

	svg "github.com/ajstarks/svgo"
)

func maxLength(dice []Group) int {
	max := 0

	for _, die := range dice {
		length := len(die.Rolls)
		if length > max {
			max = length
		}
	}
	return max
}

func generate(w http.ResponseWriter, result Result) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)

	canvasx := maxLength(result.Dice)*100 + 50
	canvasy := len(result.Dice)*100 + len(result.Modifiers)*100 + 200

	s.Start(canvasx, canvasy)
	x := 60
	y := 50

	s.Text(25, 30, result.Input, "font-size:25px;fill:black;font-weight:bold")

	for i := 0; i < len(result.Dice); i++ {
		x = 60
		size := result.Dice[i].Size

		text(3, y+35, s, fmt.Sprintf("d%d", size))
		x += 35
		text(55, y+35, s, result.Dice[i].Operation)

		s.Roundrect(x-10, y-10, 60*len(result.Dice[i].Kept)+5, 70, 10, 10, "fill:none;stroke:black;stroke-width:5")

		for j := 0; j < len(result.Dice[i].Kept); j++ {
			die(x, y, s, size, result.Dice[i].Kept[j])
			x += 60
		}

		for k := 0; k < len(result.Dice[i].Unkept); k++ {
			die(x, y, s, size, result.Dice[i].Unkept[k])
			x += 60
		}

		text(x+10, y+35, s, "=")
		text(x+35, y+35, s, strconv.Itoa(result.Dice[i].Subtotal))
		y += 80
	}

	if len(result.Modifiers) > 0 {
		text(3, y+35, s, "Modifiers:")
	}

	x = 150

	for i := 0; i < len(result.Modifiers); i++ {
		text(x, y+35, s, fmt.Sprintf("%d", result.Modifiers[i]))
		x += 60
	}

	text(3, y+80, s, fmt.Sprintf("Total: %d", result.Total))
	s.End()
}

func text(x int, y int, s *svg.SVG, text string) {
	switch text {
	case "+":
		s.Text(x, y, "+", "font-size:30px;fill:blue;font-weight:bold")
	case "-":
		s.Text(x+5, y, "-", "font-size:30px;fill:red;font-weight:bold")
	case "=":
		s.Text(x, y, "=", "font-size:30px;fill:black;font-weight:bold")
	default:
		s.Text(x, y, text, "font-size:30px;fill:black;font-weight:bold")
	}
}

func die(x int, y int, s *svg.SVG, size int, value int) {
	switch size {
	case 2:
		d2(s, value, x, y)
	case 4:
		d4(s, value, x, y)
	case 6:
		d6(s, value, x, y)
	case 8:
		d8(s, value, x, y)
	case 10:
		d10(s, value, x, y)
	case 12:
		d12(s, value, x, y)
	case 20:
		d20(s, value, x, y)
	case 100:
		s.Text(x+20, y+35, strconv.Itoa(value), "font-size:20pt;fill:black")
	}
}

func d2(s *svg.SVG, value, x int, y int) {
	s.Circle(x+25, y+25, 25, "fill:gray;stroke:black")
	s.Text(x+20, y+35, strconv.Itoa(value), "font-size:20pt;fill:white")
}

func d4(s *svg.SVG, value, x int, y int) {
	xcoords := []int{x, x + 25, x + 50}
	ycoords := []int{y + 50, y, y + 50}

	s.Polygon(xcoords, ycoords, "fill:purple;stroke:black")
	s.Text(x+20, y+40, strconv.Itoa(value), "font-size:20pt;fill:white")
}

func d6(s *svg.SVG, value, x int, y int) {
	s.Rect(x, y, 50, 50, "fill:blue;stroke:black")
	s.Text(x+20, y+35, strconv.Itoa(value), "font-size:20pt;fill:white")
}

func d8(s *svg.SVG, value, x int, y int) {
	xcoords := []int{x, x + 25, x + 50, x + 25}
	ycoords := []int{y + 25, y - 5, y + 25, y + 55}

	s.Polygon(xcoords, ycoords, "fill:#ff66cc;stroke:black")
	s.Text(x+20, y+35, strconv.Itoa(value), "font-size:20pt;fill:white")
}

func d10(s *svg.SVG, value, x int, y int) {
	x += 25
	y += 25
	xcoords := []int{x + 12, x - 13, x - 25, x - 12, x + 12, x + 25}
	ycoords := []int{y - 22, y - 22, y, y + 22, y + 22, y}

	s.Polygon(xcoords, ycoords, "fill:#ffa31a;stroke:black")
	s.Text(x-5, y+10, strconv.Itoa(value), "font-size:20pt;fill:white")
}

func d12(s *svg.SVG, value, x int, y int) {
	x += 25
	y += 25
	xcoords := []int{x + 8, x - 8, x - 20, x - 25, x - 20, x - 8, x + 8, x + 20, x + 25, x + 20}
	ycoords := []int{y - 24, y - 24, y - 15, y, y + 15, y + 24, y + 24, y + 15, y, y - 15}

	s.Polygon(xcoords, ycoords, "fill:green;stroke:black")
	s.Text(x-10, y+10, strconv.Itoa(value), "font-size:20pt;fill:white")
}

func d20(s *svg.SVG, value, x int, y int) {
	x += 25
	y += 25
	xcoords := []int{x + 12, x - 13, x - 25, x - 12, x + 12, x + 25}
	ycoords := []int{y - 22, y - 22, y, y + 22, y + 22, y}



	
	s.Polygon(xcoords, ycoords, "fill:black;stroke:black")
	s.Text(x-10, y+10, strconv.Itoa(value), "font-size:20pt;fill:white")
}
