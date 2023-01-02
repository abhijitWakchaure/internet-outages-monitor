package main

import (
	_ "embed"
	"log"
	"time"

	"github.com/abhijitWakchaure/internet-outages-monitor/env"
	"github.com/abhijitWakchaure/internet-outages-monitor/notifier"
	"github.com/abhijitWakchaure/internet-outages-monitor/storage"
)

var (
	notif                  Notifier
	internetAvailable      = false
	interval, tickInterval time.Duration
	nsDomain               string
	nsPort                 string
	//go:embed VERSION
	version string
)

func main() {
	log.Printf("Starting Internet Outages Monitor [v%s]\n", version)
	notif = &notifier.Slack{}
	err := notif.Register()
	if err != nil {
		log.Printf("Error registering notifier: %s. I'll not be able to send notifications\n", err)
	}
	storage.Init()
	tickInterval, err = time.ParseDuration(env.Read(env.ENVTICKINTERVAL))
	if err != nil {
		panic(err)
	}
	interval = tickInterval
	nsDomain = env.Read(env.ENVNCDOMAIN)
	nsPort = env.Read(env.ENVNCPORT)
	log.Printf("Tick interval set to: %s\n", tickInterval)
	for {
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
		time.Sleep(interval)
	}
}
