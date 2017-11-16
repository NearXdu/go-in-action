package main

import (
	"log"
	"os"
	_ "zxcode/go-in-action/rss/matchers"
	"zxcode/go-in-action/rss/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
