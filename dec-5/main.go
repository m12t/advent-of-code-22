//

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	partOneSolution := partOne("input.txt")
	// partTwoSolution := partTwo("input.txt")
	fmt.Println("part one:", partOneSolution)
	// fmt.Println("part two:", partTwoSolution)
}

type Node struct {
	key  rune
	next *Node
}

type Stack struct {
	head     *Node
	size     int
	capacity int
}

func buildStack() [9]Stack {

	var stacks [9]Stack

	one := []rune{'N', 'R', 'G', 'P'}
	two := []rune{'J', 'T', 'B', 'L', 'F', 'G', 'D', 'C'}
	three := []rune{'M', 'S', 'V'}
	four := []rune{'L', 'S', 'R', 'C', 'Z', 'P'}
	five := []rune{'P', 'S', 'L', 'V', 'C', 'W', 'D', 'Q'}
	six := []rune{'C', 'T', 'N', 'W', 'D', 'M', 'S'}
	seven := []rune{'H', 'D', 'G', 'W', 'P'}
	eight := []rune{'Z', 'L', 'P', 'H', 'S', 'C', 'M', 'V'}
	nine := []rune{'R', 'P', 'F', 'L', 'W', 'G', 'Z'}

	all := [][]rune{one, two, three, four, five, six, seven, eight, nine}
	for i, a := range all {
		s := Stack{nil, len(a), 64}
		for _, v := range a {
			s.push(v)
		}
		stacks[i] = s
	}
	return stacks
}

func (stack *Stack) push(key rune) bool {
	if stack.isFull() {
		// fmt.Println("ERROR: not enough space!")
		return false
	}
	node := Node{key, stack.head}
	stack.head = &node
	stack.size++
	return true
}

func (stack *Stack) peek() rune {
	return stack.head.key
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
			fmt.Printf("\n%s\n---", string(node.key))
		} else {
			fmt.Printf("%s -> ", string(node.key))
		}
	}
	fmt.Printf("\n")
}

func move(stacks *[9]Stack, instruction string) {
	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllString(instruction, -1)
	var instructions [3]int
	for i, v := range matches {
		instructions[i], _ = strconv.Atoi(v)
	}
	instructions[1]--
	instructions[2]--
	for i := 0; i < instructions[0]; i++ {
		if node, ok := stacks[instructions[1]].pop(); ok {
			stacks[instructions[2]].push(node.key)
		} else {
			panic("whyy")
		}
	}
}

func partOne(path string) string {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	// loop over the file line by line:
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	stacks := buildStack()

	for scanner.Scan() {
		// do something
		move(&stacks, scanner.Text())
	}
	topCrates := ""
	for i := 0; i < 9; i++ {
		topCrates += string(stacks[i].peek())
	}
	return topCrates
}

// BCTSFHPLC not right
