/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2026-01-05
 * @fileoverview Trivia game final project
 */
// Declare variables
let score: number = 0;

const questions: string[][] = [
  // ---- Computer Science ----
  ["What does CPU stand for?", "central processing unit"],
  ["What loop is used when you know how many times to repeat?", "for"],
  ["What symbol starts a single-line comment?", "//"],
  ["What does SSD stand for?", "solid state drive"],
  ["What data type stores true or false?", "boolean"],

  // ---- Math ----
  ["What is 6 times 12?", "72"],
  ["What is 15 percent of 200?", "30"],
  ["What is the square root of 81?", "9"],
  ["What is 12 squared?", "144"],
  ["What is 125 divided by 5?", "25"],

  // ---- Science ----
  ["What is the compound name for water?", "h2o"],
  ["What planet is closest to the Sun?", "mercury"],
  ["Which is the most abundant element in the universe?", "hydrogen"],
  ["What do you call an animal that eats a variety of other organisms, including plants, animals and fungi?", "omnivore"],
  ["What gas do humans breathe in?", "oxygen"],

  // ---- Geography ----
  ["What is the capital of Canada?", "ottawa"],
  ["What is the largest ocean on Earth?", "pacific"],
  ["Which country has the largest population in the world?", "china"],
  ["What is the capital of Mexico?", "mexico city"],
  ["What is the name of the largest country in the world?", "russia"],

  // ---- General ----
  ["How many days are in a leap year?", "366"],
  ["What is Arachnophobia a fear of?", "spiders"],
  ["In what year did the Titanic sink?", "1912"],
  ["Which instrument has 88 keys?", "piano"],
  ["What is the hardest natural substance on Earth?", "diamond"],
];

// shuffle function
function shuffle(list: string[][]): void {
  list.sort(() => Math.random() - 0.5);
}

// Shuffle array
shuffle(questions);

// input amount of questions
const totalString: string =
  prompt("How many questions would you like to answer(1-25)?") || ("25");

const totalNumber: number = parseInt(totalString);

// question loop
for (let counter: number = 0; counter < totalNumber; counter++) {
  // process user input
  const input = prompt(questions[counter][0]);
  if (input === null) break;
  const userInput = input.toLowerCase().trim();

  // compare users answer to question answer
  if (userInput === questions[counter][1]) {
    console.log("Correct!");
    score++;
  } else {
    console.log("Incorrect!");
  }
}
// calculate final percent
const percent: number = score / totalNumber * 100;

// print final score
console.log(
  `Game Over!\nFinalScore: ${score}/${totalNumber}(${percent.toFixed(2)}%)`,
);

console.log("\nDone.");
