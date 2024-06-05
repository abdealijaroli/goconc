package main

import "fmt"

func generate(ch chan int) {
	for i := 2; i < 100; i++ {
		ch <- i
	}
}

func filter(in chan int, out chan int, prime int) {
	for {
		i := <-in

		if i%prime != 0 {
			out <- i
		}
	}
}

// The out channel will have all numbers that are not multiples of the specific prime value used in the filter function. These numbers are not necessarily prime numbers themselves, they are just not divisible by the specific prime value.


func main() {
	ch := make(chan int)

	go generate(ch)

	for {
		prime := <-ch
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1 					// (1)
		fmt.Println(prime)
	}
}

// In the context of the Sieve of Eratosthenes algorithm, each filter function is responsible for filtering out multiples of a specific prime number.

// (1): The out channel of one filter function feeds into the in channel of the next filter function. This creates a pipeline of filters, and the numbers that make it through all the filters (i.e., are not divisible by any of the primes) are the prime numbers.




// BETTER APPROACH WITH FUNCTIONS - sieve(), generate(), filter() --------------------------------------------


package main

import "fmt"

func generate() chan int {
	ch := make(chan int)

	go func() {
		for i := 2; i < 100; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)

	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()

	return out
}

func sieve() chan int {
	res := make(chan int)

	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			res <- prime
		}
	}()
	return res
}

func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}