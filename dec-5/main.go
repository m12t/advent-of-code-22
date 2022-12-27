//

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// used to red the input and build the stack
const (
	firstCrateIdx       = 1
	crateOffset         = 4
	firstLetterByteCode = 65
)

func main() {
	partOneSolution := partOne("full_input.txt")
	// partTwoSolution := partTwo("input.txt")
	fmt.Println("part one:", partOneSolution)
	// fmt.Println("part two:", partTwoSolution)
}

type Node struct {
	key  byte
	next *Node
}

type Stack struct {
	head *Node
	size int
}

func buildStack(inputs *[9][]byte) [9]Stack {

	var stacks [9]Stack

	for i, a := range inputs {
		s := Stack{nil, len(a)}
		for _, v := range a {
			s.push(v)
		}
		stacks[i] = s
	}
	return stacks
}

func (stack *Stack) push(key byte) bool {
	node := Node{key, stack.head}
	stack.head = &node
	stack.size++
	return true
}

func (stack *Stack) peek() byte {
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
	tmp := Stack{nil, 0}
	for i := 0; i < instructions[0]; i++ {
		if node, ok := stacks[instructions[1]].pop(); ok {
			tmp.push(node.key)
		}
	}
	for i := 0; i < instructions[0]; i++ {
		if node, ok := tmp.pop(); ok {
			stacks[instructions[2]].push(node.key)
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

	var inputs [9][]byte

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		for i := 0; i < 9; i++ {
			r := line[i*crateOffset+1]
			if r == ' ' || r < firstLetterByteCode {
				continue
			}
			inputs[i] = append(inputs[i], r)
		}
	}
	stackArray := buildStack(&inputs)
	for i := 0; i < 9; i++ {
		stackArray[i].show(false)
	}

	for scanner.Scan() {
		// do something
		// stackBuilt = true
		move(&stackArray, scanner.Text())
	}
	topCrates := ""
	for i := 0; i < 9; i++ {
		topCrates += string(stackArray[i].peek())
	}
	return topCrates
}
