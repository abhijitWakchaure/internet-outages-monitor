package main

import (
	"fmt"
	"log"
	"time"

	"github.com/abhijitWakchaure/internet-outages-monitor/storage"
)

// Event ...
type Event string

type internetStatus bool

// Constants for internet connection events
const (
	InternetConnected          Event          = "InternetConnected"
	InternetDisconnected       Event          = "InternetDisconnected"
	StatusInternetConnected    internetStatus = true
	StatusInternetDisconnected internetStatus = false
)

// assume internet was available
var lastStatus internetStatus = StatusInternetConnected

func recordEvent(event Event) {
	switch event {
	case InternetConnected:
		if lastStatus == StatusInternetConnected {
			break
		}
		lastStatus = StatusInternetConnected
		startTime := storage.ReleaseLock()
		// notify the user that internet is back
		msg := fmt.Sprintf("We are back! Internet was out for %s", time.Now().Sub(startTime))
		log.Printf("%s\n", msg)
		notif.Notify(msg)
		// internet is back, restore the original tick interval
		interval = tickInterval
		log.Printf("Internet is back, restoring tick interval to %s\n", interval)
	case InternetDisconnected:
		if lastStatus == StatusInternetDisconnected {
			break
		}
		lastStatus = StatusInternetDisconnected
		// internet is disconnected, let's check status more frequently
		interval = time.Second * 3
		log.Printf("Internet disconnected, setting tick interval to %s\n", interval)
		storage.AquireLock()
	}
}
