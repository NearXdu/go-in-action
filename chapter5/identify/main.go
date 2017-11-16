package main

import (
	"fmt"
	"zxcode/go-in-action/chapter5/identify/counter"
)

func main() {
	cnt := counter.New(45)
	fmt.Println(cnt)
}
