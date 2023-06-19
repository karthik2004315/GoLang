// Check the readme once
package main

import (
	"fmt"
	"time"
)

func printIt(i int){
	time.Sleep(1 * time.Second)
	fmt.Println(i)
}

func main(){
	start := time.Now()
	for i:=0; i<100000; i++{
		printIt(i)
	}
	since := time.Since(start)
}

// Now when we execute what happens is it will print numbers from 1 to 1000000 and they will come 1 by 1 with a sleep time of 1 second.
// when we calculate the time of execution it takes 1000000 seconds 