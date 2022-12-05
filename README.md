# advent-of-code-22

This repo houses my solutions to the 2022 [Advent of Code](https://adventofcode.com/) and my analysis.


#  Day 1: Calorie Counting
**Example input:**
```
7532
37124

37309

7616
2128
2657
8061
8565
7990
5085
6046
7685
5581

35538
10913
```
**Part A Problem:**
* Each line represents the calories of various items. Each grouping of items (successive non-empty lines) is the items carried by one elf.
* Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

**Notes:**
* null

**Part A Approach:**
1. Loop over the input file line by line, aggregating calories locally until an empty line is found
2. Recalculate the global max with a banchless operation like `highest = max(highest, local)`
3. Return `highest`

**Part A Analysis:**
* Time Complexity: `O(n)`
* Space Complexity: `O(1)`

<><><><><>

**Part B Problem:**
* Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?

**Notes:**
* Now the top 3 elves calories must be stored instead of just the single highest elf score. This makes the comparison of whether a local count is one of the highest 3 more challenging. I wanted to find a more elegant approach than brute forcing all 3 positions. I decided to implement a singly linked list due to it's `O(1)` insertion time complexity.
* An array would have also been satisfactory here due to the constant number of elves to track. The downside of array insertion could be overcome with a temp variable and manual swapping of elements.
* The head of the linked list would be the _lowest_ of the 3 highest elves' totals. This would allow "short circuiting" where the most common comparison is with the lowest of the top 3 elves. If this fails, there's no need to continue traversing the list.

**Part B Approach:**
1. Loop over the input file line by line, aggregating calories locally until an empty line is found
2. Pass the local `count` to a function, `insort()` that will insert into its place in the linked list if `count` is one of the new highest 3 calorie counts
3. Return the sum of the linked list


#  Day 2: Rock Paper Scissors
**Example Input:**
```
A Z
A Z
C Y
A X
```
**Part A Problem:**
* Each line represents a round of Rock, Paper, Scissors. The first letter is the opponent's play `{A: Rock, B: Paper, C: Scissors}`. The second letter is what I should play `{X: Rock, Y: Paper, Z: Scissors}`
* A win results in `6` points, a draw results in `3`, and there are no points for a loss
* There are also "gimme" points given out based on what move I make. Playing `Rock` gives `3`, choosing `Paper` gives `2` points, and `1` is awarded for playing `Scissors`
* What would your total score be if everything goes exactly according to your strategy guide?

**Notes:**
* I used 3 hashmaps to solve part A:
    ```go
    var convert = map[string]string{"A": "X", "B": "Y", "C": "Z"}
    var moveScores = map[string]int{"X": 1, "Y": 2, "Z": 3}
    var beatenBy = map[string]string{"X": "Z", "Y": "X", "Z": "Y"}
    ```
* The first, `convert`, does what it sounds. It unifies the encoding to facilitate comparisons
* The second, `moveScores` tracks the gimme points and is used to increment the score by the chosen move
* The final map, `beatenBy` encodes the essence of the game. In pseudocode, it reads like `beatenBy["Rock"] >>> "Scissors"`


**Part A Approach:**
1. To score each round:
    1. The inputs were mapped to the same encoding, using `convert`
    1. The gimme points were added
    1. The outcome was found by the following code:
        ```go
        if myMove == opponentMove {
            *score += drawPoints
        } else if opponentMove == beatenBy[myMove] {
            *score += winPoints
        } else {
        }
        ```
1. Return `score`

**Part A Analysis:**
Time Complexity: `O(n)`
Space Complexity: `O(1)`

<><><><><>

**Part B Problem:**
* Each line represents a round of Rock, Paper, Scissors. The first letter is still the opponent's play `{A: Rock, B: Paper, C: Scissors}`. However, the second letter is now the desired outcome `{X: Lose, Y: Draw, Z: Win}`. _I have to figure out what to play to ensure that outcome_
* A win results in `6` points, a draw results in `3`, and there are no points for a loss
* There are still "gimme" points given out based on what move I make. Playing `Rock` gives `3`, choosing `Paper` gives `2` points, and `1` is awarded for playing `Scissors`
* What would your total score be if everything goes exactly according to your new strategy guide?

**Notes:**
* I used 4 hashmaps to solve part B:
    ```go
    //   - The map `convert` maps the second move encoding {"A", "B", "C"}
    //     to {"X", "Y", "Z"} for easier comparisons between moves.
    var convert = map[string]string{"A": "X", "B": "Y", "C": "Z"}

    //   - The map `moveScores` maps the "gimme" values the
    //     player wins each round based on the choice of move.
    var moveScores = map[string]int{"X": 1, "Y": 2, "Z": 3}

    //   - The map `beatenBy` holds maps a move to the move it beats.
    //     eg. "Rock" ("X") beats "Scissors" ("Z")
    var beatenBy = map[string]string{"X": "Z", "Y": "X", "Z": "Y"}

    //   - Essentially just a reversed list of `beatenBy` that's used
    //     in part two when figuring out the move needed to win.
    var beats = map[string]string{"Z": "X", "X": "Y", "Y": "Z"}
    ```
"Scissors"`


**Part B Approach:**
1. To score each round:
    1. The inputs were mapped to the same encoding, using `convert`
    1. The outcome was found by the following code:
        ```go
        if outcome == "X" {
			// must lose
			score += moveScores[beatenBy[opponentMove]]
		} else if outcome == "Y" {
			// must tie
			score += moveScores[opponentMove] + drawPoints
		} else {
			// must win
			score += moveScores[beats[opponentMove]] + winPoints
		}
        ```
1. Return `score`

**Part B Analysis:**
Time Complexity: `O(n)`
Space Complexity: `O(1)`