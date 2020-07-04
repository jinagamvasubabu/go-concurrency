package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

/**
  URL checker and Page downloader

Requirements:
	* Pass a string of urls to a function that can save the response body in a file

Note: With Goroutines and WaitGroups
*/
func main() {
	urls := []string{
		"http://www.google.com",
		"https://www.golang.com",
		"https://www.wikipedia.com",
		"http://www.facebook.com",
		"https://www.sixt.com",
		"https://www.uber.com",
		"http://www.ola.com",
		"https://www.swiggy.com",
		"https://www.zomato.com",
	}
	start := time.Now()
	var wg sync.WaitGroup
	for _, v := range urls {
		go CheckAndSaveBody_Gophers(v, &wg)
		fmt.Printf("%s \n", strings.Repeat("######", 20))
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

func CheckAndSaveBody_Gophers(url string, wg *sync.WaitGroup) {
	wg.Add(1)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err")
		fmt.Printf("%s is down", url)
		wg.Done()
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			bytes, _ := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"
			err := ioutil.WriteFile(file, bytes, 0664)
			if err != nil {
				log.Fatal("Error while writing")
			}
		}
		wg.Done()
	}
}
