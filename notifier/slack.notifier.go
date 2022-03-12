package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/abhijitWakchaure/internet-outages-monitor/env"
)

// Slack implements Notifier interface
type Slack struct {
	loc        *time.Location
	webhookURL string
}

// Register registers a Slack Notifier
func (s *Slack) Register() error {
	var err error
	s.loc, err = time.LoadLocation(env.Read(env.ENVTIMEZONE))
	if err != nil {
		return err
	}
	s.webhookURL = env.Read(env.ENVSLACKWEBHOOKURL)
	if s.webhookURL == "" {
		return fmt.Errorf("%s must be set", env.ENVSLACKWEBHOOKURL)
	}
	notifyOnRegister := env.ReadBool(env.ENVSLACKNOTIFYONREGISTER)
	if !notifyOnRegister {
		return nil
	}
	hostname, err := os.Hostname()
	if err != nil && hostname == "" {
		hostname = "unknown"
	}
	err = s.sendMessage("Slack Notifier registered on host: " + hostname)
	if err != nil {
		return err
	}
	return nil
}

// Notify sends a message to pre-registered Slack channel
func (s *Slack) Notify(message string) error {
	return s.sendMessage(message)
}

func (s *Slack) sendMessage(message string) error {
	msg := map[string]string{
		"text": fmt.Sprintf("%s\n\n- %s", message, s.getTimeString()),
	}
	mBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	res, err := http.Post(s.webhookURL, "application/json", bytes.NewBuffer(mBytes))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Slack returned status: [%s]", res.Status)
	}
	return nil
}

func (s *Slack) getTimeString() string {
	now := time.Now().In(s.loc)
	// Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	// ANSIC       = "Mon Jan _2 15:04:05 2006"
	// UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	// RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	// RFC822      = "02 Jan 06 15:04 MST"
	// RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	// RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	// RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	// RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	// RFC3339     = "2006-01-02T15:04:05Z07:00"
	// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	// Kitchen     = "3:04PM"
	// // Handy time stamps.
	// Stamp      = "Jan _2 15:04:05"
	// StampMilli = "Jan _2 15:04:05.000"
	// StampMicro = "Jan _2 15:04:05.000000"
	// StampNano  = "Jan _2 15:04:05.000000000"

	return now.Format(time.RFC1123)
}
