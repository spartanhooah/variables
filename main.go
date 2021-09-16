package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	rand.Seed(time.Now().UnixNano())
	var firstNumber = random1to10()
	var secondNumber = random1to10()
	var subtraction = random1to10()
	answer := firstNumber*secondNumber - subtraction

	runTheGame(firstNumber, secondNumber, subtraction, answer)
}

func random1to10() int {
	return rand.Intn(8) + 2
}

func runTheGame(firstNumber, secondNumber, subtraction, answer int) {
	reader := bufio.NewReader(os.Stdin)

	// display a welcome message with instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("\nThink of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take user through the game
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give the answer
	fmt.Println("The answer is", answer)
}
