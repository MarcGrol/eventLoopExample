package main

import (
	"log"
	"time"
)

func performFetchWithResultOnChannel(processChannel chan ProcessSlice) {
	defer leave(enter("performFetchWithResultOnChannel"))

	psList, err := fetchCurrentProcesses()
	if err == nil {
		processChannel <- psList
	}
}

func processFetcherEventLoop(intervalSecs int, commandChannel chan bool, returnChannel chan ProcessSlice) {
	defer leave(enter("processFetcherEventLoop"))

	terminating := false
	fetchResultChannel := make(chan ProcessSlice)

	fetchInProgress := false
	initial := time.After(0) // immediately
	for {
		repetitiveTimer := time.After(time.Duration(intervalSecs) * time.Second)
		select {

		case <-initial:
			log.Printf("Got initial timer command")
			commandChannel <- true

		case <-repetitiveTimer:
			log.Printf("Got repetitive timer command")
			commandChannel <- true

		case command := <-commandChannel:
			if command == false {
				if fetchInProgress == false {
					log.Printf("Terminating fetcher-loop immediately")
					returnChannel <- ProcessSlice{} // empty procesSlice and terminate goroutine
					return
				} else {
					log.Printf("Waitfor fetch-in-progress before terminating fetcher-loop")
					terminating = true
				}
			} else if fetchInProgress {
				log.Printf("Ignore command because fetch-in-progress")
			} else {
				fetchInProgress = true
				go performFetchWithResultOnChannel(fetchResultChannel)
			}

		case psList := <-fetchResultChannel:
			log.Printf("Got fetch result %+v", psList)
			fetchInProgress = false
			returnChannel <- psList
			if terminating {
				log.Printf("Terminating fetcher-loop")
				return
			}
		}
	}
}

func startProcessFetcherLoop(intervalSecs int, inputChannel chan bool) chan ProcessSlice {
	defer leave(enter("runProcessFetcher"))

	processChannel := make(chan ProcessSlice)

	// kick off fetcher in seperate go-routine
	go processFetcherEventLoop(intervalSecs, inputChannel, processChannel)

	return processChannel
}
