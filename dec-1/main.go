// package main solves advent of code for Dec. 1, 2022
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Problem:
// The jungle must be too overgrown and difficult to navigate in vehicles or access from the air;
// the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin
// taking inventory of their supplies. One important consideration is food - in particular,
// the number of Calories each Elf is carrying (your puzzle input).

// The Elves take turns writing down the number of Calories contained by the various meals,
// snacks, rations, etc. that they've brought with them, one item per line.
//Each Elf separates their own inventory from the previous Elf's inventory (if any) by a blank line.

// For example, suppose the Elves finish writing their items' Calories and end up with the following list:

// 1000
// 2000
// 3000

// 4000

// 5000
// 6000

// 7000
// 8000
// 9000

// 10000

// This list represents the Calories of the food carried by five Elves:

// The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
// The second Elf is carrying one food item with 4000 Calories.
// The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
// The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
// The fifth Elf is carrying one food item with 10000 Calories.
// In case the Elves get hungry and need extra snacks, they need to know which Elf to ask:
// they'd like to know how many Calories are being carried by the Elf carrying the most Calories.
// In the example above, this is 24000 (carried by the fourth Elf).

// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

// --- Part Two ---
// By the time you calculate the answer to the Elves' question, they've already realized
// that the Elf carrying the most Calories of food might eventually run out of snacks.

// To avoid this unacceptable situation, the Elves would instead like to know the total
// Calories carried by the top three Elves carrying the most Calories. That way, even if
// one of those Elves runs out of snacks, they still have two backups.

// In the example above, the top three Elves are the fourth Elf (with 24000 Calories),
// then the third Elf (with 11000 Calories), then the fifth Elf (with 10000 Calories).
// The sum of the Calories carried by these three elves is 45000.

// Find the top three Elves carrying the most Calories.
// How many Calories are those Elves carrying in total?

func main() {
	partOneSolution := partOne("input.txt")
	partTwoSolution := partTwo("input.txt")
	fmt.Println("part one:", partOneSolution)
	fmt.Println("part two:", partTwoSolution)
}

// <<<<<<<<<<<<<<<<<<<<<<<<< Part one >>>>>>>>>>>>>>>>>>>>>>>>>

func partOne(path string) int {
	// algorithm description:
	// two counter vars: count, max
	// read the data line-by-line
	// if a newline is reached:
	// 1. calculate the max: max = math.Max(max, curr)
	// 2. reset the counter to 0
	// else add the value to the counter

	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	count, highest := 0, 0

	// loop over the file line by line:
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			highest = int(math.Max(float64(count), float64(highest)))
			count = 0
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			count += val
		}
	}
	return highest
}

// <<<<<<<<<<<<<<<<<<<<<<<<< Part Two >>>>>>>>>>>>>>>>>>>>>>>>>
type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

// initialize a singly linked list of size 3
func initializeLinkedList() *LinkedList {
	list := LinkedList{}
	node := &Node{}
	list.head = node
	for i := 0; i < 2; i++ {
		node.next = &Node{}
		node = node.next
	}
	return &list
}

// traverse the sorted linked list from the tail (smallest value)
// and insert the new value where it belongs, preserving sorting.
func (list *LinkedList) insort(val int) {

	if val <= list.head.val {
		return
	}
	newNode := &Node{val: val}

	// find the insertion slot
	node := list.head
	var last *Node
	for node != nil && node.val < val {
		if node.next == nil {
			// newNode it's the new largest node
			node.next = newNode
			list.head = list.head.next
			return
		}
		last = node
		node = node.next
	}
	// if execution reaches this point, `node.val` > val.
	// newNode must be inserted in the first or second spot
	newNode.next = node
	if last == nil {
		// the new value is greater that the current smallest value.
		// It's the new head
		list.head = newNode
		return
	}
	// if execution gets here, the new node goes somewhere in the middle
	// `last` stores the last node whose val <= newNode.val
	// `node` stores the node whose val > newNode.val
	last.next = newNode
	newNode.next = node
	list.head = list.head.next
}

func (list *LinkedList) print() {
	node := list.head
	for node.next != nil {
		fmt.Printf("%v -> ", node.val)
		node = node.next
	}
	fmt.Printf("%v\n", node.val)
}

func (list *LinkedList) sum() int {
	count := 0
	node := list.head
	for node != nil {
		count += node.val
		node = node.next
	}
	return count
}

func partTwo(path string) int {
	// algorithm description:
	// * Part two uses similar code to part one in regards to
	//   reading the data line-by-line.
	// * This time only one counter is used and in place of the `highest`
	//   variable is a linked list of the 3 current highest calorie counts.
	// * Once a newline is reached, the `count` is sent to `insort()` which
	//   will insert the count into the linked list if it is greater
	//   than the current lowest value.
	// * The linked list is a singly linked list with a fixed length of 3 nodes.
	//   The nodes are sorted from smallest -> largest.
	//   `list.head` points to the smallest value

	// linked of size 3, sorted, walk backwards from the smallest to largest
	list := initializeLinkedList()

	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	count := 0

	// loop over the file line by line:
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			list.insort(count)
			count = 0
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			count += val
		}
	}
	list.print()
	return list.sum()
}
