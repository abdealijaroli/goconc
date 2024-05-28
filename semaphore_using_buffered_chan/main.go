package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    sem := make(chan struct{}, 2) // semaphore with a capacity of 2

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            sem <- struct{}{} // acquire: send into the semaphore channel
            defer func() { <-sem }() // release: receive from the semaphore channel

            fmt.Println(i, "is starting work")
            time.Sleep(2 * time.Second) // simulate work
            fmt.Println(i, "has finished work")
        }(i)
    }

    wg.Wait()
}