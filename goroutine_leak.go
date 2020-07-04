package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
	//"sync"
)

func main() {
	fmt.Printf("Starting:%d", runtime.NumGoroutine())
	res, err := getNumbers()
	fmt.Println("\n starting")
	fmt.Printf("%d", res)
	fmt.Println()
	fmt.Printf("%s", err)
	fmt.Println()
	fmt.Printf("Ending:%d", runtime.NumGoroutine())
}

func getNumbers() (int32, error) {
	var evenNumber int32
	var oddNumber int32
	//var wg sync.WaitGroup
	routinesCount := 2
	errCh := make(chan error, routinesCount)
	//wg.Add(2)
	go func() {
		resp, err := GetEvenNumber()
		if err == nil {
			evenNumber = resp
		}
		errCh <- err
		//wg.Done()
	}()
	go func() {
		resp, err := GetOddNumber()
		if err == nil {
			oddNumber = resp
		}
		errCh <- err
		//wg.Done()
	}()
	//wg.Wait()
	for i := 0; i < 3; i++ {
		err := <-errCh
		if err != nil {
			return 0, err
		}
	}
	return 1, nil // 1 means success and 0 means error
}

func GetEvenNumber() (int32, error) {
	fmt.Printf("\n In Even Func")
	time.Sleep(5 * time.Second)
	return 2, errors.New("error in even function")
}

func GetOddNumber() (int32, error) {
	fmt.Printf("\n In Odd Func")
	time.Sleep(5 * time.Second)
	return 1, nil
}
