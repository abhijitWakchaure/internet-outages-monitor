package main

import (
	"fmt"
	"log"
	"time"
)

// Event ...
type Event string

// Constants for internet connection events
const (
	InternetConnected    Event = "InternetConnected"
	InternetDisconnected Event = "InternetDisconnected"
)

// assume internet was available
var lastStatus = true
var startTime time.Time

func recordEvent(event Event) {
	switch event {
	case InternetConnected:
		if !lastStatus {
			// notify the user that internet is back
			msg := fmt.Sprintf("We are back! Internet was out for %s", time.Now().Sub(startTime))
			log.Printf("%s\n", msg)
			notif.Notify(msg)
			// internet is back, restore the original tick interval
			interval = tickInterval
			log.Printf("Internet is back, restoring tick interval to %s\n", interval)
		}
		lastStatus = true
	case InternetDisconnected:
		if lastStatus {
			startTime = time.Now()
			// internet is disconnected, let's check status more frequently
			interval = time.Second * 3
			log.Printf("Internet disconnected, setting tick interval to %s\n", interval)
		}
		lastStatus = false
	}
}
