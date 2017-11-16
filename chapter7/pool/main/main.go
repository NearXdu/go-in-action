package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"zxcode/go-in-action/chapter7/pool"
)

const (
	maxGroutines    = 25
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection ", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGroutines)

	p, err := pool.New(createConnection, pooledResources)

	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	wg.Wait()
	log.Println("Shutdown Program.")
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}