// Do check the notes once

package main

import "fmt"

type Fruits struct{
	color string
	price int
}

func main(){
// first way of declaration and intiliaztion

	st1 := new(Fruits)
	fmt.Println(st1)

// second way of declaration and initialization

	st2 := Fruits("red", 100) // we are assigning values in the way the the data types are declared
	fmt.Println(st2)

// Accessing and changing values of structure

	st2.color = "orange"
	st2.price = 50

// Methods 

	func (f *Fruits) demo(){
		f.color = "green"
	}

	f = Fruits{price: 20}
	demo(&f)
	fmt.Println(f)
}