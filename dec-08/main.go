// --- Day 8: Treetop Tree House ---
// The expedition comes across a peculiar patch of tall trees all planted carefully in a grid. The Elves explain that a previous expedition planted these trees as a reforestation effort. Now, they're curious if this would be a good location for a tree house.

// First, determine whether there is enough tree cover here to keep a tree house hidden. To do this, you need to count the number of trees that are visible from outside the grid when looking directly along a row or column.

// The Elves have already launched a quadcopter to generate a map with the height of each tree (your puzzle input). For example:

// 30373
// 25512
// 65332
// 33549
// 35390
// Each tree is represented as a single digit whose value is its height, where 0 is the shortest and 9 is the tallest.

// A tree is visible if all of the other trees between it and an edge of the grid are shorter than it. Only consider trees in the same row or column; that is, only look up, down, left, or right from any given tree.

// All of the trees around the edge of the grid are visible - since they are already on the edge, there are no trees to block the view. In this example, that only leaves the interior nine trees to consider:

// The top-left 5 is visible from the left and top. (It isn't visible from the right or bottom since other trees of height 5 are in the way.)
// The top-middle 5 is visible from the top and right.
// The top-right 1 is not visible from any direction; for it to be visible, there would need to only be trees of height 0 between it and an edge.
// The left-middle 5 is visible, but only from the right.
// The center 3 is not visible from any direction; for it to be visible, there would need to be only trees of at most height 2 between it and an edge.
// The right-middle 3 is visible from the right.
// In the bottom row, the middle 5 is visible, but the 3 and 4 are not.
// With 16 trees visible on the edge and another 5 visible in the interior, a total of 21 trees are visible in this arrangement.

// Consider your map; how many trees are visible from outside the grid?

// NOTES:
// This should be a straightforward O(n) solution. Can branchless make it faster?

package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	maxTreeHeight = 9
)

type Grid [][]int

var (
	forest  Grid
	counted = make(map[string]bool)
)

func main() {
	partOneSolution, partTwoSolution := solve("input.txt")
	fmt.Println("part one:", partOneSolution)
	fmt.Println("part two:", partTwoSolution)
}

func (forest *Grid) findVisible() int {
	// * Need to loop over 4x and see the number of visible trees.
	// * No optimizations for the time being
	numVisible, sideLength := 0, len(*forest)
	numVisible += forest.findVisibleFromEastWest(sideLength)
	numVisible += forest.findVisibleFromNorthSouth(sideLength)
	return numVisible
}

func (forest *Grid) findVisibleFromEastWest(sideLength int) int {
	numVisible := 0 // the first row is all visible
	for row := 0; row < sideLength; row++ {
		tallestInEast, tallestInWest := -1, -1
		for col := 0; col < sideLength; col++ {
			westTree, eastTree := (*forest)[row][col], (*forest)[row][sideLength-col-1]
			if westTree > tallestInWest {
				tallestInWest = westTree
				key := fmt.Sprintf("[%d][%d]", row, col)
				if _, present := counted[key]; !present {
					numVisible += 1
					counted[key] = true
				}
			}
			if eastTree > tallestInEast {
				tallestInEast = eastTree
				key := fmt.Sprintf("[%d][%d]", row, sideLength-col-1)
				if _, present := counted[key]; !present {
					numVisible += 1
					counted[key] = true
				}
			}
		}
	}
	return numVisible
}

func (forest *Grid) findVisibleFromNorthSouth(sideLength int) int {
	numVisible := 0 // the first row is all visible
	for col := 0; col < sideLength; col++ {
		tallestInNorth, tallestInSouth := -1, -1
		for row := 0; row < sideLength; row++ {
			northTree, southTree := (*forest)[row][col], (*forest)[sideLength-row-1][col]
			if northTree > tallestInNorth {
				tallestInNorth = northTree
				key := fmt.Sprintf("[%d][%d]", row, col)
				if _, present := counted[key]; !present {
					numVisible += 1
					counted[key] = true
				}
			}
			if southTree > tallestInSouth {
				tallestInSouth = southTree
				key := fmt.Sprintf("[%d][%d]", sideLength-row-1, col)
				if _, present := counted[key]; !present {
					numVisible += 1
					counted[key] = true
				}
			}
		}
	}
	return numVisible
}

func solve(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// loop over the file line by line:
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, v := range line {
			row = append(row, int(v-'0'))
		}
		forest = append(forest, row)
	}

	return forest.findVisible(), 0
}
