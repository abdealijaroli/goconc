package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump()
	go suck(stream)
	// the above 2 lines can be shortened to: go suck( pump() )
	time.Sleep(2 * 1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// func suck(ch chan int) {
// 	for {
// 		fmt.Println(<-ch)
// 	}
// }

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

// The pump function starts a goroutine that continuously sends integers to a channel. The suck function starts another goroutine that continuously reads from this channel and prints the integers. However, these goroutines are stopped when the main goroutine ends.