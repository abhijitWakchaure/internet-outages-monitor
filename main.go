package main

import (
	"log"
	"time"

	"github.com/abhijitWakchaure/internet-outages-monitor/env"
	"github.com/abhijitWakchaure/internet-outages-monitor/notifier"
)

var (
	notif             Notifier
	internetAvailable = false
)

func main() {
	notif = &notifier.Slack{}
	err := notif.Register()
	if err != nil {
		panic(err)
	}
	// err = notifier.Notify("Hello")
	// if err != nil {
	// 	panic(err)
	// }
	interval, err := time.ParseDuration(env.Read(env.ENVTICKINTERVAL))
	if err != nil {
		panic(err)
	}

	for range time.Tick(interval) {
		internetAvailable, err = checkInternetStatus()
		if err != nil {
			log.Printf("failed to check internet status due to %v\n", err)
			panic(err)
		}
		// log.Printf("internetAvailable: %v\n", internetAvailable)
		if !internetAvailable {
			recordEvent(InternetDisconnected)
		} else {
			recordEvent(InternetConnected)
		}
	}

}
