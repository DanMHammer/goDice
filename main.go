package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
)

//Result ...
type Result struct {
	Input         string  `json:"input"`
	Valid         bool    `json:"valid"`
	Dice          []Group `json:"dice"`
	RollTotal     int     `json:"roll_total"`
	Modifiers     []int   `json:"modifiers"`
	ModifierTotal int     `json:"modifier_total"`
	Total         int     `json:"total"`
	Image         string  `json:"image"`
}

//Group ...
type Group struct {
	Size       int    `json:"size"`
	Count      int    `json:"count"`
	Operation  string `json:"operation"`
	High       int    `json:"high"`
	Low        int    `json:"low"`
	Multiplier int    `json:"multiplier"`
	Rolls      []int  `json:"rolls"`
	Kept       []int  `json:"kept"`
	Subtotal   int    `json:"subtotal"`
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/roll/{input}", roll)
	router.HandleFunc("/image/{id}", image)
	log.Fatal(http.ListenAndServe(":8080", router))
}

//4d20H3+3d4L1+12-3
func roll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	input := mux.Vars(r)["input"]
	result := performCommands(input)
	result.Valid = true
	result.Input = input

	id, _ := shortid.Generate()
	result.Image = "localhost:8080/image/" + id
	saveJSON(id, result)
	json.NewEncoder(w).Encode(result)
}

func image(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	resultJSON := getJSON(id)
	result := Result{}
	json.Unmarshal([]byte(resultJSON), &result)
	generate(w, result)
}
