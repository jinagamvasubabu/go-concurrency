package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
)

/**
  URL checker and Page downloader

Requirements:
	* Pass a string of urls to a function that can save the response body in a file

Note: With Goroutines and WaitGroups
*/

type urlStatus struct {
	err error
	url string
}

func main() {

	cpuProfile, _ := os.Create("cpuprofile")
	memProfile, _ := os.Create("memprofile")
	pprof.StartCPUProfile(cpuProfile)
	urls := []string{
		"http://www.google.com",
		"https://www.golang.com",
		"https://www.wikipedia.com",
		"http://www.facebook.com",
		"https://www.uber.com",
		"http://www.ola.com",
		"https://www.swiggy.com",
		"https://www.zomato.com",
	}
	start := time.Now()
	log.Printf("Start: No of Goroutines %d", runtime.NumGoroutine())
	input := make(chan string, len(urls))
	for _, v := range urls {
		input <- v
	}
	//close(input)
	output := make(chan urlStatus, len(urls))
	for range urls {
		go asyncHttp(input, output)
	}
	log.Println("closing")
	for i := 0; i < len(urls); i++ {
		log.Println("looping")
		result := <-output
		if result.err != nil {
			log.Printf("Error while calling %s", result.url)
		}
		log.Printf("Got the response from %s", result.url)
	}

	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
	log.Printf("End: No of Goroutines %d", runtime.NumGoroutine())
	pprof.StopCPUProfile()
	_ = pprof.WriteHeapProfile(memProfile)
	http.ListenAndServe(":8080", nil)

}

func asyncHttp(input <-chan string, output chan<- urlStatus) {
	for v := range input {
		resp, err := http.Get(v)
		if err != nil {
			fmt.Printf("err")
			fmt.Printf("%s is down", v)
			output <- urlStatus{
				err: err,
				url: v,
			}
		} else {
			defer resp.Body.Close()
			bytes, _ := ioutil.ReadAll(resp.Body)
			file := strings.Split(v, "//")[1]
			file += ".txt"
			err := ioutil.WriteFile(file, bytes, 0664)
			if err != nil {
				log.Fatal("Error while writing")
			}
			output <- urlStatus{
				err: err,
				url: v,
			}
		}
	}

}
