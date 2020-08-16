package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
)

var cacheEngineFlag = flag.String("engine", "gocache", "Storage engine to use for hashes and messages.  Supported: redis, gocache. Default: gocache")

// Cache - Cache Engine for saving results
var Cache CacheEngine

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
	Errors        string  `json:"error"`
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
	Unkept     []int  `json:"unkept"`
	Subtotal   int    `json:"subtotal"`
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	Cache, _ = SetupCache()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/roll/{input}", roll)
	router.HandleFunc("/image/{id}", image)
	router.HandleFunc("/rollImage/{input}", rollImage)
	log.Fatal(http.ListenAndServe(":3000", router))
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

	Cache.SaveResult(id, result)
	json.NewEncoder(w).Encode(result)
}

func image(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	result := Cache.GetResult(id)
	generate(w, result)
}

func rollImage(w http.ResponseWriter, r *http.Request) {
	input := mux.Vars(r)["input"]
	result := performCommands(input)
	result.Input = input
	generate(w, result)
}
