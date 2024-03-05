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
		go printIt(i)			// small change is we added go keyword before the function so that each will execute in seperate go routine
	}
	time.Sleep(2 * time.Second) // we added this because main function doesnt wait for all the go routines to finish so in order to get the output we will put a time out
	since := time.Since(start)
}
