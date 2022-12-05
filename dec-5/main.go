//

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	partOneSolution := partOne("input.txt")
	// partTwoSolution := partTwo("input.txt")
	fmt.Println("part one:", partOneSolution)
	// fmt.Println("part two:", partTwoSolution)
}

// stacks!

[C]         [Q]         [V]    
[D]         [D] [S]     [M] [Z]
[G]     [P] [W] [M]     [C] [G]
[F]     [Z] [C] [D] [P] [S] [W]
[P] [L]     [C] [V] [W] [W] [H] [L]
[G] [B] [V] [R] [L] [N] [G] [P] [F]
[R] [T] [S] [S] [S] [T] [D] [L] [P]
[N] [J] [M] [L] [P] [C] [H] [Z] [R]

type Node struct {
	key  int
	next *Node
}

type Stack struct {
	head     *Node
	size     int
	capacity int
}


func (stacks *Stack) buildStackOfSize(n int, evensOnly bool) {
	one := []rune{'N', 'R', 'G', 'P'}
	two := []rune{'J', 'T', 'B', 'L', 'F', 'G', 'D', 'C'}
	three := []rune{'M', 'S', 'V'}
	four := []rune{'L', 'S', 'R', 'C', 'Z', 'P'}
	five := []rune{'P', 'S', 'L', 'V', 'C', 'W', 'D', 'Q'}
	six := []rune{'C', 'T', 'N', 'W', 'D', 'M', 'S'}
	seven := []rune{'H', 'G', 'D', 'W', 'P'}
	eight := []rune{'Z', 'L', 'P', 'H', 'S', 'C', 'M', 'V'}
	nine := []rune{'R', 'P', 'F', 'L', 'W', 'G', 'Z'}  
	for i := 0; i < n; i++ {
		stack.push(i)
		if evensOnly && i%2 == 1 {
			stack.pop()
		}
	}
}

func (stack *Stack) push(key int) bool {
	if stack.isFull() {
		// fmt.Println("ERROR: not enough space!")
		return false
	}
	node := Node{key, stack.head}
	stack.head = &node
	stack.size++
	return true
}

func (stack *Stack) peek() (*Node, bool) {
	if stack.len() > 0 {
		return stack.head, true
	}
	return nil, false
}

func (stack *Stack) pop() (*Node, bool) {
	if stack.size <= 0 {
		// fmt.Println("ERROR: cannot pop from empty stack!")
		return nil, false
	}
	node := stack.head
	stack.head = node.next
	stack.size--
	return node, true
}

func (stack *Stack) clear() {
	for stack.size > 0 {
		stack.pop()
	}
}

func (stack *Stack) isEmpty() bool {
	return stack.size == 0
}

func (stack *Stack) isFull() bool {
	return stack.size == stack.capacity
}

func (stack *Stack) len() int {
	return stack.size
}

func (stack *Stack) show(vertical bool) {
	for node := stack.head; node != nil; node = node.next {
		if vertical {
			fmt.Printf("\n%d\n---", node.key)
		} else {
			fmt.Printf("%d -> ", node.key)
		}
	}
	fmt.Printf("\n")
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

	numFullyContained := 0
	stacks := [9]Stack

	stacks.buildStack()

	for scanner.Scan() {
		// do something
	}
	return numFullyContained
}
