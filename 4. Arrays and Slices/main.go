// Do check the notes once
package main

import "fmt"

func main(){
	fruits := [...]string{ "apple", "banana", "grapes"}

	for _, value := range fruits{
		fmt.Println(value)
	}

	slice1 = fruits[:1]
	slice2 = fruits[1:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	slice3 := append(slice1, slice2...)
	fmt.Println(slice3)

	mappy := make(map[int]string)
	mappy[1] = "first"
	mappy[2] = "second"
	mappy[3] = "third"
	fmt.Println(mappy)
}