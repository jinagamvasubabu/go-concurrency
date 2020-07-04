package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  Race conditions:
   * In this example i want to replicate a Race condition where two or more go routines
   * access the same variable N one adds and one decrements, ideally when all the go routines
   * finishes it jobs, the N value should be same as at the time of initialization, but let's see what happens

   You will get some unexpected values every time you run :(

   How do you detect race condition: provide -race flag while running main.go

  `go run -race main_data_race_example_5.go`

  This will result two race condition identified and in the next program we are going to solve this with Mutex lock and unlock
*/
func main() {
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go func() { //anonymous function aka go routine, his job is to increment 100 times
			wg.Add(1)
			n++
			time.Sleep(1 * time.Second) // to give the real time delay, just to give the feel it does complex increment :)
			wg.Done()
		}()
		go func() { //anonymous function aka go routine, his job is to increment 100 times
			wg.Add(1)
			n--
			time.Sleep(1 * time.Second) // to give the real time delay, just to give the feel it does complex increment :)
			wg.Done()
		}()
	}
	fmt.Printf("Final value %d", n)
	wg.Wait()
}
