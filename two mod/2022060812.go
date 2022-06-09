package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	xx     = 0
	wg4    sync.WaitGroup
	rwlock sync.RWMutex
)

func read() {
	defer wg4.Done()
	rwlock.RLock()
	fmt.Println(xx)
	time.Sleep(time.Millisecond)
	rwlock.Unlock()
}

func write() {
	defer wg4.Done()
	rwlock.Lock()
	xx += 1
	time.Sleep(time.Millisecond * 5)
	rwlock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg4.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg4.Add(1)
	}
	wg4.Wait()
	fmt.Println(time.Since(start))
}
