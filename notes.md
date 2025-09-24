# 2025-09-24
Created local API (single GET response), single route at /play/ 
(this was simply a matter of re-factoring some of the code and using GIN)

Deployed API: 
https://perfect-21-api.onrender.com/play

Created frontend (vite or nextjs) to display JSON response and "pretty cards"
(on refresh)

Created Action buttons and logic to "Quiz" (true / false) component with Vite buttons

Deployed on Vercel


# 2025-09-23a
Creating and more refined JSON response. First entry is a pretty string. 

"Scenario" object create. Will use to pass to "DetermineAction" function

Scenario object creation re-factored in helper function.

Committing before working on "recommended action"

Perfect strategy PDF downloaded and added to repo. 

Resource: https://www.blackjackapprenticeship.com/blackjack-strategy-charts/

Created the recommended action function. 

Now JSON response "works as is". Will work on the API next. 

(thinking: probably only need a GET route ... frontend will be like a quiz)

{
  "dealer": {
    "prettyString": [
      "4♥",
      "?"
    ],
    "upCardValue": 4,
    "holeCardValue": 10,
    "isBlackJack": false
  },
  "player": {
    "prettyString": [
      "4♦",
      "K♥"
    ],
    "card1Value": 4,
    "card2Value": 10,
    "playerTotal": 14,
    "isPair": false,
    "isSoft": false,
    "isBlackJack": false
  },
  "correctAction": "STAND"
}


# 2025-09-22
Created basic main.go that prints to terminal one dealer upcard, and 2 player cards. From a 6-deck shoe, and a game struct.

Now output JSON with the dealer card and the 2 player cards (string only)

Next steps:
-refine the JSON output: keep string for "pretty", and integer values (no suit) for future decisions on strategy
-make this into a simple API with a single GET (outputs JSON)
-Create "perfect21" JSON (for H17 game, or S17 game or other)
-Add a POST to API for selecting the one correct action

