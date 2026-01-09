/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2026-01-05
 * @fileoverview Trivia game final project
 */

package main

import (
	"fmt"
	"math/rand/v2"
)

// each question has the prompt and the correct response
var questions = [][]string{
	// ---- Computer Science ----
	{"What does CPU stand for?", "central processing unit"},
	{"What loop is used when you know how many times to repeat?", "for"},
	{"What symbol starts a single-line comment?", "//"},
	{"What does SSD stand for?", "solid state drive"},
	{"What data type stores true or false?", "boolean"},

	// ---- Math ----
	{"What is 6 times 12?", "72"},
	{"What is 15 percent of 200?", "30"},
	{"What is the square root of 81?", "9"},
	{"What is 12 squared?", "144"},
	{"What is 125 divided by 5?", "25"},

	// ---- Science ----
	{"What is the compound name for water?", "H2O"},
	{"What planet is closest to the Sun?", "mercury"},
	{"Which is the most abundant element in the universe?", "hydrogen"},
	{"What do you call an animal that eats a variety of other organisms, including plants, animals and fungi?", "omnivore"},
	{"What gas do humans breathe in?", "oxygen"},

	// ---- Geography ----
	{"What is the capital of Canada?", "ottawa"},
	{"What is the largest ocean on Earth?", "pacific"},
	{"Which country has the largest population in the world?", "china"},
	{"What is the capital of Mexico?", "mexico city"},
	{"What is the name of the largest country in the world?", "russia"},

	// ---- General ----
	{"How many days are in a leap year?", "366"},
	{"What is Arachnophobia a fear of?", "spiders"},
	{"In what year did the Titanic sink?", "1912"},
	{"Which instrument has 88 keys?", "piano"},
	{"What is the hardest natural substance on Earth?", "diamond"},
}

// shuffle function
func shuffle(list [][]string) {
	// loop through the list
	for counter := 0; counter < len(list); counter++ {
		// randomly pick from the list
		loop := rand.IntN(len(list))
		// swap the two items picked from the randomizer
		list[counter], list[loop] = list[loop], list[counter]
	}
}

func main() {
	// shuffle questions
	shuffle(questions)
	fmt.Println(questions)
	// declare variables
	// var total int
	// var score int = 0

	fmt.Println("\nDone.")
}
