package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER when ready."

func main() {
	var firstNumber = 2
	var secondNumber = 5
	var subraction = 7
	// var answer int

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

	fmt.Println("Now subtract", subraction, prompt)
	reader.ReadString('\n')

	// give the answer
}