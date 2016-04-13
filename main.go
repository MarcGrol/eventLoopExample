package main

import (
	"log"
	"time"
)

func mainEventLoop() {
	defer leave(enter("mainEventLoop"))
	terminating := false

	networkIoEvent := time.Tick(time.Duration(1) * time.Second)                         // simulate io-event every 5 seconds
	programTermination := time.After(time.Duration(60) * time.Second)                   // quit after 100 secs
	processFetchCommandChannel := make(chan bool, 20)                                   // buffered channel: trigger fetch-processes on demand
	processListResultChannel := startProcessFetcherLoop(10, processFetchCommandChannel) // result of fetch process

	for {
		select {

		case <-networkIoEvent:
			log.Printf("fire fetch command")
			processFetchCommandChannel <- true // buffered channel because this event-loop must not block

		case processList := <-processListResultChannel:
			if terminating == true {
				log.Printf("Termination complete")
				return
			}
			// TODO do something smart with the most up to date process-list
			log.Printf("processes:%+v", processList)

		case <-programTermination:
			log.Println("Start termination")
			terminating = true
			processFetchCommandChannel <- false

		}
	}

}

func main() {
	mainEventLoop()
}
