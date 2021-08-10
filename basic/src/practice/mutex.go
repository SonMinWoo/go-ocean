package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	myChannel := make(chan int)
	// counter +2
	wg.Add(2)
	go runLoopSend(10, myChannel)
	go runLoopReceive(myChannel)
	// wait wait counter to 0
	wg.Wait()
}

func runLoopSend(n int, ch chan int) {
	defer wg.Done() // wait group counter minus 1
	for i:= 0; i<n; i++ {
		ch <- i
	}
	close(ch)
}

func runLoopReceive(ch chan int) {
	defer wg.Done() // wait group counter minus 1
	for {
		i, ok := <- ch
		if !ok {
			break
		}
		fmt.Println("Received value:", i)
	}
}