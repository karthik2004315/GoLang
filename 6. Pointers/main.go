package main

import "fmt"

// Function to update the value of an integer by reference
func update(i *int) {
	*i += 20
}

// Function to modify the first element of a slice
func modify(str []int) {
	str[0] = 10 // No need for pointers since slices are reference types
}

// Function to modify a map
func modifyMaps(m map[string]int) {
	m["third"] = 3 // No need for pointers since maps are reference types
}

func main() {
	i := 10
	var ptr1 *int = &i
	fmt.Println("value is:", *ptr1, "address is:", ptr1)

	var ptr2 = &i
	fmt.Println("value is:", *ptr2, "address is:", ptr2)

	ptr3 := &i
	fmt.Println("value is:", *ptr3, "address is:", ptr3)

	// Passing by reference
	update(&i)
	fmt.Println(i)

	// Slices
	slice := []int{1, 2, 3}
	modify(slice)
	fmt.Println(slice)

	// Maps
	mappy := map[string]int{
		"first":  1,
		"second": 2,
	}

	modifyMaps(mappy)
	fmt.Println(mappy)
}
