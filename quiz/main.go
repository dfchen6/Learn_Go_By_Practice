package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
)

type Record struct {
	given string
	actual string
	index int
}

func main() {
	filePath := "problems.csv"
	// load csv file
	f,_ := os.Open(filePath)

	problemSet := make(map[string]string)

	// create a new reader
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		question := strings.TrimSpace(record[0])
		answer := strings.TrimSpace(record[1])
		problemSet[question] = answer
	}

	var score = 0
	var total = len(problemSet)

	reader := bufio.NewReader(os.Stdin)

	for k, v := range problemSet {
		fmt.Printf(k + ": ")
		text,_ := reader.ReadString('\n')
		if (strings.TrimRight(text, "\n") == string(v)) {
			score++
		}
	}

	fmt.Printf("Total socre: %v/%v\n", score, total)
}