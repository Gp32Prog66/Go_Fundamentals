package main

/*
Movie Script from Genius.com
https://genius.com/Monty-python-monty-python-and-the-holy-grail-bridge-of-death-annotated
*/

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	fmt.Println(input)


	//Testing String Import Library
	str1 := "Testing String";

	res1 := strings.ToLower(str1)

	fmt.Println(res1)

	//Declaring Variables
	//var name string

	name := "";
	nameStr := strings.ToLower(name)

	//var quest string

	quest := "";
	questStr := strings.ToLower(quest)


	//var color string
	color := "";
	colorStr := strings.ToLower(color)

	//Strings 
	

	//var capital string
	capital := "";
capitalStr := strings.ToLower(color)

	var finalQuestion string

	//Questions
	fmt.Print("What is your name? ")
	fmt.Scanln(&nameStr)

	fmt.Print("What is your quest? ")
	fmt.Scanln(&questStr)
	fmt.Print("What is your favorite color? ")
	fmt.Scanln(&colorStr)

	//Off guard questions
	fmt.Print("What is the capital of Bahrain?")
	fmt.Scanln(&capitalStr)

		//Determine Correct Answer for Bahrain Capital
	   if(capital == "Manama") {
	   fmt.Print("Right. Off you go.")
	   } else {
	 	fmt.Print("Auuuuuuuuuuuuuuuuugh")  
	   }
	

	//Final Questions
	fmt.Print("What is the air-speed velocity of an unladen swallow? ")
	fmt.Scanln(&finalQuestion)

	//Answer
	if(finalQuestion == "What do you mean? An African or European swallow?") {
		fmt.Print("What? I don't know that! Auuuuuuuuuuuuuuuuugh");
	} else {
		fmt.Print("Auuuuuuuuuuugh");
	}
	
}
