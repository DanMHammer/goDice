# goDice - An over-engineered go microservice for dice rolling

## Done

- Simple API for dice rolling based on string input.
- Unit testing
- Validate Die size (2, 4, 6, 8, 10, 12, 20, 100)
- Error handling

## In Progress

- Generate image of roll - Cache JSON in redis and generate SVG on demand when image is requested

## Todo

- Helm deployment for service
- Github Actions test -> build -> deploy
- Frontend on danhammer.dev

## Build and run

docker run -p 6379:6379 redis -d
go get
go build .\
go test -v
go run .\

## Usage

Dice code reference: https://rolz.org/

Accepts possible standard DnD dice 2, 4, 6, 8, 10, 12, and 20 sided as well as 100 sided.

H - highest

L - lowest

+/- add or subtract number from total

x multiply result

Use %2B for +

GET /roll/4d20H3%2B3d4L1%2B12-3

    4 d20: 4 20 sided dice
    H3: keep highest 3 of 4 d20
    3d4: 3 4 sided dice
    L1: keep lowest 1
    + 12: add 12
    - 3: subtract 3

Response:

```json
{
  "input": "4d20H3+3d4+12-3",
  "valid": true,
  "dice": [
    {
      "size": 20,
      "count": 4,
      "Operation": "+",
      "high": 3,
      "low": 0,
      "multiplier": 0,
      "rolls": [5, 6, 10, 15],
      "kept": [6, 10, 15],
      "subtotal": 31
    },
    {
      "size": 4,
      "count": 3,
      "Operation": "+",
      "high": 0,
      "low": 0,
      "multiplier": 0,
      "rolls": [2, 1, 2],
      "kept": [2, 1, 2],
      "subtotal": 5
    }
  ],
  "roll_total": 36,
  "modifiers": [12, -3],
  "modifier_total": 9,
  "total": 45,
  "image": ""
}
```

GET /roll/4d21H3

    4 d21: 4 21 sided dice - not possible

Response:

```json
{
  "input": "4d21H3",
  "valid": true,
  "dice": null,
  "roll_total": 0,
  "modifiers": null,
  "modifier_total": 0,
  "total": 0,
  "error": "Size is not valid. Size: 21 is not a standard dice with 2, 4, 6, 8, 10, 12, 20, or 100 sides"
}
```
