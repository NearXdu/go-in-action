package main

import (
	"log"
	"sync"
	"time"
	"zxcode/go-in-action/chapter7/work"
)

var names = []string{
	"zhangxiao",
	"weileqi",
	"lihuan",
	"liuhuan",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	//2 goroutine
	p := work.New(2)

	var wg sync.WaitGroup

	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()

}
