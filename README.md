###goDice - An over-engineered and verbose go microservice for dice rolling

DONE:

TODO:

- Simple API for dice rolling based on string input.
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
  "input": "4d20H3+3d4L1+12-3",
  "valid": true,
  "dice": [
    {
      "size": 20,
      "count": 4,
      "high": 3,
      "rolls": [3, 5, 10, 12],
      "kept": [5, 10, 12],
      "subtotal": 27
    },
    {
      "size": 4,
      "count": 3,
      "low": 1,
      "rolls": [1, 3, 1],
      "kept": [1],
      "subtotal": 1
    }
  ],
  "modifier": 9,
  "total": 37,
  "image": "https://...."
}
```

GET /roll?string=4d21H3

    4 d21: 4 21 sided dice - not possible

Response:

```json
    {
        "input": "4d21H3"
        "valid": "false"
    }
```
