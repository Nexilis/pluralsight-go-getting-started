package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// let's work on 8 processors max
	runtime.GOMAXPROCS(8)

	// creating printing channel
	ch := make(chan string)
	// creating shutdown channel
	shutdownCh := make(chan bool)

	go printEnglishAlphabetWithChannel(ch)
	go channelPrinter(ch, shutdownCh)

	// applications waits for a message comming out from the channel
	<-shutdownCh
}

func printEnglishAlphabetWithChannel(ch chan string) {
	// channels - sharing data between threads
	// byte('a') == explicit cast to byte
	for l := byte('a'); l <= byte('z'); l++ {
		// channel, receive, operator, message
		ch <- string(l)
	}

	// we can't receive more messages anymore
	close(ch)
}

func channelPrinter(ch chan string, shutdownCh chan bool) {
	for l := range ch {
		fmt.Println(l)
	}

	shutdownCh <-true
}


func concurrencyChaosGoroutines() {
	// goroutines - thread like constructs
	// (but not threads, not so problematic for managing memory)
	// for asynchronous tasks
	
	// concurrency is not parallelism! ;)

	go printEnglishAlphabet()
	fmt.Println("This comes first!")

	// we need to sleep so other goroutines have time to execute
	time.Sleep(100 * time.Millisecond)
}

func printEnglishAlphabet() {
	// byte('a') == explicit cast to byte
	for l := byte('a'); l <= byte('z'); l++ {
		// and here we create 24 goroutines
		// goroutines are extremely light and inexpensive to start
		go fmt.Println(string(l))
	}
}
