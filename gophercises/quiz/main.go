package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a CSV file in the format of 'question and answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	// csvFileName is a pointer to a filename
	flag.Parse()
	// _ = csvFileName // for Code Compilation
	file, err := os.Open(*csvFilename)
	// file (of the type ioReader) will be the actual filename of the file we wish to open
	if err != nil {
		exit(fmt.Sprintf("failed to open the CSV file: %s\n", *csvFilename))
	}
	// _ = file // for Code Compilation

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to read the CSV file.")
	}

	problems := parseLines(lines)
	// Create a new timer to track the time elapsed
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	// waits for the message from the channel
	// <-timer.C

	correct := 0
problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, p.q)
		// create an answer channel
		answerCh := make(chan string)
		// write a go routine:
		go func() {
			var answer string
			// reads the answer from the user.
			fmt.Scanf("%s\n", &answer)
			// sending answer (string) to answerCh (channel)
			answerCh <- answer
		}()
		/* */
		select {
		case <-timer.C: // waits for the message from the channel
			// fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
			fmt.Println()
			// return
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
}
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
