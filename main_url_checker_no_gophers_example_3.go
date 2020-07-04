package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

/**
  URL checker and Page downloader

Requirements:
	* Pass a string of urls to a function that can save the response body in a file

Note: With No Goroutines
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
	for _, v := range urls {
		CheckAndSaveBody_No_Gophers(v)
		strings.Repeat("######", 20)
	}
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

func CheckAndSaveBody_No_Gophers(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err")
		fmt.Printf("%s is down", url)
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
	}
}
