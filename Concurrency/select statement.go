// Do check the readme once
package main

import (
	"fmt"
	"time"
)

func main(){
	ch1 := make (chan int)
	ch2 := make (chan int)
	go first(ch1)
	go second(ch2)
	select {
	case val1 := ch1 :
		fmt.Prinln(val1)
	case val2 := ch2 : 
		fmt.Println(val2)
	}
	time.Sleep(1 * time.Second)
}

func first(ch1 chan int){
	ch1 <- 10
}

func second(ch2 chan int){
	ch2 <- 20
}