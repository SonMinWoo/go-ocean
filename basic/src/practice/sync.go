package main

import (
	"fmt"
	// "time"
)

func runSomeLoop(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Printing:", i)
	}
}

func main() {
	go runSomeLoop(10)

	// time.Sleep(2*time.Second)
	fmt.Println("Hello, playground")
}