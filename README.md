# goDice - A go microservice for dice rolling

See it in Action: https://k8s.danhammer.dev/rollImage/-3d2H2%2B5d4H1%2B4d6H3%2B3d8%2B2d10H1%2B3d12-2d20

## Done

- Simple API for dice rolling based on string input.
- Unit testing
- Validate Die size (2, 4, 6, 8, 10, 12, 20, 100)
- Error handling
- Generate image of roll - Cache JSON in redis and generate SVG on demand when image is requested
- Helm deployment
- Github Actions build -> deploy

## In Progress

## Todo

- Github Actions test
- Automatic deployment to heroku
- Add redis to heroku

## Build and run

```
go build -o app
PORT=5000 ./app
```

Test it out! http://0.0.0.0:3000/rollImage/-3d2H2%2B5d4H1%2B4d6H3%2B3d8%2B2d10H1%2B3d12-2d20

Example roll:

![image](https://user-images.githubusercontent.com/35697323/135666980-4f081643-385a-4b20-af8f-6d30d1a05862.png)

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
