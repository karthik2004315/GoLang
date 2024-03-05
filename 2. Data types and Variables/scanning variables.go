// Do check the notes once

package main

import "fmt"

func main(){
	var name string
	fmt.Print("Enter a name: ")
	fmt.Scanf("%s", &name)

	// It returns two values count and error for above example

	var num int
	fmt.printf("Enter a number: ")
	count, err := fmt.Scanf("%d", &num)
	fmt.Println(count) // this prints number of successful values taken input
	fmt.Println(err)   // this prints an error if we have assigned wrong values to wrong data types
	
}