package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status string
}

var errRequestFailed error = errors.New("Request Failed")

func main() {
	results := make(map[string]string)
	urls := []string{
		"https://github.com/",
		"https://nomadcoders.co/",
		"https://www.google.com/",
		"https://www.naver.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.amazon.com/",
	}

	c := make(chan requestResult)
	for _, url := range urls {
		go hitUrl(url, c)
	}

	for i:=0; i<len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitUrl(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}