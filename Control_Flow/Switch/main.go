package main

import "fmt"

func main() {
	switch "Powers"{
	case "MoneyPenny":
		fmt.Println("I'm Miss Money Penny")
	case "Bond":
		fmt.Println("The name is Bond, James Bond")
	case "Austin":
		fmt.Println("Austin, Austin Powers Baby yea.")
		//fallthrough // This will allow the next case statement to evaluate.
					// Generally DO NOT USE FALLTHROUGH!
	case "Powers":
		fmt.Println("Powers, Austin Powers, Do I make you horny baby? Do I?")
	default:
		fmt.Println("Default Case")
	}
}
