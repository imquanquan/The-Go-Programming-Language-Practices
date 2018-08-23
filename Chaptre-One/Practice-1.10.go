package main

import (
	"time"
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Fetch(url string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main()  {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go Fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<- ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}