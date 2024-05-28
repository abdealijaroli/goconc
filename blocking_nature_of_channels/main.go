package main
import "fmt"
import "time"

func main() {
	c := make(chan int)
	go func() {  
		time.Sleep(5 * 1e9)
		fmt.Println("received", <- c) // value received from channel
	}() 
	fmt.Println("sending", 10)
	c <- 10 // putting 10 on the channel
	fmt.Println("sent", 10) 
}