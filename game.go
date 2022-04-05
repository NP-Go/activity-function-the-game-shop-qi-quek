package main

import "fmt"

//Declare Struct for Game

type gameStruct struct { //defining type gameStruct with keyword -struct
	title string
	price float64
}

//Declare method for Game

//same name method but different of different types
func (game gameStruct) print() {
	//this is for list
	fmt.Println(game.title, game.price) //prints the attributes of the struct this .print is the .print for line 12
}
