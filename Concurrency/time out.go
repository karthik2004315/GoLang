// Do check the readme once
package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make( chan int)
	go sent(ch)

	select{
	case val := <-ch:
		fmt.Println("value is: ", val)
		
	case <- time.After(1* time.Second):
		fmt.Println("timed out")
	}
}

func sent(ch chan int){
	time.Sleep(3 * time.Second)
	ch <- 10
}

// now when executed we get result as "timed out" this is because go routine is waiting for 3 seconds.
// where as go routine is timed out after 1 second so we get result as "timed out"