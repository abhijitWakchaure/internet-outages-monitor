package env

import (
	"log"
	"os"
	"strconv"
)

// Env constants for config
const (
	ENVTIMEZONE = "TIMEZONE"

	ENVTICKINTERVAL = "TICK_INTERVAL"
	ENVNCDOMAIN     = "NC_DOMAIN"
	ENVNCPORT       = "NC_PORT"

	ENVSLACKWEBHOOKURL       = "SLACK_WEBHOOK_URL"
	ENVSLACKNOTIFYONREGISTER = "SLACK_NOTIFY_ON_REGISTER"
)

var defaults map[string]string = map[string]string{
	"TIMEZONE":                 "Asia/Kolkata",
	"TICK_INTERVAL":            "30s",
	"SLACK_NOTIFY_ON_REGISTER": "true",
	"NC_DOMAIN":                "google.com",
	"NC_PORT":                  "443",
}

// Read ...
func Read(name string, defaultVal ...string) string {
	// Lookup env override
	v, ok := os.LookupEnv(name)
	if ok {
		log.Printf("Environment override detected for [%v]\n", name)
		return v
	}
	// Check user default
	if len(defaultVal) > 0 {
		log.Printf("Inline override detected for [%v]\n", name)
		return defaultVal[0]
	}
	v, ok = defaults[name]
	if !ok {
		log.Printf("Failed to lookup value for env var: %s\n", name)
		return ""
	}
	log.Printf("Using default value for [%v]\n", name)
	return v
}

// ReadBool ...
func ReadBool(env string, defaultVal ...string) bool {
	v := Read(env, defaultVal...)
	b, _ := strconv.ParseBool(v)
	return b
}
