// --- Day 3: Rucksack Reorganization ---
// One Elf has the important job of loading all of the rucksacks with
// supplies for the jungle journey. Unfortunately, that Elf didn't quite
// follow the packing instructions, and so a few items now need to
// be rearranged.

// Each rucksack has two large compartments. All items of a given type
// are meant to go into exactly one of the two compartments.
// The Elf that did the packing failed to follow this rule for
// exactly one item type per rucksack.

// The Elves have made a list of all of the items currently in each
// rucksack (your puzzle input), but they need your help finding the errors.
// Every item type is identified by a single lowercase or uppercase letter
// (that is, a and A refer to different types of items).

// The list of items for each rucksack is given as characters all on a single line.
// A given rucksack always has the same number of items in each of its two compartments,
// so the first half of the characters represent items in the first compartment,
// while the second half of the characters represent items in the second compartment.

// For example, suppose you have the following list of contents from six rucksacks:

// vJrwpWtwJgWrhcsFMMfFFhFp
// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
// PmmdzqPrVvPwwTWBwg
// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
// ttgJtRGJQctTZtZT
// CrZsJsPPZsGzwwsLwLmpwMDw
// The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its
// first compartment contains the items vJrwpWtwJgWr, while the second compartment
// contains the items hcsFMMfFFhFp. The only item type that appears in both compartments
// is lowercase p.
// The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL.
// The only item type that appears in both compartments is uppercase L.
// The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common
// item type is uppercase P.
// The fourth rucksack's compartments only share item type v.
// The fifth rucksack's compartments only share item type t.
// The sixth rucksack's compartments only share item type s.
// To help prioritize item rearrangement, every item type can be converted to a priority:

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
// In the above example, the priority of the item type that appears in both compartments
// of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s);
// the sum of these is 157.

// Find the item type that appears in both compartments of each rucksack.
// What is the sum of the priorities of those item types?

// --- Part Two ---
// As you finish identifying the misplaced items, the Elves come to you with another issue.

// For safety, the Elves are divided into groups of three. Every Elf carries a badge that
// identifies their group. For efficiency, within each group of three Elves, the badge is the
// only item type carried by all three Elves. That is, if a group's badge is item type B,
// then all three Elves will have item type B somewhere in their rucksack, and at most two
// of the Elves will be carrying any other item type.

// The problem is that someone forgot to put this year's updated authenticity
// sticker on the badges. All of the badges need to be pulled out of the
// rucksacks so the new authenticity stickers can be attached.

// Additionally, nobody wrote down which item type corresponds to each group's badges.
// The only way to tell which item type is the right one is by finding the one item type
// that is common between all three Elves in each group.

// Every set of three lines in your list corresponds to a single group,
// but each group can have a different badge item type. So, in the above example,
// the first group's rucksacks are the first three lines:

// vJrwpWtwJgWrhcsFMMfFFhFp
// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
// PmmdzqPrVvPwwTWBwg
// And the second group's rucksacks are the next three lines:

// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
// ttgJtRGJQctTZtZT
// CrZsJsPPZsGzwwsLwLmpwMDw

// In the first group, the only item type that appears in all three rucksacks
// is lowercase r; this must be their badges. In the second group,
// their badge item type must be Z.

// Priorities for these items must still be found to organize the sticker attachment
// efforts: here, they are 18 (r) for the first group and 52 (Z) for the second group.
// The sum of these is 70.

// Find the item type that corresponds to the badges of each three-Elf group.
// What is the sum of the priorities of those item types?

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const (
	upperOffset = 38
	lowerOffset = 96
)

func main() {
	partOneSolution := partOne("input.txt")
	partTwoSolution := partTwo("input.txt")
	fmt.Println("part one:", partOneSolution)
	fmt.Println("part two:", partTwoSolution)
}

func incrementPriority(c rune, totalPriority *int) {
	// a ... z : 1 ... 26
	// A ... Z : 27 ... 52
	if unicode.IsUpper(c) {
		*totalPriority += int(c) - upperOffset
	} else {
		*totalPriority += int(c) - lowerOffset
	}
}

func findDupe(rucksackOne, rucksackTwo string) rune {
	seen := make(map[rune]bool)
	for _, c := range rucksackOne {
		// build a set of the "items" in the first rucksack
		seen[c] = true
	}
	for _, c := range rucksackTwo {
		// * Iterate over the second rucksack and look for dupes
		//   in O(n) time by comparing with the set from above
		if seen[c] {
			return c
		}
	}
	panic("No dupes found!")
}

func partOne(path string) int {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	// loop over the file line by line:
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	totalPriority := 0

	for scanner.Scan() {
		line := scanner.Text()
		mid := len(line) / 2
		rucksackOne, rucksackTwo := line[:mid], line[mid:]
		dupe := findDupe(rucksackOne, rucksackTwo)
		incrementPriority(dupe, &totalPriority)
	}
	return totalPriority

}

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

func partTwo(path string) int {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	// loop over the file line by line:
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	i, totalPriority := 0, 0
	var group [3]string

	for scanner.Scan() {
		rucksack := scanner.Text()
		group[i] = rucksack
		i++
		if i%3 == 0 {
			dupe := findGroupBadge(&group)
			incrementPriority(dupe, &totalPriority)
			i = 0
		}
	}
	return totalPriority

}
