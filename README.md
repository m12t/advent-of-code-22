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

**Part A Notes:**

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

**Part B Notes:**
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

**Part A Notes:**
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
            score += drawPoints
        } else if opponentMove == beatenBy[myMove] {
            score += winPoints
        } else {
        }
        ```

**Part A Analysis:**
* Time Complexity: `O(n)`
* Space Complexity: `O(1)`

<><><><><>

**Part B Problem:**
* Each line represents a round of Rock, Paper, Scissors. The first letter is still the opponent's play `{A: Rock, B: Paper, C: Scissors}`. However, the second letter is now the desired outcome `{X: Lose, Y: Draw, Z: Win}`. _I have to figure out what to play to ensure that outcome_
* A win results in `6` points, a draw results in `3`, and there are no points for a loss
* There are still "gimme" points given out based on what move I make. Playing `Rock` gives `3`, choosing `Paper` gives `2` points, and `1` is awarded for playing `Scissors`
* What would your total score be if everything goes exactly according to your new strategy guide?

**Part B Notes:**
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
* Time Complexity: `O(n)`
* Space Complexity: `O(1)`


#  Day 3: Rucksack Reorganization
**Example Input:**
```
DMwrszrfMzSSCpLpfCCn
RMvhZhQqlvhMvRtbvbcPclPlncddppLTdppd
tVMQhFtjjWmsFJsmsW
trRtvNhfJhSzzSTFVhQQZQhHGphP
CnLMBWLwDMgMcwwdngdHGPVTQGpTHZdGPGpd
LLDqcDgwqCMnLWqtvzrzbbtJqPjJ
wQQwHNQLmbWQbQRHwHNFBbwqPfjqlzRMGRqzpSfvPlzplM
nCtGCZZtsGsrtDMZpfMpSlMlvlZq
cJctJCgVJsCJnDTsCthGhGLwBWBbbQmLbgQLQQdWbbbQ
```
**Part A Problem:**
* Each line represents the items in two rucksacks carried by one elf (characters are case-sensitive).
* The lines contain an even number of characters and both rucksacks have the same number of items inside.
* Items have a "priority" associated with them that corresponds to the following: `a...z : 1...26, A...Z : 27...52`
* Find the item type that appears in both compartments of each rucksack. What is the sum of the priorities of those item types?

**Part A Notes:**

**Part A Approach:**
1. Split each string into the two rucksacks by splitting the string in two
1. Find the item in both rucksacks by:
    1. Creating a set of the first rucksack
    1. Iterating over the items in the second rucksack and checking if each item is in the first rucksack. When a match is found, return the `rune` of the item.
1. Use the following logic to find the "priority" for the char, where you get the ascii value of the character and use a constant offset that maps the char to the desired "priority" of the problem:
    ```go
    if unicode.IsUpper(c) {
		totalPriority += int(c) - upperOffset
	} else {
		totalPriority += int(c) - lowerOffset
	}
    ```

**Part A Analysis:**
* Time Complexity: `O(n)`
* Space Complexity: `O(m)` the sets take up linear extra space with respect to the number of items in the rucksack, `m`

<><><><><>

**Part B Problem:**
* Each 3 lines represents a group of elves and their respective rucksacks.
* Find the one item that all 3 elves possess

**Part B Notes:**


**Part B Approach:**
1. Loop over the input and create an array of the groups rucksacks as 3 strings. The array looks like `var group [3]string`
1. Once the three rucksacks are added to the group, find the common item by:
    ```go
    func findGroupBadge(rucksacks *[3]string) rune {
        // frequency is a global map that stores the
        // count associated with a rune where a single
        // rucksack can only increment an item once.
        frequency := make(map[rune]int)
        for i := 0; i < 3; i++ {
            // loop over each rucksack and create a local
            // set `s` to eliminate duplicate items within
            // the same rucksack
            s := make(map[rune]bool)
            for _, c := range rucksacks[i] {
                // loop over each rune in the rucksack
                if _, found := s[c]; !found {
                    // if this item hasn't yet been seen
                    // in this particular rucksack, increment
                    // the global map `frequency`
                    frequency[c]++
                }
                s[c] = true
            }
        }

        for rune, count := range frequency {
            if count == 3 {
                return rune
            }
        }
	    panic("No badge found!")
    }
    ```
1. Find the "priority" for the common item using the same logic from part A

**Part B Analysis:**
* Time Complexity: `O(n*m)`, where the number of items in a rucksack is `n` and the number of elves in a grouop is `m`.
    * In this case, there are a constant number of elves, 3, and so the overall asymptotic time complexity is just `O(n)`.
* Space Complexity: `O(n)`

#  Day 4: Camp Cleanup
**Example Input:**
```
2-4,6-8
1-71,2-71
11-74,74-75
44-96,43-96
79-79,3-78
```
**Part A Problem:**
* Each line represents the region of the camp that two elves are assigned to clean. The two ranges are separated by a comma.
* Within the first pair of Elves, the first Elf was assigned sections 2-4 (sections 2, 3, and 4), while the second Elf was assigned sections 6-8 (sections 6, 7, 8).
* Some of the pairs have noticed that one of their assignments fully contains the other. For example, 2-8 fully contains 3-7, and 6-6 is fully contained by 4-6. In pairs where one assignment fully contains the other, one Elf in the pair would be exclusively cleaning sections their partner will already be cleaning, so these seem like the most in need of reconsideration. In this example, there are 2 such pairs.
* In how many assignment pairs does one range fully contain the other?

**Part A Notes:**


**Part A Approach:**
1. Read the input into a 2D array of two assignments and their respective low and high bounds.
    * eg `2-4,6-8` converts to `[[2, 4],[6,8]]`
1. First check for the cases that are guaranteed to be fully contained:
    ```go
    if assignments[1][0] == assignments[0][0] ||
		assignments[1][1] == assignments[0][1] {
		// * The sections are guaranteed to overlap
		//   since they have the same lower or upper.
		return 1
	}
    ```
1. Find the array index of the assignment with the lesser low bound:
    ```go
    var x uint8 = 0
	if assignments[0][0] > assignments[1][0] {
		// * Set `x` to the array index of the
		//   range with the lower starting bound.
		x = 1
	}
    ```
1. If the upper bound of the array with the lesser low bound is greater than or equal to the upper bound of the other array, the second array is fully contained. Here there's some bitwise trickery to `xor` the value of `x` with 1 (which will flip 0 -> 1 and 1 -> 0):
    ```go
    if assignments[x][1] >= assignments[x^1][1] {
		return 1
	}
	return 0
    ```

**Part A Analysis:**
* Time Complexity: `O(n)`
* Space Complexity: `O(1)`

<><><><><>

**Part B Problem:**
* It seems like there is still quite a bit of duplicate work planned. Instead, the Elves would like to know the number of pairs that overlap at all.
* In how many assignment pairs do the ranges overlap?

**Part B Notes:**


**Part B Approach:**
1. Repeat the same steps from part A but with the addition of some logic:
    ```go
    if partialOverlap && assignments[x][1] >= assignments[x^1][0] {
		return 1
	}
    ```
    * If the upper bound of the assignment _with the lesser low bound_ is greater than or equal to the lower bound of the assignment with the _greater low bound_, there is overlap.

**Part B Analysis:**
* Time Complexity: `O(n)`
* Space Complexity: `O(1)`
