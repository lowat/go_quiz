package main

type quiz struct {
	problems []problem
}

type problem struct {
	description string
	answer      string
}

type state struct {
	quiz       *quiz
	numCorrect int
	inProgress bool
}
