package main

import ("fmt",
		"strconv"
)

func main(){
	var i int = 10
	var f float64 = float64(i)
	fmt.Println("integer to float: ", f) // same as float64 wrapped around variable i we can wrap int around float value to conver into int

    var str string = Itoa(i)
	fmt.Println("we converted integer to string: ", str) // we used "Itoa" method to convert integer to string

	// we can use "Atoi" method to convert string to integer

}