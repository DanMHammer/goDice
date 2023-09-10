package dice

/*

{
	"dice": [
		{
			"size": 20,
			"count": 4,
			"highest": 3,
			"lowest": 1,
			"rolls": [1, 2, 3, 4],
			"kept": [2, 3, 4],
			"unkept": [1],
			"subtotal": 9
		}
	],
	"roll_total": 9,
	"error": "",
}
*/

type DieRequest struct {
	Size    int `json:"size"`
	Count   int `json:"count"`
	Highest int `json:"highest"`
	// Lowest  int `json:"lowest"`
}

type DieResponse struct {
	DieRequest
	Rolls       []int `json:"rolls"`
	HighestKept []int `json:"highest-kept"`
	// LowestKept  []int `json:"lowest-kept"`
	Unkept   []int `json:"unkept"`
	Subtotal int   `json:"subtotal"`
}

type RollRequest struct {
	Dice []DieRequest `json:"dice"`
}

type RollResponse struct {
	Dice     []DieResponse `json:"dice"`
	Total    int           `json:"total"`
	Error    string        `json:"error"`
	ImageUrl string        `json:"image"`
}
