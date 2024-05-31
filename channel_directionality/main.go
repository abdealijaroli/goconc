package main

func main() {
	var c = make(chan int) // bidirectional
	go source(c)
	go sink(c)
}

// var send_only chan<- int // data can only be sent (written) to the channel
// var recv_only <-chan int // data can only be received (read) from the channel

func source(ch chan<- int) { // send_only channel
	for {
		ch <- 1 // sending data to ch channel
	} 
}
func sink(ch <-chan int) {  // recv_only channel
	for {
		<-ch  // receiving data from ch channel
	} 
}



