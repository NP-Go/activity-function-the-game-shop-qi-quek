package main

import "fmt"

func main() {
	//Create objects

	var (
		minecraft       = gameStruct{title: "minecraft", price: 5}
		worldOfWarcraft = gameStruct{title: "World of warcraft", price: 19}
		eliteDangerous  = gameStruct{title: "Elite Dangerous", price: 54}
	)

	var items []*gameStruct
	items = append(items, &minecraft, &worldOfWarcraft, &eliteDangerous)
	//append items in the address to the list
	fmt.Println(items)
	//print the list
	fmt.Println(*items[0])
	//print item of index 0 from list
	my := list(items)

	my.print()

}
