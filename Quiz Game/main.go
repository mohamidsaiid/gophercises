package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type problem struct {
	ques string
	ans  string
}

func main() {
	fileName := flag.String("fileName", "test.csv", "get the name of csv file")
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problems := NewQuiz(records)

	//starting the test
	for {
		var s string
		fmt.Println("to Start the Test enter 's'")
		fmt.Scan(&s)
		if s == "s" {
			break
		}
	}
	timer := time.NewTimer(30 * time.Second)
	fmt.Println("time starts now you got 30 seconds till the end of the quiz")
	var correctAns int
	for _, problem := range problems {
		fmt.Printf("%s : ", problem.ques)
		answerCh := make(chan string)

		go getAnswer(answerCh)

		select {
		case answer := <-answerCh:
			if answer == problem.ans {
				correctAns++
			}
		case <-timer.C:
			fmt.Printf("\nTIME OUT!!\nquiz is finished\nyou got %d questions correct out of %d", correctAns, len(problems))
			return
		}
		close(answerCh)
	}
}

func NewQuiz(recs [][]string) []problem {
	problems := make([]problem, len(recs))
	for i, v := range recs {
		problems[i] = problem{
			ques: v[0],
			ans:  v[1],
		}
	}
	return problems
}

func getAnswer(answerCh chan string) {
	var s string
	fmt.Scan(&s)
	answerCh <- s
}
