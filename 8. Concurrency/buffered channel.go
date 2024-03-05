// Do check the readme once
package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int, 3)
	go sent(ch, &wg)
	wg.Wait()
}
func sent(ch chan int, s *sync.WaitGroup){
	ch <- 1
	ch <- 2
	ch <- 3
	go recieved(ch, s)
	fmt.Println("sent data to the channel")
	s.Done()
}

func buy(ch chan int, s *sync.WaitGroup){
	fmt.Println("waiting for the data")
	fmt.Println("recieved data - ", <- ch)
	s.Done()
}