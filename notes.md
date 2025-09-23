# 2025-09-23
Creating and more refine JSON response. First entry is a pretty string. 

Committing before working on "recommended action"


# 2025-09-22
Created basic main.go that prints to terminal one dealer upcard, and 2 player cards. From a 6-deck shoe, and a game struct.

Now output JSON with the dealer card and the 2 player cards (string only)

Next steps:
-refine the JSON output: keep string for "pretty", and integer values (no suit) for future decisions on strategy
-make this into a simple API with a single GET (outputs JSON)
-Create "perfect21" JSON (for H17 game, or S17 game or other)
-Add a POST to API for selecting the one correct action

