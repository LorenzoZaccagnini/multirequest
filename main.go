package main

import (
	"multirequest/args"
	"multirequest/client"
	"multirequest/logger"
	"multirequest/openfile"
	"sync"
)

func main() {

	//logger setup
	fileLogger := logger.InitLogger()

	fileLogger.Info().Timestamp().Msg("starting program")

	//args
	file_path := args.InitArgs(fileLogger)

	//scan file
	urls := openfile.Init(file_path, fileLogger)

	var wg sync.WaitGroup
	wg.Add(len(urls))
	// loop through the urls and log the status code
	// should we limit the number of goroutines? https://dave.cheney.net/2013/06/02/why-is-a-goroutines-stack-infinite
	for _, u := range urls {
		go func(url string) {
			client.CheckRequest(url, fileLogger, &wg)
		}(u)
	}

	wg.Wait()

}
