# advent-of-code-22

This repo houses my solutions to the 2022 [Advent of Code](https://adventofcode.com/) and my analysis.


##  Day 1: Calorie Counting
**Example input:**
Each line represents the calories of various items. Each grouping of items (successive non-empty lines) is the items carried by one elf.
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
Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

**Notes:**
null

**Part A Approach:**
1. Loop over the input file line by line, aggregating calories locally until an empty line is found
2. Recalculate the global max with a banchless operation like `highest = max(highest, local)`
3. Return `highest`

**Part A Analysis:**
Time Complexity: `O(n)`
Space Complexity: `O(1)`

<><><><><>

**Part B Problem:**
Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?

**Notes:**
Now the top 3 elves calories must be stored instead of just the single highest elf score. This makes the comparison of whether a local count is one of the highest 3 more challenging. I wanted to find a more elegant approach than brute forcing all 3 positions. I decided to implement a singly linked list due to it's `O(1)` insertion time complexity. An array would have also been satisfactory here due to the constant number of elves to track. The downside of array insertion could be overcome with a temp variable and manual swapping of elements. Nonetheless, the head of the linked list would be the _lowest_ of the 3 highest elves' totals. This would allow "short circuiting" where the most common comparison is with the lowest of the top 3 elves. If this fails, there's no need to continue traversing the list.

**Part B Approach:**
1. Loop over the input file line by line, aggregating calories locally until an empty line is found
2. Pass the local `count` to a function, `insort()` that will insert into its place in the linked list if `count` is one of the new highest 3 calorie counts
3. Return the sum of the linked list


