# goDice - An over-engineered and verbose go microservice for dice rolling

DONE:

- Simple API for dice rolling based on string input.

TODO:

- (Maybe) Generate image of roll - Cache JSON in redis and generate on demand when image is requested
- Unit testing for API
- Helm deployment for service
- Github Actions test -> build -> deploy
- Frontend on danhammer.dev

API:

Dice code reference: https://rolz.org/

Accepts possible standard DnD dice 2, 4, 6, 8, 10, 12, and 20 sided as well as 100 sided.

H - highest

L - lowest

+/- add or subtract number from total

x multiply result

GET /roll?string=4d20H3+3d4L1+12-3

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

GET /roll?string=4d21H3

    4 d21: 4 21 sided dice - not possible

Response:

```json
{
  "input": "4d21H3",
  "valid": false
}
```
