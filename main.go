package main

import (
	"multirequest/args"
	"multirequest/client"
	"multirequest/logger"
	"multirequest/openfile"
	"os"
	"sync"
)

func main() {

	//logger setup
	fileLogger, err := logger.InitLogger()

	if err != nil {
		//print error and exit
		println(err.Error())
		os.Exit(1)
	}

	fileLogger.Info().Timestamp().Msg("starting program")

	//args
	filePath := args.InitArgs(fileLogger)

	//scan file
	urls, err := openfile.Init(filePath, fileLogger)
	if err != nil {
		//print error and exit
		println(err.Error())
		os.Exit(1)
	}

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
