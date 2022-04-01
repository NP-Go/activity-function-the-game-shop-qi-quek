package main

//declare List

type list []*gameStruct

//declare method for List Print

func (l list) print() {
	//method for type list
	for _, item := range l {
		item.print()
	}
}
