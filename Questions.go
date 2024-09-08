package main

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
	fmt.Print("What is the capital of Bahrain ")
	fmt.Scanln(&capital)

	/*
	   if(capital = "Manama")
	   {
	   fmt.Print("Right. Off you go.")
	   }
	*/

	fmt.Print("What is the air-speed velocity of an unladen swallow? ")
	fmt.Scanln(&finalQuestion)

}
