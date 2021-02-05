package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
)

/**
  URL checker and Page downloader

Requirements:
	* Pass a string of urls to a function that can save the response body in a file

Note: With Goroutines and WaitGroups
*/
type UrlStatus struct {
	err error
	url string
}

func main() {
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
	noOfGoRoutines := 0
	log.Printf("Start: No of Goroutines %d", runtime.NumGoroutine())
	ch := make(chan UrlStatus, len(urls))
	for _, v := range urls {
		go asyncHttpGet(v, ch)
		noOfGoRoutines++
	}
	log.Println("closing")
	for i := 0; i < noOfGoRoutines; i++ {
		log.Println("looping")
		result := <-ch
		if result.err != nil {
			log.Printf("Error while calling %s", result.url)
		}
		log.Printf("Got the response from %s", result.url)
	}

	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
	log.Printf("End: No of Goroutines %d", runtime.NumGoroutine())

}

func asyncHttpGet(url string, ch chan UrlStatus) {
	log.Println("calling")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err")
		fmt.Printf("%s is down", url)
		ch <- UrlStatus{
			err: err,
			url: url,
		}
	} else {
		defer resp.Body.Close()
		bytes, _ := ioutil.ReadAll(resp.Body)
		file := strings.Split(url, "//")[1]
		file += ".txt"
		err := ioutil.WriteFile(file, bytes, 0664)
		if err != nil {
			log.Fatal("Error while writing")
		}
		ch <- UrlStatus{
			err: nil,
			url: url,
		}
	}
}
