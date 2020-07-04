package main

import (
	"fmt"
	"sync"
)

/**
  Race conditions:
   * In this example i want to solve  Race condition using mutex
   * Mutex allows two or more goroutines share the same memory (variable here) but the catch is not simultaneously
   * Mutex is under go sync package, has lock and unlock methods which helps you do the above



  `go run -race main_data_race_mutex_example_6.go` - run this again to see if there are any race conditions
  `go run main_data_race_mutex_example_6.go` - always value is zero

*/
func main() {
	var n int32
	var wg sync.WaitGroup
	var mx sync.Mutex
	wg.Add(200)
	for i := 0; i < 100; i++ {
		go func() { //anonymous function aka go routine, his job is to increment 100 times
			mx.Lock()
			//defer mx.Unlock() // which does the same writing after increment operation
			n++
			mx.Unlock()
			//time.Sleep(1 * time.Second) // to give the real time delay, just to give the feel it does complex increment :)
			wg.Done()
		}()
		go func() { //anonymous function aka go routine, his job is to increment 100 times
			mx.Lock()
			n--
			mx.Unlock()
			//time.Sleep(1 * time.Second) // to give the real time delay, just to give the feel it does complex increment :)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Final value %d", n)
}
