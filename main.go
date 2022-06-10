package main

import "fmt"
// card := "Ace of spades" // error as initialization is outside function body

// var card string 
// We can declare variable outside but can't initialise it


func main() {
	// var card string = "Ace of spades" // 1st way of declaration

	// card := "Ace of spades" // 2nd way of declaration

	var card string // We can first only declare variable and then later initialize it
	card = "Ace of spades"
	

	fmt.Println(card)
}
