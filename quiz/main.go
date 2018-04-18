package main

import (
	"encoding/csv"
	"fmt"
	"flag"
	"os"
	"io"
	"bufio"
	"strings"
	"time"
	"log"
)

func main() {
	timeOut := flag.Int("timeout", 2, "time out without any input in seconds")
	filePath := flag.String("filepath", "problems.csv", "path of the file to read")
	flag.Parse()

	// load csv file
	f,_ := os.Open(*filePath)

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
		case <- time.After(time.Duration(*timeOut) * time.Second):
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