package main

import (
	"fmt"
	"time"
)

func runLoopSend(n int, ch chan int) {
	for i:=0; i<n; i++ {
		ch <- i
	}
	close(ch)
}

func runLoopReceive(ch chan int) {
	for {
		i, ok := <- ch // ok (channel open : true)
		if !ok {
			break
		}
		fmt.Println("Received val:", i, ok)
	}
}

func selectReceive(ch1 chan int, ch2 chan int) {
	select {
	case i:= <-ch1:
		fmt.Println("Received val:", i)
	case ch2 <- 12:
		fmt.Println("Sent value 12 to ch2")
	case <- time.After(1*time.Second):
		fmt.Println("timed out")
	default :
		fmt.Println("No channel ready")
	}
}

func main() {
	myChannel := make(chan int)
	selectChannel := make(chan int)
	//buffer 
	//myBufferedChannel := make(chan int, 10) // save 10 int values
	go runLoopSend(10, myChannel)
	go selectReceive(myChannel, selectChannel)
	go runLoopReceive(selectChannel)
	time.Sleep(2*time.Second)
}