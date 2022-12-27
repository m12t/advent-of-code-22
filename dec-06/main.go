// --- Day 6: Tuning Trouble ---
// The preparations are finally complete; you and the Elves leave camp on foot and begin to make your way toward the star fruit grove.

// As you move through the dense undergrowth, one of the Elves gives you a handheld device. He says that it has many fancy features, but the most important one to set up right now is the communication system.

// However, because he's heard you have significant experience dealing with signal-based systems, he convinced the other Elves that it would be okay to give you their one malfunctioning device - surely you'll have no problem fixing it.

// As if inspired by comedic timing, the device emits a few colorful sparks.

// To be able to communicate with the Elves, the device needs to lock on to their signal. The signal is a series of seemingly-random characters that the device receives one at a time.

// To fix the communication system, you need to add a subroutine to the device that detects a start-of-packet marker in the datastream. In the protocol being used by the Elves, the start of a packet is indicated by a sequence of four characters that are all different.

// The device will send your subroutine a datastream buffer (your puzzle input); your subroutine needs to identify the first position where the four most recently received characters were all different. Specifically, it needs to report the number of characters from the beginning of the buffer to the end of the first such four-character marker.

// For example, suppose you receive the following datastream buffer:

// mjqjpqmgbljsphdztnvjfqwrcgsmlb
// After the first three characters (mjq) have been received, there haven't been enough characters received yet to find the marker. The first time a marker could occur is after the fourth character is received, making the most recent four characters mjqj. Because j is repeated, this isn't a marker.

// The first time a marker appears is after the seventh character arrives. Once it does, the last four characters received are jpqm, which are all different. In this case, your subroutine should report the value 7, because the first start-of-packet marker is complete after 7 characters have been processed.

// Here are a few more examples:

// bvwbjplbgvbhsrlpgdmjqwftvncz: first marker after character 5
// nppdvjthqldpwncqszvftbrmjlhg: first marker after character 6
// nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: first marker after character 10
// zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: first marker after character 11
// How many characters need to be processed before the first start-of-packet marker is detected?

// NOTES:
// - Since there's no linebreaks, there should be a delimiter to read bit-sized chunks into memory.
// - While the whole file could be read into memory, streaming is more efficient.

// --- Part Two ---
// Your device's communication system is correctly detecting packets, but still isn't working. It looks like it also needs to look for messages.

// A start-of-message marker is just like a start-of-packet marker, except it consists of 14 distinct characters rather than 4.

// Here are the first positions of start-of-message markers for all of the above examples:

// mjqjpqmgbljsphdztnvjfqwrcgsmlb: first marker after character 19
// bvwbjplbgvbhsrlpgdmjqwftvncz: first marker after character 23
// nppdvjthqldpwncqszvftbrmjlhg: first marker after character 23
// nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: first marker after character 29
// zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: first marker after character 26
// How many characters need to be processed before the first start-of-message marker is detected?

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const (
	chunksize   = 1024
	packetSize  = 4
	messageSize = 14
)

var (
	count int
	err   error
)

func main() {
	partOneSolution := solve("input.txt", packetSize)
	partTwoSolution := solve("input.txt", messageSize)
	fmt.Println("part one:", partOneSolution)
	fmt.Println("part two:", partTwoSolution)
}

func solve(path string, size int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// loop over the file line by line:
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	section := make([]byte, chunksize)
	window := make([]byte, size)
	hashmap := make(map[byte]int)
	i := 0

	for {
		if count, err = reader.Read(section); err != nil {
			break
		}
		buffer.Write(section[:count])
		section = section[:count]
		for _, c := range section {
			// * Use an array to keep track of a rolling window
			//   of bytes.
			// * A map is used with a count of occurances, since
			//   a "set" with boolean values will run into issues
			//   if the expired value to be deleted also (rightfully)
			//   appears later in the window. This will cause the
			//   correct instance of the char to be removed as well.
			// * By storing a value, it is easy to identify
			//   duplicated values within the window by checking
			//   the value of the map for that char.
			hashmap[window[i%size]]--
			if hashmap[window[i%size]] < 1 {
				delete(hashmap, window[i%size])
			}
			window[i%size] = c
			hashmap[c]++
			if len(hashmap) == size {
				return i + 1
			}
			i++
		}
	}
	return -1
}
