package main

func main() {
	type Empty interface{}
	var empty Empty

	N := 10000

	data := make([]float64, N)
	res := make([]float64, N)
	sem := make(chan Empty, N) // semaphore

	for i, xi := range data {
		go func(i int, xi float64) {
			res[i] = doSomething(i, xi)
			sem <- empty
		}(i, xi)
	}

	// wait for goroutines to finish
	for i := 0; i < N; i++ {
		<-sem
	}

	//12: the current i and xi are passed to the closure as parameters, masking the i and xi variables from the outer for-loop. This allows each goroutine to have its own copy of i and xi; otherwise, the next iteration of the for-loop would update i and xi in all goroutines. On the other hand, the res slice is not passed to the closure, since each goroutine does not need a separate copy of it. The res slice is part of the closure’s environment but is not a parameter.

	// each iteration in the for-loop is done in parallel:
	for i, v := range data {
		go func(i int, v float64) {
			doSomething(i, v)
		}(i, v)
	}
}