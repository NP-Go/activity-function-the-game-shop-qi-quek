package main

import "fmt"

//Declare Struct for Game

type gameStruct struct {
	//structure define
	title string
	price float64
}

//Declare method for Game

//declare method for same type but

func (game gameStruct) print() { //same name method but different of different types
	//this is for list
	fmt.Println(game.title, game.price)
}
