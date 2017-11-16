package main

import (
	"log"
	"os"
	"time"
	"zxcode/go-in-action/chapter7/runner"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting Work.")

	r := runner.New(timeout)
	r.Start()
	r.Add(createTask(),
		createTask(),
		createTask(),
		createTask())
	//r.Add(createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)

		}
	}
	log.Println("Process ended")

}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
