package main

//declare List

type list []*gameStruct

// a slice of pointers to be returned
//in this case, it is a slice of the struct pointers

func (l list) print() { //method for type list
	for _, item := range l { //prints item for each element within the slice
		item.print()
	}
}
