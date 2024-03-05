// Do check the notes once
package main

import "fmt"

func main(){
	i := 10
	var ptr1 *int = &i
	fmt.Println("value is: ",*ptr1, "address is: ", ptr1)

	var ptr2 = &i
	fmt.Println("value is: ",*ptr2, "address is: ", ptr2)

	ptr3 := &i
	fmt.Println("value is: ",*ptr3, "address is: ", ptr3)\

// Passing by reference

	func update(i *int){
		*i += 20
	}

	update(&i)
	fmt.Println(i)

// slices

	func modify(str []int){
		str[0] = 10           // see we dont need the concept of pointers
	}

	slice = []int{1, 2, 3}
	modify(slice)
	fmt.Println(slice)

// maps 

	func modifyMaps(m map[string]int){
		m["third"] = 3        // see we didnt use the concept of pointers for maps
	}

	mappy := map[string]int
	mappy["first"]  = 1
	mappy["second"] = 2

	modifyMaps(mappy)
	fmt.Println(mappy)
}