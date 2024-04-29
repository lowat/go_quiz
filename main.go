package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	quiz, err := readProblemsFromCSV("problems.csv")
	if err != nil {
		log.Fatal("Error generating quiz from csv", err)
	}
	// for i, problem := range quiz.problems {
	// 	fmt.Printf("Question %v: %v\n", i+1, problem.description)
	// }

	state := newState(quiz)
	startQuiz(state)
	quizResults(state)

}

func readProblemsFromCSV(filepath string) (*quiz, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var problems []problem

	for _, record := range records {
		currProblem := problem{
			description: record[0],
			answer:      record[1],
		}
		problems = append(problems, currProblem)
	}

	quiz := quiz{
		problems: problems,
	}

	return &quiz, nil
}

func newState(quiz *quiz) *state {
	return &state{
		quiz:       quiz,
		numCorrect: 0,
		inProgress: false,
	}
}

func startQuiz(state *state) {
	reader := bufio.NewScanner(os.Stdin)
	for i, problem := range state.quiz.problems {
		fmt.Printf("Question %v: %v\n", i+1, problem.description)
		reader.Scan()
		userAnswer := cleanedInput(reader.Text())
		if isAnswerCorrect(userAnswer, problem.answer) {
			state.numCorrect += 1
		}
	}
}

func isAnswerCorrect(userAnswer, correctAnswer string) bool {
	return userAnswer == correctAnswer
}

func quizResults(state *state) {
	fmt.Printf("The final score is: %.2f", 100.0*float32(state.numCorrect)/float32(len(state.quiz.problems)))
}

func cleanedInput(userInput string) string {
	return strings.Fields(strings.ToLower(userInput))[0]
}
