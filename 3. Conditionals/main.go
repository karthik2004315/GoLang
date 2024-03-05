// Do check the notes once before you see the program

package main

import "fmt"

func main(){
	i := 4
	for i>=0 {
		fmt.Println(i)
		i--
	}

	name := "Go"
	if name == "Go" {
		fmt.Println(name, " You are in serious trouble")
	}
	else {
		fmt.Println("Ok you can leave")
	}
}