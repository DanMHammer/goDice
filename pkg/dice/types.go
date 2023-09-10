package dice

// Result ...
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

// Group ...
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
