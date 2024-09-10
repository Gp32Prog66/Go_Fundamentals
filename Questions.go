package main

/*
Movie Script from Genius.com
https://genius.com/Monty-python-monty-python-and-the-holy-grail-bridge-of-death-annotated
*/

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	fmt.Println(input)

	var name string
	var quest string
	var color string

	var capital string
	var finalQuestion string

	//Questions
	fmt.Print("What is your name? ")
	fmt.Scanln(&name)
	fmt.Print("What is your quest? ")
	fmt.Scanln(&quest)
	fmt.Print("What is your favorite color? ")
	fmt.Scanln(&color)

	//Off guard questions
	fmt.Print("What is the capital of Bahrain?")
	fmt.Scanln(&capital)

	   if(capital == "Manama") {
	   fmt.Print("Right. Off you go.")
	   } else {
	 	fmt.Print("Auuuuuuuuuuuuuuuuugh")  
	   }
	

	fmt.Print("What is the air-speed velocity of an unladen swallow? ")
	fmt.Scanln(&finalQuestion)

	
	if(finalQuestion == "What do you mean? An African or European swallow?") {
		fmt.Print("What? I don't know that! Auuuuuuuuuuuuuuuuugh");
	} else {
		fmt.Print("Auuuuuuuuuuugh");
	}
	
}
