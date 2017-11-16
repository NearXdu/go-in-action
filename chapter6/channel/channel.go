package main

import (
	"log"
	"sync"
	"time"
)

var (
	ch chan int
	wg sync.WaitGroup
)

func main() {
	ch := make(chan int)
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second * time.Duration(1))
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for c := range ch {
			log.Println(c)
		}
	}()

	wg.Wait()

}
