package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	//feeds stream
	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err)
	}

	//result channel
	results := make(chan *Result)

	//countDown Latch
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	//every feed need a goroutine
	for _, feed := range feeds {
		//
		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	//monitor
	go func() {
		waitGroup.Wait()

		//close channel
		close(results)
	}()

	Display(results)
}

//Key - Value mapping  (key: feed type) (value: matcher)
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
