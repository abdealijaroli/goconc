package main

import (
    "fmt"
    "sync"
)

var (
    mu sync.Mutex
    counter int
)

func worker(wg *sync.WaitGroup, id int) {
    defer wg.Done()

    mu.Lock() // acquire the lock
    defer mu.Unlock() // release the lock when the function returns

    fmt.Printf("Worker %d is starting work\n", id)
    counter++
    fmt.Printf("Worker %d has finished work. Counter: %d\n", id, counter)
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go worker(&wg, i)
    }

    wg.Wait()
}