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
		}
		lastStatus = true
	case InternetDisconnected:
		if lastStatus {
			// fmt.Printf("InternetDisconnected...starting timer now\n")
			startTime = time.Now()
		}
		lastStatus = false
	}
}
