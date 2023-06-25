// Do check the readme once
package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	ch := make (chan string)
	go sent(ch)
	go recieved(ch)
}

// To send the data to a channel
func sent(ch chan string){
	ch <- "This is channel"
	fmt.Println("Sent data to the channel")
}

// To recieve data from a channel
func recieved(ch chan string){
	fmt.Println("waiting for data")
	val := <- ch
	fmt.Println("recieved data - ", val)
}