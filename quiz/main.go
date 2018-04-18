package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"time"
	"log"
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
	var counter = 1

	inputChan := make(chan string)
	go getInput(inputChan)

	for k, v := range problemSet {
		fmt.Printf("Problem %v: %v = ", counter, k)
		counter++
		select {
		case text := <-inputChan:
			if (strings.TrimRight(text, "\n") == string(v)) {
				score++
			}
		case <- time.After(2000 * time.Millisecond):
			fmt.Printf("Total socre: %v/%v\n", score, total)
			return
		}
	}

	fmt.Printf("Total socre: %v/%v\n", score, total)
}

func getInput(input chan string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input <- text
	}
}