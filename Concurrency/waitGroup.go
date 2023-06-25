// Do check the Readme once
package main

import (
	"fmt"
	"time"
	"sync"
)

func mul(i int, wg *sync.WaitGroup){
	fmt.Println(i * i)
	wg.Done()
}
func main(){
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0; i<10; i++{
		mul(i, &wg)
	}
	wg.Wait()
}