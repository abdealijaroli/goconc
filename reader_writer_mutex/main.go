package main

import (
    "fmt"
    "sync"
)

var (
    rwmu sync.RWMutex
    counter int
)

func reader(wg *sync.WaitGroup, id int) {
    defer wg.Done()

    rwmu.RLock() // acquire the read lock
    defer rwmu.RUnlock() // release the read lock when the function returns

    fmt.Printf("Reader %d: counter = %d\n", id, counter)
}

func writer(wg *sync.WaitGroup, id int) {
    defer wg.Done()

    rwmu.Lock() // acquire the write lock
    defer rwmu.Unlock() // release the write lock when the function returns

    counter++
    fmt.Printf("Writer %d: counter = %d\n", id, counter)
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go reader(&wg, i)
    }

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go writer(&wg, i)
    }

    wg.Wait()
}