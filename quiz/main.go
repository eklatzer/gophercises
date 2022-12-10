package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"gophercises/quiz/csv"
)

var inputFile string
var secondsForGame int
var shuffleQuestions bool

var correctAnswers int

func init() {
	flag.StringVar(&inputFile, "in", "problems.csv", "Path of file with the questions for the quiz")
	flag.IntVar(&secondsForGame, "time", 30, "Total time for the game in seconds")
	flag.BoolVar(&shuffleQuestions, "shuffle", true, "Whether the questions are shuffled each game")
	flag.Parse()
}

type question struct {
	Text   string `csv:"text"`
	Answer string `csv:"answer"`
}

func main() {
	csvParser := csv.NewCsvParser()

	var questions = []question{}
	err := csvParser.ParseFile(inputFile, &questions)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if shuffleQuestions {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(questions), func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
		})
	}

	log.Printf("are you ready for %d questions in max. %d seconds? press ENTER", len(questions), secondsForGame)
	fmt.Scanln()
	time.AfterFunc(time.Duration(secondsForGame)*time.Second, func() {
		log.Println("-----THE TIME IS OVER-----")
		printResult(correctAnswers, len(questions))
		os.Exit(0)
	})

	for i, question := range questions {
		log.Printf("question #%d: %s", (i + 1), question.Text)

		var answer string
		_, err := fmt.Scanln(&answer)
		if err != nil {
			log.Printf("failed to scan answer: %v", err)
			continue
		}

		if strings.TrimSpace(answer) == question.Answer {
			correctAnswers++
		}
	}
	printResult(correctAnswers, len(questions))
}

func printResult(correct, total int) {
	log.Println("-----GAME STATS-----")
	log.Printf("correct questions: %d/%d\n", correct, total)
}
