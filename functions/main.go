// Do check the notes once
package main

import "fmt"

func main(){
// normal way

	func add(a int, b int)(int, int){
		c := a*b
		d := a/b
		return c, d
	}

	fmt.Println(add(20, 30))

// using variadic function

	func factorial(num ...int)int{
		mul := 1
		for _,val :=range num{
			mul *= val
		}
		returm mul
	}
	
	fmt.Println(factorial(5,4,3,2,1))

// anonymoys function

	x := func(num int) int{
		return x
	}(6)

	fmt.Println("just normally printing value: ", x)

// defer 

	func greet1(str string) {
		fmt.Println("Hello ", str)
	}

	func greet2(str string) {
		fmt.Println("Hola ", str)
	}

	greet1("Go")
	defer greet2("friends")
}