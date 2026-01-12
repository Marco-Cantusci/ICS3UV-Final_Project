/**
 * @author Marco Cantusci
 * @version 2.0.0
 * @date 2026-01-05
 * @fileoverview Trivia game final project
 */

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
		loop := rand.Intn(len(list))
		// swap the two items picked from the randomizer
		list[counter], list[loop] = list[loop], list[counter]
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	// declare variables
	var totalString string
	var totalInt int
	var score int = 0
	var playAgain string = "yes"
	var grade string = ""

	// start replay loop
	for playAgain == "yes" {

		// shuffle questions
		shuffle(questions)

		// input amount of questions
		fmt.Print("How many questions would you like to answer(1-25)? ")
		totalString, _ = reader.ReadString('\n')
		totalString = strings.TrimSpace(totalString)
		totalInt, _ = strconv.Atoi(totalString)

		// question loop
		for counter := 0; counter < totalInt; counter++ {
			fmt.Println("\n" + questions[counter][0])
			var input string
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)

			input = strings.ToLower(strings.TrimSpace(input))

			// compare users answer to question answer
			if input == questions[counter][1] {
				fmt.Println("Correct!")
				score++
				} else {
					fmt.Println("Incorrect!")
				}
			}

		// calculate final percent
		var percent float64 = float64(score) / float64(totalInt) * 100

		if percent >= 90 {
			grade = "A"
		} else if percent >= 80 {
			grade = "B"
		} else if percent >= 70 {
			grade = "C"
		} else if percent >= 60 {
			grade = "D"
		} else {
			grade = "F"
		}

		// print final score
		fmt.Printf("\nGame Over!\nFinal Score: %d/%d(%.2f%%)\nGrade: %s\n", score, totalInt, percent, grade)

		score = 0

		// replay input
		fmt.Print("Would you like to play again? (yes/no): ")
		playAgain, _ = reader.ReadString('\n')
		playAgain = strings.TrimSpace(playAgain)
		playAgain = strings.ToLower(playAgain)
	}

	fmt.Println("\nThanks for playing!")

	fmt.Println("\nDone.")
}
